package services

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	extractor "github.com/Yavuzlar/CodinLab/pkg/code_extractor"
	"github.com/Yavuzlar/CodinLab/pkg/docker"
	"github.com/Yavuzlar/CodinLab/pkg/file"
)

type codeService struct {
	dockerSDK      docker.IDockerSDK
	labRoadService domains.ILabRoadService
	labService     domains.ILabService
	roadService    domains.IRoadService
}

func newCodeService(
	labRoadService domains.ILabRoadService,
	labService domains.ILabService,
	roadService domains.IRoadService,
) domains.ICodeService {
	dockerSDK, err := docker.NewDockerSDK()
	if err != nil {
		panic(err)
	}

	return &codeService{
		dockerSDK:      dockerSDK,
		labRoadService: labRoadService,
		labService:     labService,
		roadService:    roadService,
	}
}

func (s *codeService) Pull(ctx context.Context, imageReference string) (err error) {
	resultChan := make(chan error)

	// Asenkron olarak Docker image pull işlemini başlatın
	go func() {
		err := s.dockerSDK.Images().Pull(ctx, imageReference)
		resultChan <- err
	}()

	// İndirme işlemi tamamlandığında kanaldan hata bilgisi alınır
	select {
	case err := <-resultChan:
		return err
	case <-ctx.Done():
		return ctx.Err() // Eğer context iptal edilirse, bu hata döndürülür
	}
}

func (s *codeService) IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error) {
	isExsits, err = s.dockerSDK.Images().IsImageExists(ctx, imageReference)
	if err != nil {
		return false, service_errors.NewServiceErrorWithMessage(400, "Docker Image Is Exist Error")
	}

	return
}

func (s *codeService) RunContainerWithTar(ctx context.Context, image, tmpCodePath, fileName string, cmd []string) (string, error) {

	logs, err := s.dockerSDK.Container().RunContainerWithTar(ctx, image, cmd, tmpCodePath, fileName)
	if err != nil {
		fmt.Println(err)
		return "", service_errors.NewServiceErrorWithMessage(500, "Unable to read docker logs")
	}

	return logs, nil
}

// Bunu Answer Kısmınlarında Kullanacaksın.
func (s *codeService) UploadUserCode(ctx context.Context, userID, programmingID, labPathID string, codeType, fileExtention, content string) (string, error) {
	if err := s.createCodeFile(userID); err != nil {
		return "", err
	}

	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(400, "Invalid Programming Language ID")
	}

	intLabPathID, err := strconv.Atoi(labPathID)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(400, "Invalid Lab ID")
	}

	codePath := s.generateUserCodePath(userID, codeType, intProgrammingID, intLabPathID, fileExtention)
	codeTmpPath := s.generateUserCodeTmpPath(userID, codeType, intProgrammingID, intLabPathID, fileExtention)

	if codeType == domains.TypeLab {
		if err := s.CreateFileAndWrite(codePath, content); err != nil {
			return "", err
		}
		if err := s.CreateFileAndWrite(codeTmpPath, ""); err != nil {
			return "", err
		}
	}

	if codeType == domains.TypePath {
		if err := s.CreateFileAndWrite(codePath, content); err != nil {
			return "", err
		}
		if err := s.CreateFileAndWrite(codeTmpPath, ""); err != nil {
			return "", err
		}
	}

	return codeTmpPath, nil
}

func (s *codeService) CodeDockerTemplateGenerator(templatePath, funcName, userCode string, tests []domains.Test) (string, error) {
	templateMap, err := s.readTemplate(templatePath)
	if err != nil {
		return "", err
	}
	if !strings.Contains(userCode, funcName) {
		return "", service_errors.NewServiceErrorWithMessage(400, fmt.Sprintf("Need %v function", funcName))
	}

	frontImports, cleanedCode := extractor.ExtractImports(userCode, true)
	newUserCode, err := extractor.ExtractMainFunction(cleanedCode)
	if err != nil {
		return "", err
	}

	docker := templateMap["docker"]

	checks := s.createChecks(templateMap["check"], tests)

	docker = strings.Replace(docker, "$checks$", checks, -1)
	docker = strings.Replace(docker, "$usercode$", newUserCode, -1)
	docker = strings.Replace(docker, "$funcname$", funcName, -1)
	docker = strings.Replace(docker, "$success$", "Test Passed", -1)

	dockerImports, newDocker := extractor.ExtractImports(docker, false)

	allImports := s.bindImports(dockerImports, frontImports)

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

func (s *codeService) GetFrontendTemplate(userID, programmingID, labPathID, labRoadType string, fileExtention string) (string, error) {
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(400, "Invalid Programming Language ID")
	}

	intLabPathID, err := strconv.Atoi(labPathID)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(400, "Invalid Lab ID")
	}

	var frontendTemplate string
	if labRoadType == domains.TypeLab {
		history := s.readFrontendTemplateHistory(userID, intProgrammingID, intLabPathID, labRoadType, fileExtention)
		if len(history) > 0 {
			return history, nil
		}

		lab, err := s.labService.GetLabByID(userID, labPathID)
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Lab Not Found")
		}

		for _, ct := range lab.GetQuest().GetCodeTemplates() {
			if ct.GetProgrammingID() == intProgrammingID {
				frontendTemplate, err = s.CodeFrontendTemplateGenerator(ct.GetTemplatePath(), lab.GetQuest().GetFuncName())
				if err != nil {
					return "", service_errors.NewServiceErrorWithMessage(404, "Error while getting frontend template")
				}

				break
			}
		}

	} else if labRoadType == domains.TypePath {
		history := s.readFrontendTemplateHistory(userID, intProgrammingID, intLabPathID, labRoadType, fileExtention)
		if len(history) > 0 {
			return history, nil
		}

		path, err := s.roadService.GetRoadByID(userID, programmingID, labPathID)
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Path Not Found")
		}
		for _, ct := range path.GetQuest().GetCodeTemplates() {
			if ct.GetProgrammingID() == intProgrammingID {
				frontendTemplate, err = s.CodeFrontendTemplateGenerator(ct.GetTemplatePath(), path.GetQuest().GetFuncName())
				if err != nil {
					return "", service_errors.NewServiceErrorWithMessage(404, "Error while getting frontend template")
				}

				break
			}
		}

	}

	if frontendTemplate == "" {
		return "", service_errors.NewServiceErrorWithMessage(400, "Frontend Template could not be created")
	}

	return frontendTemplate, nil
}

func (s *codeService) DeleteFrontendTemplateHistory(userID, programmingID, labPathID, labRoadType, fileExtention string) (err error) {
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(400, "Invalid Programming Language ID")
	}

	intLabPathID, err := strconv.Atoi(labPathID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(400, "Invalid Lab ID")
	}

	codePath := s.generateUserCodePath(userID, labRoadType, intProgrammingID, intLabPathID, fileExtention)
	if _, err := os.Stat(codePath); err == nil {
		err := os.Remove(codePath)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(500, "Failed to delete file")
		}
	}

	codeTmpPath := s.generateUserCodeTmpPath(userID, labRoadType, intProgrammingID, intLabPathID, fileExtention)
	if _, err := os.Stat(codeTmpPath); err == nil {
		err := os.Remove(codeTmpPath)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(500, "Failed to delete file")
		}
	}

	return nil
}

func (s *codeService) CreateFileAndWrite(filePath, content string) (err error) {
	err = file.CreateFileAndWrite(filePath, content)
	if err != nil {
		return service_errors.NewServiceErrorWithMessage(500, "Could not create file and write data into it")
	}
	return file.CreateFileAndWrite(filePath, content)
}

func (s *codeService) createChecks(check string, tests []domains.Test) string {
	var checks strings.Builder

	for i, test := range tests {
		tmp := check
		tmp = strings.Replace(tmp, "$rnd$", fmt.Sprintf("%v", i), -1)

		// Handle input replacement
		var inputs []string
		for _, in := range test.GetInput() {
			switch v := in.(type) {
			case string:
				// Add double quotes around string inputs
				inputs = append(inputs, `"`+v+`"`)
			case []string:
				// If input is a []string, format it correctly with double quotes around each element
				var quotedStrings []string
				for _, elem := range v {
					quotedStrings = append(quotedStrings, `"`+elem+`"`)
				}
				inputs = append(inputs, strings.Join(quotedStrings, ", "))
			case []interface{}:
				// Handle []interface{} with mixed types
				var arrayElements []string
				for _, elem := range v {
					switch e := elem.(type) {
					case string:
						arrayElements = append(arrayElements, `"`+e+`"`)
					case []string:
						// Handle nested []string within []interface{}
						var quotedStrings []string
						for _, s := range e {
							quotedStrings = append(quotedStrings, `"`+s+`"`)
						}
						arrayElements = append(arrayElements, "["+strings.Join(quotedStrings, ", ")+"]")
					default:
						// Handle other types normally
						arrayElements = append(arrayElements, fmt.Sprintf("%v", e))
					}
				}
				inputs = append(inputs, strings.Join(arrayElements, ", "))
			default:
				// Directly use other types (int, etc.)
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
				// Add double quotes around string outputs
				outputs = append(outputs, `"`+v+`"`)
				fails = append(fails, fmt.Sprintf("%v", v))
			case []string:
				// If output is a []string, format it correctly with double quotes around each element
				var quotedStrings []string
				for _, elem := range v {
					quotedStrings = append(quotedStrings, `"`+elem+`"`)
				}
				outputs = append(outputs, strings.Join(quotedStrings, ", "))
				fails = append(fails, strings.Join(quotedStrings, ", "))
			case []interface{}:
				// Handle []interface{} with mixed types
				var arrayElements []string
				for _, elem := range v {
					switch e := elem.(type) {
					case string:
						arrayElements = append(arrayElements, `"`+e+`"`)
					case []string:
						// Handle nested []string within []interface{}
						var quotedStrings []string
						for _, s := range e {
							quotedStrings = append(quotedStrings, `"`+s+`"`)
						}
						arrayElements = append(arrayElements, "["+strings.Join(quotedStrings, ", ")+"]")
					default:
						// Handle other types normally
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

		checks.WriteString(tmp + "\n       ")
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
			return err
		}
	}

	// Check and create userDir, labDir, pathDir
	if err := file.CheckDir(userDir); err != nil {
		if err = file.CreateDir(userDir); err != nil {
			return err
		}
	}

	// Check and create labDir if not exists
	if err := file.CheckDir(labDir); err != nil {
		if err = file.CreateDir(labDir); err != nil {
			return err
		}
	}

	// Check and create pathDir if not exists
	if err := file.CheckDir(pathDir); err != nil {
		if err = file.CreateDir(pathDir); err != nil {
			return err
		}
	}

	return nil
}

func (s *codeService) readTemplate(templatePath string) (map[string]string, error) {
	// Dosyayı oku
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(500, "File could not be read")
	}
	content := string(templateData)

	// Metni ## işaretlerine göre böl
	sections := strings.Split(content, "##")

	// Değişkenleri tanımla
	template := make(map[string]string)

	// Bölümleri tara ve ilgili anahtarlara ata
	for _, section := range sections {
		trimmedSection := strings.TrimSpace(section)

		// Eğer bölüm başlığı varsa bu başlığı anahtar olarak kullan ve içeriği map'e ata
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

func (s *codeService) readFrontendTemplateHistory(userID string, programmingLanguageID, PathLabID int, codeType, fileExtention string) string {
	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingLanguageID, PathLabID, fileExtention)

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

func (s *codeService) generateUserCodePath(userID, labRoadType string, programmingID, labPathID int, fileExtention string) string {
	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingID, labPathID, fileExtention)

	var codePath string

	if labRoadType == domains.TypeLab {
		codePath = labDir + fileName
	} else if labRoadType == domains.TypePath {
		codePath = pathDir + fileName
	}

	return codePath
}

func (s *codeService) generateUserCodeTmpPath(userID, labRoadType string, programmingID, labPathID int, fileExtention string) string {
	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingID, labPathID, fileExtention)

	var codeTmpPath string

	if labRoadType == domains.TypeLab {
		codeTmpPath = labDir + "tmp-" + fileName
	} else if labRoadType == domains.TypePath {
		codeTmpPath = pathDir + "tmp-" + fileName
	}

	return codeTmpPath
}
