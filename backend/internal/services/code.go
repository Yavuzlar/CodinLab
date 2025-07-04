package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	extractor "github.com/Yavuzlar/CodinLab/pkg/code_extractor"
	"github.com/Yavuzlar/CodinLab/pkg/docker"
	"github.com/Yavuzlar/CodinLab/pkg/file"
	"github.com/docker/docker/api/types/container"
	"github.com/gofiber/websocket/v2"
)

type codeService struct {
	dockerSDK     docker.IDockerSDK
	commonService domains.ICommonService
	labService    domains.ILabService
	roadService   domains.IRoadService
}

func newCodeService(
	commonService domains.ICommonService,
	labService domains.ILabService,
	roadService domains.IRoadService,
) domains.ICodeService {
	dockerSDK, err := docker.NewDockerSDK()
	if err != nil {
		panic(err)
	}

	return &codeService{
		dockerSDK:     dockerSDK,
		commonService: commonService,
		labService:    labService,
		roadService:   roadService,
	}
}

func (s *codeService) ParseCodeLog(log string) (*domains.UserLog, error) {
	var userLog domains.UserLog
	logArr := strings.Split(log, "|||")
	err := "go: golang.org/x/tool"

	if len(logArr) != 4 {
		if strings.Contains(log, err) {
			return nil, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidTemplateOutput)
		}
		re := regexp.MustCompile(`(?m)^\)?go: downloading golang.org/x/.*\n`)
		log = re.ReplaceAllString(log, "")

		re2 := regexp.MustCompile(`v\d+\.\d+\.\d+`)
		log = re2.ReplaceAllString(log, "")

		userLog.ErrorMessage = log
		return &userLog, nil
	}

	if logArr[0] != "_" {
		if strings.Contains(logArr[0], "Test Passed") {
			userLog.IsCorrect = true
		}
	}

	if logArr[1] != "_" {
		userLog.Output = logArr[1]
	}

	if logArr[2] != "_" {
		userLog.ExpectedOutput = logArr[2]
	}

	if logArr[3] != "_" {
		userLog.ErrorMessage = logArr[3]
	}

	return &userLog, nil
}

func (s *codeService) Pull(ctx context.Context, imageReference, programmingLanguage string, conn *websocket.Conn) error {
	if conn != nil {
		errSocket := conn.WriteJSON(domains.Response{
			Type: "Pull",
			Data: struct {
				Status              int    `json:"status"`
				ProgrammingLanguage string `json:"programminglanguage"`
				Message             string `json:"message"`
			}{
				Status:              200,
				ProgrammingLanguage: programmingLanguage,
				Message:             "Started",
			},
		})
		if errSocket != nil {
			return errSocket
		}
	}
	// Start Docker image pull process
	err := s.dockerSDK.Images().Pull(ctx, imageReference)
	if conn != nil {
		if err != nil {
			errSocket := conn.WriteJSON(domains.Response{
				Type: "Pull",
				Data: struct {
					Status              int    `json:"status"`
					ProgrammingLanguage string `json:"programminglanguage"`
					Message             string `json:"message"`
				}{
					Status:              400,
					ProgrammingLanguage: programmingLanguage,
					Message:             err.Error(),
				},
			})
			if errSocket != nil {
				return errSocket
			}
		} else {
			errSocket := conn.WriteJSON(domains.Response{
				Type: "Pull",
				Data: struct {
					Status              int    `json:"status"`
					ProgrammingLanguage string `json:"programminglanguage"`
					Message             string `json:"message"`
				}{
					Status:              200,
					ProgrammingLanguage: programmingLanguage,
					Message:             "Finished",
				},
			})
			if errSocket != nil {
				return errSocket
			}
		}
	}
	return err
}

func (s *codeService) IsImageExists(ctx context.Context, imageReference string) (isExists bool, err error) {
	resultChan := make(chan struct {
		isExists bool
		err      error
	})

	// Start Docker image existence check asynchronously
	go func() {
		isExists, err := s.dockerSDK.Images().IsImageExists(ctx, imageReference)
		resultChan <- struct {
			isExists bool
			err      error
		}{isExists, err}
	}()

	// After the process is completed, the result is received from the channel
	select {
	case result := <-resultChan:
		if result.err != nil {
			return false, service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerImage)
		}
		return result.isExists, nil
	case <-ctx.Done():
		return false, ctx.Err() // If the context is canceled
	}
}

func (s *codeService) StopContainer(ctx context.Context, containerID string) error {
	errCh := make(chan error)

	go func() {
		errCh <- s.dockerSDK.Container().StopContainer(ctx, containerID)
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrDockerContainerStop, err)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *codeService) createContainerWithCMD(ctx context.Context, image string, cmd []string) (*container.CreateResponse, error) {
	respCh := make(chan *container.CreateResponse)
	errCh := make(chan error)

	go func() {
		resp, err := s.dockerSDK.Container().CreateContainerWithCMD(ctx, image, cmd)
		if err != nil {
			errCh <- err
			close(errCh)
		} else {
			respCh <- resp
			close(respCh)
		}
	}()

	select {
	case resp := <-respCh:
		return resp, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *codeService) CreateBashFile(cmd []string, tests []domains.Test, userID, pathDir string) error {
	templateData, err := os.ReadFile(fmt.Sprintf("%s/main.sh", pathDir))
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(404, domains.ErrRoadMainSHNotFound)
	}
	content := string(templateData)

	var testsStrBuilder strings.Builder
	for _, test := range tests {
		for _, t := range test.GetOutput() {
			switch v := t.(type) {
			case int:
				testsStrBuilder.WriteString(fmt.Sprintf(" %d", v))
			case string:
				testsStrBuilder.WriteString(fmt.Sprintf(" \"%s\"", v))
			default:
				testsStrBuilder.WriteString(fmt.Sprintf(" %v", v))
			}
		}

	}
	testsStr := testsStrBuilder.String()
	content = strings.Replace(content, "-tests-", testsStr, -1)

	if err := file.CheckDir(fmt.Sprintf("usercodes/%v/paths", userID)); err != nil {
		return err
	}

	file.CreateFileAndWrite(fmt.Sprintf("usercodes/%v/paths/main.sh", userID), content)

	return nil
}

func (s *codeService) RunContainerWithTar(ctx context.Context, image, tmpCodePath, fileName string, cmd []string, conn *websocket.Conn) (string, error) {
	// Start the container run process
	resp, err := s.createContainerWithCMD(ctx, image, cmd)
	if err != nil {
		return "", err
	}

	if conn != nil {
		err = conn.WriteJSON(domains.Response{
			Type: "container",
			Data: struct {
				ID string `json:"id"`
			}{
				ID: resp.ID,
			},
		})
		if err != nil {
			return "", err
		}
	}

	logs, err := s.dockerSDK.Container().RunContainerWithTar(ctx, tmpCodePath, fileName, *resp)
	if err != nil {
		if strings.Contains(err.Error(), "No such image") {
			re := regexp.MustCompile(`No such image: (.+)`)
			matches := re.FindStringSubmatch(err.Error())
			if len(matches) > 1 {
				imageName := matches[1]
				return "", service_errors.NewServiceErrorWithMessage(404, fmt.Sprintln(domains.ErrDockerImageNotFound, imageName))
			}
			return "", service_errors.NewServiceErrorWithMessage(404, domains.ErrDockerImageNotFound)
		}
		return "", service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerLogs)
	}

	return logs, nil
}

func (s *codeService) UploadUserCode(userID, programmingID, labPathID, codeType, fileExtension, content string) (string, string, error) {
	if err := s.createCodeFile(userID); err != nil {
		return "", "", err
	}

	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return "", "", service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
	}

	intLabPathID, err := strconv.Atoi(labPathID)
	if err != nil {
		return "", "", service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidLabID)
	}

	codePath := s.generateUserCodePath(userID, codeType, intProgrammingID, intLabPathID, fileExtension)
	codeTmpPath := s.generateUserCodeTmpPath(userID, codeType, intProgrammingID, intLabPathID, fileExtension)

	if codeType == domains.TypeLab {
		if err := s.CreateFileAndWrite(codePath, content); err != nil {
			return "", "", err
		}
		if err := s.CreateFileAndWrite(codeTmpPath, ""); err != nil {
			return "", "", err
		}
	}

	if codeType == domains.TypePath {
		if err := s.CreateFileAndWrite(codePath, content); err != nil {
			return "", "", err
		}
		if err := s.CreateFileAndWrite(codeTmpPath, ""); err != nil {
			return "", "", err
		}
	}

	return codePath, codeTmpPath, nil
}

func (s *codeService) CodeDockerTemplateGenerator(templatePath, funcName, userCode string, tests []domains.Test, inventoryName string) (string, error) {
	templateMap, err := s.readTemplate(templatePath)
	if err != nil {
		return "", err
	}
	if !strings.Contains(userCode, funcName) {
		return "", service_errors.NewServiceErrorWithMessage(400, fmt.Sprintf("Need %v function", funcName))
	}

	frontImports, cleanedCode := extractor.ExtractImports(userCode, true)
	if !strings.EqualFold(funcName, "main") {
		cleanedCode, err = extractor.ExtractMainFunction(cleanedCode)
		if err != nil {
			return "", err
		}
	} else {
		funcName = "test"
		cleanedCode = extractor.ExtractFuncName(cleanedCode, funcName)
	}

	docker := templateMap["docker"]

	checks := s.createChecks(templateMap["check"], tests)
	docker = strings.Replace(docker, "$checks$", checks, -1)
	docker = strings.Replace(docker, "$res$", fmt.Sprint(len(tests)-1), -1)
	docker = strings.Replace(docker, "$usercode$", cleanedCode, -1)
	docker = strings.Replace(docker, "$funcname$", funcName, -1)
	docker = strings.Replace(docker, "$success$", "Test Passed", -1)

	dockerImports, newDocker := extractor.ExtractImports(docker, false)

	allImports := ""
	if strings.EqualFold(inventoryName, "GO") && !strings.EqualFold(funcName, "main") {
		allImports = extractor.CombineImportsForGolang(frontImports, dockerImports)
	} else {
		allImports = s.bindImports(dockerImports, frontImports)
	}

	docker = strings.Replace(newDocker, "$imps$", allImports, -1)
	return docker, nil
}

func (s *codeService) bindImports(dockerImports, frontImports string) string {
	dockerImportsLines := strings.Split(dockerImports, "\n")
	frontImportsLines := strings.Split(frontImports, "\n")

	for _, dockerImport := range dockerImportsLines {
		if !strings.Contains(frontImports, dockerImport) {
			frontImportsLines = append(frontImportsLines, dockerImport)
		}
	}

	return strings.Join(frontImportsLines, "\n")
}

func (s *codeService) CodeFrontendTemplateGenerator(templatePath, funcName string) (string, error) {
	templateMap, err := s.readTemplate(templatePath)
	if err != nil {
		return "", err
	}
	frontend := templateMap["frontend"]

	frontend = strings.Replace(frontend, "$funcname$", funcName, -1)

	return frontend, nil
}

func (s *codeService) GetFrontendTemplate(userID, programmingID, labPathID, labRoadType string, fileExtension string, checkHistory bool) (string, error) {
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
	}

	intLabPathID, err := strconv.Atoi(labPathID)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidLabID)
	}

	var frontendTemplate string
	if labRoadType == domains.TypeLab {
		if checkHistory {
			history := s.readFrontendTemplateHistory(userID, intProgrammingID, intLabPathID, labRoadType, fileExtension)
			if len(history) > 0 {
				return history, nil
			}
		}

		lab, err := s.labService.GetLabByID(userID, labPathID)
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, domains.ErrLabNotFound)
		}

		for _, ct := range lab.GetQuest().GetCodeTemplates() {
			if ct.GetProgrammingID() == intProgrammingID {
				frontendTemplate, err = s.CodeFrontendTemplateGenerator(ct.GetTemplatePath(), lab.GetQuest().GetFuncName())
				if err != nil {
					return "", service_errors.NewServiceErrorWithMessage(500, domains.ErrGettingFrontendTemplate)
				}

				break
			}
		}

	} else if labRoadType == domains.TypePath {
		if checkHistory {
			history := s.readFrontendTemplateHistory(userID, intProgrammingID, intLabPathID, labRoadType, fileExtension)
			if len(history) > 0 {
				return history, nil
			}
		}

		path, err := s.roadService.GetPathByID(userID, programmingID, labPathID)
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, domains.ErrPathNotFound)
		}
		for _, ct := range path.GetQuest().GetCodeTemplates() {
			if ct.GetProgrammingID() == intProgrammingID {
				frontendTemplate, err = s.CodeFrontendTemplateGenerator(ct.GetTemplatePath(), path.GetQuest().GetFuncName())
				if err != nil {
					return "", service_errors.NewServiceErrorWithMessage(500, domains.ErrGettingFrontendTemplate)
				}

				break
			}
		}

	}

	if frontendTemplate == "" {
		return "", service_errors.NewServiceErrorWithMessage(404, domains.ErrTemplateNotFound)
	}

	return frontendTemplate, nil
}

func (s *codeService) DeleteFrontendTemplateHistory(userID, programmingID, labPathID, labRoadType, fileExtension string) (err error) {
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
	}

	intLabPathID, err := strconv.Atoi(labPathID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidLabID)
	}

	codePath := s.generateUserCodePath(userID, labRoadType, intProgrammingID, intLabPathID, fileExtension)
	if _, err := os.Stat(codePath); err == nil {
		err := os.Remove(codePath)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(404, domains.ErrUserCodeNotFound)
		}
	}

	codeTmpPath := s.generateUserCodeTmpPath(userID, labRoadType, intProgrammingID, intLabPathID, fileExtension)
	if _, err := os.Stat(codeTmpPath); err == nil {
		err := os.Remove(codeTmpPath)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(404, domains.ErrUserCodeNotFound)
		}
	}

	return nil
}

func (s *codeService) CreateFileAndWrite(filePath, content string) (err error) {
	err = file.CreateFileAndWrite(filePath, content)
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerCouldNotCreateFile)
	}
	return file.CreateFileAndWrite(filePath, content)
}

func (s *codeService) createChecks(check string, tests []domains.Test) string {
	var checks strings.Builder

	for i, test := range tests {
		tmp := check
		tmp = strings.Replace(tmp, "$rnd$", fmt.Sprintf("%v", i), -1)

		var inputs []string
		for _, in := range test.GetInput() {
			switch v := in.(type) {
			case string:
				// Add double quotes around string inputs
				inputs = append(inputs, `"`+v+`"`)
			case []string:
				// If input is a []string, it will be formatted correctly with double quotes around each element
				var quotedStrings []string
				for _, elem := range v {
					quotedStrings = append(quotedStrings, `"`+elem+`"`)
				}
				inputs = append(inputs, strings.Join(quotedStrings, ", "))
			case []interface{}:
				var arrayElements []string
				for _, elem := range v {
					switch e := elem.(type) {
					case string:
						arrayElements = append(arrayElements, `"`+e+`"`)
					case []string:
						// nested []string is handled within []interface{}
						var quotedStrings []string
						for _, s := range e {
							quotedStrings = append(quotedStrings, `"`+s+`"`)
						}
						arrayElements = append(arrayElements, "["+strings.Join(quotedStrings, ", ")+"]")
					default:
						// other types are handled normally
						arrayElements = append(arrayElements, fmt.Sprintf("%v", e))
					}
				}
				inputs = append(inputs, strings.Join(arrayElements, ", "))
			default:
				inputs = append(inputs, fmt.Sprintf("%v", v))
			}
		}
		tmp = strings.Replace(tmp, "$input$", strings.Join(inputs, ", "), -1)

		// Handle output replacement
		var outputs []string
		var fails []string
		for _, out := range test.GetOutput() {
			switch v := out.(type) {
			case string:
				outputs = append(outputs, `"`+v+`"`)
				fails = append(fails, fmt.Sprintf("%v", v))
			case []string:
				var quotedStrings []string
				for _, elem := range v {
					quotedStrings = append(quotedStrings, `"`+elem+`"`)
				}
				outputs = append(outputs, strings.Join(quotedStrings, ", "))
				fails = append(fails, strings.Join(quotedStrings, ", "))
			case []interface{}:
				var arrayElements []string
				for _, elem := range v {
					switch e := elem.(type) {
					case string:
						arrayElements = append(arrayElements, `"`+e+`"`)
					case []string:
						var quotedStrings []string
						for _, s := range e {
							quotedStrings = append(quotedStrings, `"`+s+`"`)
						}
						arrayElements = append(arrayElements, "["+strings.Join(quotedStrings, ", ")+"]")
					default:
						arrayElements = append(arrayElements, fmt.Sprintf("%v", e))
					}
				}
				outputs = append(outputs, strings.Join(arrayElements, ", "))
				fails = append(fails, strings.Join(arrayElements, ", "))
			default:
				// Directly use other types (int, etc.)
				outputs = append(outputs, fmt.Sprintf("%v", v))
				fails = append(fails, fmt.Sprintf("%v", v))
			}
		}
		tmp = strings.Replace(tmp, "$output$", strings.Join(outputs, ", "), -1)
		if strings.Contains(check, "$out$") {
			tmp = strings.Replace(tmp, "$out$", strings.Join(fails, ", "), -1)
		}

		checks.WriteString(tmp + "\n")
	}

	return checks.String()
}

func (s *codeService) createCodeFile(userID string) (err error) {
	mainDir := "usercodes"
	userDir := fmt.Sprintf("%v/%v", mainDir, userID)
	labDir := fmt.Sprintf("%v/labs", userDir)
	pathDir := fmt.Sprintf("%v/paths", userDir)

	// Check and create mainDir if not exists
	if err := file.CheckDir(mainDir); err != nil {
		if err = file.CreateDir(mainDir); err != nil {
			return service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerFileError)
		}
	}

	// Check and create userDir, labDir, pathDir
	if err := file.CheckDir(userDir); err != nil {
		if err = file.CreateDir(userDir); err != nil {
			return service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerFileError)
		}
	}

	// Check and create labDir if not exists
	if err := file.CheckDir(labDir); err != nil {
		if err = file.CreateDir(labDir); err != nil {
			return service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerFileError)
		}
	}

	// Check and create pathDir if not exists
	if err := file.CheckDir(pathDir); err != nil {
		if err = file.CreateDir(pathDir); err != nil {
			return service_errors.NewServiceErrorWithMessage(500, domains.ErrDockerFileError)
		}
	}

	return nil
}

func (s *codeService) readTemplate(templatePath string) (map[string]string, error) {
	// Read the file
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(404, domains.ErrTemplateNotFound)
	}
	content := string(templateData)

	// Split text by ## signs
	sections := strings.Split(content, "##")

	// Define variables
	template := make(map[string]string)

	// Scan sections and assign them to relevant keys
	for _, section := range sections {
		trimmedSection := strings.TrimSpace(section)

		// If there is a section title, use this title as a key and assign the content to the map.
		if strings.HasPrefix(trimmedSection, "FRONTEND") {
			template["frontend"] = strings.TrimSpace(strings.TrimPrefix(trimmedSection, "FRONTEND"))
		} else if strings.HasPrefix(trimmedSection, "DOCKER") {
			template["docker"] = strings.TrimSpace(strings.TrimPrefix(trimmedSection, "DOCKER"))
		} else if strings.HasPrefix(trimmedSection, "CHECK") {
			template["check"] = strings.TrimSpace(strings.TrimPrefix(trimmedSection, "CHECK"))
		}
	}

	return template, nil
}

func (s *codeService) readFrontendTemplateHistory(userID string, programmingLanguageID, PathLabID int, codeType, fileExtension string) string {
	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingLanguageID, PathLabID, fileExtension)

	var codePath string

	if codeType == domains.TypeLab {
		codePath = labDir + fileName
	} else if codeType == domains.TypePath {
		codePath = pathDir + fileName
	}

	templateData, err := os.ReadFile(codePath)
	if err != nil {
		return ""
	}
	content := string(templateData)

	return content
}

func (s *codeService) generateUserCodePath(userID, labRoadType string, programmingID, labPathID int, fileExtension string) string {
	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingID, labPathID, fileExtension)

	var codePath string

	if labRoadType == domains.TypeLab {
		codePath = labDir + fileName
	} else if labRoadType == domains.TypePath {
		codePath = pathDir + fileName
	}

	return codePath
}

func (s *codeService) generateUserCodeTmpPath(userID, labRoadType string, programmingID, labPathID int, fileExtension string) string {
	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingID, labPathID, fileExtension)

	var codeTmpPath string

	if labRoadType == domains.TypeLab {
		codeTmpPath = labDir + "tmp-" + fileName
	} else if labRoadType == domains.TypePath {
		codeTmpPath = pathDir + "tmp-" + fileName
	}

	return codeTmpPath
}

func (s *codeService) SaveUserHistory(conn *websocket.Conn, messages []byte, userID string) error {
	var req domains.UserCodeRequest
	err := json.Unmarshal(messages, &req)
	if err != nil {
		return err
	}
	if conn != nil {
		if req.Type == "close" {
			stringProgrammingID := strconv.Itoa(int(req.Data.ProgrammingID))
			stringLabPathID := strconv.Itoa(int(req.Data.LabPathID))
			programmingInformation, err := s.commonService.GetInventoryInformation(stringProgrammingID, "")
			if err != nil {
				return err
			}

			if strings.EqualFold(req.Data.LabPathType, domains.TypeLab) {
				req.Data.LabPathType = domains.TypeLab
			} else if strings.EqualFold(req.Data.LabPathType, domains.TypePath) {
				req.Data.LabPathType = domains.TypePath
			} else {
				return service_errors.NewServiceErrorWithMessage(400, "invalid type")
			}

			if _, _, err := s.UploadUserCode(userID, stringProgrammingID, stringLabPathID, req.Data.LabPathType, programmingInformation.GetFileExtension(), req.Data.UserCode); err != nil {
				return err
			}

			errSocket := conn.WriteJSON(domains.Response{
				Type: "close",
				Data: struct {
					Status  int    `json:"status"`
					Message string `json:"message"`
				}{
					Status:  200,
					Message: "History Saved Successfully",
				},
			})
			if errSocket != nil {
				return errSocket
			}
		}
	} else {
		return service_errors.NewServiceErrorWithMessage(400, "Web socket connection error")
	}

	return nil
}
