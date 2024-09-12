package domains

import (
	"context"
)

type ICodeService interface {
	Pull(ctx context.Context, imageReference string) (err error)
	IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error)
	UploadUserCode(ctx context.Context, userID string, programmingLanguageID, PathLabID int, codeType, fileExtention, content string) (string, error)
	RunContainerWithTar(ctx context.Context, image, tmpCodePath, fileName string, cmd []string) (string, error)
	CreateFileAndWrite(filePath, content string) (err error)
	CodeDockerTemplateGenerator(template, check, success, userCode, funcName string, tests []Test, returns []Returns) (string, error)
	CodeFrontendTemplateGenerator(programmingName, funcName, frontend string, params []Param, returns []Returns, imports []string) string
}
