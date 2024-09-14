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
	CodeDockerTemplateGenerator(templatePath, funcName, userCode string, tests []Test) (string, error)
	CodeFrontendTemplateGenerator(templatePath, funcName string) (string, error)
	GetFrontendTemplate(userID, labRoadType string, programmingID, labPathID int, fileExtention string) (string, error)
}
