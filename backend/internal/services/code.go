package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
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
	return s.dockerSDK.Images().IsImageExists(ctx, imageReference)
}

func (s *codeService) RunContainerWithTar(ctx context.Context, image, tmpCodePath, fileName string, cmd []string) (string, error) {
	return s.dockerSDK.Container().RunContainerWithTar(ctx, image, cmd, tmpCodePath, fileName)
}

// Bunu Answer Kısmınlarında Kullanacaksın.
func (s *codeService) UploadUserCode(ctx context.Context, userID string, programmingLanguageID, PathLabID int, codeType, fileExtention, content string) (string, error) {
	if err := s.createCodeFile(userID); err != nil {
		return "", err
	}

	codePath := s.generateUserCodePath(userID, codeType, programmingLanguageID, PathLabID, fileExtention)
	codeTmpPath := s.generateUserCodeTmpPath(userID, codeType, programmingLanguageID, PathLabID, fileExtention)

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
	docker := templateMap["docker"]

	checks := s.createChecks(templateMap["check"], tests)

	docker = strings.Replace(docker, "$checks$", checks, -1)
	docker = strings.Replace(docker, "$usercode$", userCode, -1)
	docker = strings.Replace(docker, "$funcname$", funcName, -1)
	docker = strings.Replace(docker, "$success$", "Test Passed", -1)

	return docker, nil
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

func (s *codeService) GetFrontendTemplate(userID, labRoadType string, programmingID, labPathID int, fileExtention string) (string, error) {
	var frontendTemplate string

	if labRoadType == domains.TypeLab {
		history := s.readFrontendTemplateHistory(userID, programmingID, labPathID, labRoadType, fileExtention)
		if len(history) > 0 {
			return history, nil
		}

		lab, err := s.labService.GetLabByID(userID, labPathID)
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Lab Not Found")
		}

		for _, ct := range lab.GetQuest().GetCodeTemplates() {
			if ct.GetProgrammingID() == programmingID {
				frontendTemplate, err = s.CodeFrontendTemplateGenerator(ct.GetTemplatePath(), lab.GetQuest().GetFuncName())
				if err != nil {
					return "", service_errors.NewServiceErrorWithMessage(404, "Error while getting frontend template")
				}

				break
			}
		}

	} else if labRoadType == domains.TypePath {
		history := s.readFrontendTemplateHistory(userID, programmingID, labPathID, labRoadType, fileExtention)
		if len(history) > 0 {
			return history, nil
		}

		path, err := s.roadService.GetRoadByID(userID, programmingID, labPathID)
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Path Not Found")
		}
		for _, ct := range path.GetQuest().GetCodeTemplates() {
			if ct.GetProgrammingID() == programmingID {
				frontendTemplate, err = s.CodeFrontendTemplateGenerator(ct.GetTemplatePath(), path.GetQuest().GetFuncName())
				if err != nil {
					return "", service_errors.NewServiceErrorWithMessage(404, "Error while getting frontend template")
				}

				break

			}
		}

	}

	return frontendTemplate, nil
}

func (s *codeService) DeleteFrontendTemplateHistory(userID, labRoadType string, programmingID, labPathID int, fileExtention string) (err error) {

	codePath := s.generateUserCodePath(userID, labRoadType, programmingID, labPathID, fileExtention)
	if _, err := os.Stat(codePath); err == nil {
		err := os.Remove(codePath)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(500, "Failed to delete file")
		}
	}

	codeTmpPath := s.generateUserCodeTmpPath(userID, labRoadType, programmingID, labPathID, fileExtention)
	if _, err := os.Stat(codeTmpPath); err == nil {
		err := os.Remove(codeTmpPath)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(500, "Failed to delete file")
		}
	}

	return nil
}

func (s *codeService) CreateFileAndWrite(filePath, content string) (err error) {
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
			case []interface{}:
				// If input is an array, format it correctly without square brackets
				var arrayElements []string
				for _, elem := range v {
					arrayElements = append(arrayElements, fmt.Sprintf("%v", elem))
				}
				// Join array elements with a comma, without square brackets
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
			case []interface{}:
				// If output is an array, format it correctly without square brackets
				var arrayElements []string
				for _, elem := range v {
					arrayElements = append(arrayElements, fmt.Sprintf("%v", elem))
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
		return nil, err
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
