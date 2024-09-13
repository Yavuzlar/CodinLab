package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/pkg/docker"
	"github.com/Yavuzlar/CodinLab/pkg/file"
)

type codeService struct {
	dockerSDK      docker.IDockerSDK
	labRoadService domains.ILabRoadService
	labService     domains.ILabService
	roadService    domains.IRoadService
}

func NewCodeService() domains.ICodeService {
	dockerSDK, err := docker.NewDockerSDK()
	if err != nil {
		panic(err)
	}

	return &codeService{
		dockerSDK: dockerSDK,
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

	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingLanguageID, PathLabID, fileExtention)

	var codePath, codeTmpPath string

	if codeType == domains.TypeLab {
		codePath = labDir + fileName
		codeTmpPath = labDir + "tmp-" + fileName
	} else if codeType == domains.TypePath {
		codePath = pathDir + fileName
		codeTmpPath = pathDir + "tmp-" + fileName
	}

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

	checks := s.createChecks(templateMap["check"], templateMap["success"], tests)

	docker = strings.Replace(docker, "$checks$", checks, -1)
	docker = strings.Replace(docker, "$usercode$", userCode, -1)
	docker = strings.Replace(docker, "$funcname$", funcName, -1)

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

func (s *codeService) CreateFileAndWrite(filePath, content string) (err error) {
	return file.CreateFileAndWrite(filePath, content)
}

func (s *codeService) createChecks(check, success string, tests []domains.Test) string {
	var checks strings.Builder

	for i, test := range tests {
		tmp := check
		tmp = strings.Replace(tmp, "result", fmt.Sprintf("result%v", i), -1)

		// Handle input replacement
		var inputs []string
		for _, in := range test.GetInput() {
			switch v := in.(type) {
			case string:
				// Add double quotes around string inputs
				inputs = append(inputs, `"`+v+`"`)
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

			default:
				// Directly use other types (int, etc.)
				outputs = append(outputs, fmt.Sprintf("%v", v))
				fails = append(fails, fmt.Sprintf("%v", v))
			}
		}
		tmp = strings.Replace(tmp, "$output$", strings.Join(outputs, ", "), -1)
		if strings.Contains(check, "$out$") { //in c++ strings should not contain "". This will correct it. Add $out$ instead of $output$ in failed message
			tmp = strings.Replace(tmp, "$out$", strings.Join(fails, ", "), -1)
		}

		checks.WriteString(tmp + "\n       ")
	}
	checks.WriteString(success + "\n")

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
		} else if strings.HasPrefix(trimmedSection, "SUCCESS") {
			template["success"] = strings.TrimSpace(strings.TrimPrefix(trimmedSection, "SUCCESS"))
		}
	}

	return template, nil
}

/* func (s *codeService) GetTemplate(labPathID, programmingID int, userID, labPathType string) (frontendContent string, err error) {
	inventoryInformation, err := s.labRoadService.GetInventoryInformation(int32(programmingID))
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessage(404, "Programming Language Not Found")
	}

	var lab *domains.Lab
	var path *domains.Path
	var codeTemplate domains.CodeTemplate
	if labPathType == domains.TypeLab {
		if programmingID != 0 {
			lab, err = s.labService.GetLabByID(userID, labPathID)
		} else {
			lab, err = s.labService.GetLabByID(userID, 1)
		}
		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Error While Getting Lab")
		}
		if lab == nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Lab Not Found")
		}
		for _, codeTmp := range lab.GetQuest().GetCodeTemplates() {
			if codeTmp.GetProgrammingID() == inventoryInformation.GetID() {
				codeTemplate = codeTmp
			}
		}
		frontendContent := s.CodeFrontendTemplateGenerator(inventoryInformation.GetName(), lab.GetQuest().GetFuncName(), codeTemplate.GetFrontend(), lab.GetQuest().GetParams(), lab.GetQuest().GetReturns(), codeTemplate.GetQuestImports())
		return frontendContent, nil
	} else if labPathType == domains.TypePath {
		if programmingID != 0 {
			path, err = s.roadService.GetRoadByID(userID, programmingID, labPathID)
		} else {
			path, err = s.roadService.GetRoadByID(userID, programmingID, 1)
		}

		if err != nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Error While Getting Path")
		}
		if path == nil {
			return "", service_errors.NewServiceErrorWithMessage(404, "Path Not Found")
		}
		codeTemplate = path.GetQuest().GetCodeTemplates()[0]
		frontendContent := s.CodeFrontendTemplateGenerator(inventoryInformation.GetName(), path.GetQuest().GetFuncName(), codeTemplate.GetFrontend(), path.GetQuest().GetParams(), path.GetQuest().GetReturns(), codeTemplate.GetQuestImports())
		return frontendContent, nil
	}

	return frontendContent, nil

}
*/
