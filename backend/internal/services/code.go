package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/pkg/docker"
	"github.com/Yavuzlar/CodinLab/pkg/file"
)

type codeService struct {
	dockerSDK docker.IDockerSDK
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

func (s *codeService) CodeDockerTemplateGenerator(template, check, success, userCode, funcName string, tests []domains.Test, returns []domains.Returns) (string, error) {
	if !strings.Contains(userCode, funcName) {
		return "", fmt.Errorf("invalid func name")
	}

	var returnStr string
	for _, r := range returns {
		returnStr += fmt.Sprintf("%v,", r.GetType())
	}
	if len(returnStr) > 0 {
		returnStr = returnStr[:len(returnStr)-1]
	}

	checks := s.createChecks(check, success, tests)

	template = strings.Replace(template, "$checks$", checks, -1)
	template = strings.Replace(template, "$func$", funcName, -1)
	template = strings.Replace(template, "$userCode$", userCode, -1)
	template = strings.Replace(template, "$returns$", returnStr, -1)

	return template, nil
}

func (s *codeService) CodeFrontendTemplateGenerator(programmingName, funcName, frontend string, params []domains.Param, returns []domains.Returns, imports []string) string {
	var paramStr string
	var returnStr string
	var questImports string
	if programmingName == "GO" {
		for _, param := range params {
			paramStr += fmt.Sprintf("%v %v", param.GetName(), param.GetType())
		}
		if len(imports) > 1 {
			questImports = "import (\n"
			for _, imp := range imports {
				questImports += (`"` + imp + `"` + "\n")
			}
			questImports += ")\n"
		} else {
			questImports = "import " + `"` + imports[0] + `"` + "\n"
		}

	} else {
		for _, param := range params {
			paramStr += fmt.Sprintf("%v %v,", param.GetType(), param.GetName())
		}
		for _, imp := range imports {
			questImports += (imp + "\n")
		}
	}
	if len(paramStr) > 0 {
		paramStr = paramStr[:len(paramStr)-1]
	}

	for _, r := range returns {
		returnStr += fmt.Sprintf("%v,", r.GetType())
	}
	if len(returnStr) > 0 {
		returnStr = returnStr[:len(returnStr)-1]
	}

	frontend = strings.Replace(frontend, "$imports$", questImports, -1)
	frontend = strings.Replace(frontend, "$params$", paramStr, -1)
	frontend = strings.Replace(frontend, "$funcname$", funcName, -1)
	frontend = strings.Replace(frontend, "$returns$", returnStr, -1)

	return frontend
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

func (s *codeService) CreateFileAndWrite(filePath, content string) (err error) {
	return file.CreateFileAndWrite(filePath, content)
}
