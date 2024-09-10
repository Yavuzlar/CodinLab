package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	extractor "github.com/Yavuzlar/CodinLab/pkg/code_extractor"
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

// // Bu kısımda bütün diller için template oluşturma kısmı gelicek.
// func (s *codeService) CodeTemplateGenerator(programmingName, templatePathObject, content, funcName string, tests []domains.Test) (string, error) {
// 	if programmingName == "GO" {
// 		return s.goLabTemplate(templatePathObject, content, funcName, tests)
// 	}

// 	return "", service_errors.NewServiceErrorWithMessage(500, "this programming language not supported")
// }

func (s *codeService) CodeTemplateGenerator(template, check, userCode, funcName string, tests []domains.Test) (string, error) {
	if !strings.Contains(userCode, funcName) {
		return "", fmt.Errorf("invalid func name")
	}

	checks := s.createChecks(check, tests)
	template = strings.Replace(template, "$checks$", checks, -1)
	template = strings.Replace(template, "$func$", funcName, -1)
	template = strings.Replace(template, "$userCode$", userCode, -1)

	return template, nil
}

func (s *codeService) createChecks(check string, tests []domains.Test) string {
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
		for _, out := range test.GetOutput() {
			switch v := out.(type) {
			case string:
				// Add double quotes around string outputs
				outputs = append(outputs, `"`+v+`"`)
			default:
				// Directly use other types (int, etc.)
				outputs = append(outputs, fmt.Sprintf("%v", v))
			}
		}
		tmp = strings.Replace(tmp, "$output$", strings.Join(outputs, ", "), -1)

		checks.WriteString(tmp + "\n")
	}

	return checks.String()
}

func (s *codeService) goLabTemplate(templatePathObject, content, funcName string, tests []domains.Test) (string, error) {
	// Read the template file
	temp, err := os.ReadFile(templatePathObject)
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "error while reading go template", err)
	}

	// Replace placeholders with actual function names and imports
	replace := strings.Replace(string(temp), "#funccall", funcName, -1)
	imports := extractor.ExtractImports(content)
	replace = strings.Replace(replace, "#imports", imports, -1)

	// Extract the user's function from the content
	userfunc, err := extractor.ExtractFunction(content, funcName)
	if err != nil {
		return "", err
	}
	replace = strings.Replace(replace, "#funcs", userfunc, -1)

	// Build the test cases
	result := "var tests = []struct{\n input []interface{}\n output []interface{}\n}{\n"

	for _, test := range tests {
		result += "\t{input:[]interface{}{"
		for i, input := range test.GetInput() {

			switch input.(type) { //checking if string or not
			case string:
				result += fmt.Sprintf("\"%v\"", input)
			default:
				result += fmt.Sprintf("%v", input)
			}

			if len(test.GetInput()) != i+1 {
				result += ","
			}
		}
		result += "}, output:[]interface{}{"
		for i, output := range test.GetOutput() {

			switch output.(type) {
			case string:
				result += fmt.Sprintf("\"%v\"", output)
			default:
				result += fmt.Sprintf("%v", output)
			}

			if len(test.GetOutput()) != i+1 {
				result += ","
			}
		}
		result += "}},\n"
	}
	result += "}"

	// Replace the test cases placeholder in the template
	replace = strings.Replace(replace, "#tests", result, -1)

	return replace, nil
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
