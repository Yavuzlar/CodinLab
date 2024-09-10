package domains

import (
	"context"
)

type ICodeService interface {
	Pull(ctx context.Context, imageReference string) (err error)
	IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error)
	UploadUserCode(ctx context.Context, userID string, programmingLanguageID, PathLabID int, codeType, fileExtention, content string) (string, error)
	RunContainerWithTar(ctx context.Context, image, tmpCodePath string, cmd []string) (string, error)
	CreateFileAndWrite(filePath, content string) (err error)
	CodeTemplateGenerator(programmingName, templatePathObject, content, funcName string, tests []Test) (string, error)
}
