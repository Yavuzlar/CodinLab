package domains

import (
	"context"

	"github.com/gofiber/websocket/v2"
)

type ICodeService interface {
	Pull(ctx context.Context, imageReference, programmingLanguage string, conn *websocket.Conn) (err error)
	IsImageExists(ctx context.Context, imageReference string) (isExists bool, err error)
	UploadUserCode(ctx context.Context, userID, programmingID, labPathID string, codeType, fileExtention, content string) (string, error)
	RunContainerWithTar(ctx context.Context, image, tmpCodePath, fileName string, cmd []string, conn *websocket.Conn) (string, error)
	CreateBashFile(cmd []string, tests []Test, userID, pathDir string) error
	CreateFileAndWrite(filePath, content string) (err error)
	CodeDockerTemplateGenerator(templatePath, funcName, userCode string, tests []Test) (string, error)
	CodeFrontendTemplateGenerator(templatePath, funcName string) (string, error)
	GetFrontendTemplate(userID, programmingID, labPathID, labRoadType string, fileExtention string) (string, error)
	DeleteFrontendTemplateHistory(userID, programmingID, labPathID, labRoadType, fileExtention string) (err error)
	StopContainer(ctx context.Context, containerID string) error
}
