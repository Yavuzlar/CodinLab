package services

import (
	"context"
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/google/uuid"
)

type logService struct {
	logRepositories domains.ILogRepository
	utils           IUtilService
}

func newLogService(
	logRepositories domains.ILogRepository,
	utils IUtilService,
) domains.ILogService {
	return &logService{
		logRepositories: logRepositories,
		utils:           utils,
	}
}

func (l *logService) GetAllLogs(ctx context.Context, userID, languageID, labRoadID, logType, content string) (logs []domains.Log, err error) {
	var languageIDInt, labRoadIDInt int
	var userIDU uuid.UUID
	if userID != "" {
		userIDU, err = uuid.Parse(userID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
		}
	}

	if languageID != "" && labRoadID != "" {
		languageIDInt, err = strconv.Atoi(languageID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessage(400, "invalid language id")
		}
		labRoadIDInt, err = strconv.Atoi(labRoadID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessage(400, "invalid lab or road id")
		}
	}

	logFilter := domains.LogFilter{
		UserID:     userIDU,
		LanguageID: int32(languageIDInt),
		LType:      logType,
		LabRoadID:  int32(labRoadIDInt),
		Content:    content,
	}

	logs, _, err = l.logRepositories.Filter(ctx, logFilter)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

func (l *logService) GetByID(ctx context.Context, logID string) (log *domains.Log, err error) {
	logIDU, err := uuid.Parse(logID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid log id", err)
	}

	logs, _, err := l.logRepositories.Filter(ctx, domains.LogFilter{ID: logIDU})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}
	log = &logs[0]

	return
}

func (l *logService) GetByUserID(ctx context.Context, userID string) (logs []domains.Log, err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{UserID: userIDU})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Recives Logs with Spesific Type
func (l *logService) GetByType(ctx context.Context, logType string) (logs []domains.Log, err error) {
	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{LType: logType})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Recives Logs with Spesific Content
func (l *logService) GetByContent(ctx context.Context, content string) (logs []domains.Log, err error) {
	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{Content: content})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Recives Logs with Spesific Title
func (l *logService) GetByLabRoadID(ctx context.Context, labRoadID string) (logs []domains.Log, err error) {
	labRoadIDInt, err := strconv.Atoi(labRoadID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid lab or road id")
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{LabRoadID: int32(labRoadIDInt)})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Adds log
func (l *logService) Add(ctx context.Context, userID, ltype, content string, languageID, labRoadID int32) (err error) {
	// Creates new log
	newLog, err := domains.NewLog(userID, ltype, content, languageID, labRoadID)
	if err != nil {
		return err
	}

	// We save the new log to the database
	if err = l.logRepositories.Add(ctx, newLog); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while adding the log", err)
	}

	return
}

func (l *logService) IsExists(ctx context.Context, logID string) (isExists bool, err error) {
	logIDU, err := uuid.Parse(logID)
	if err != nil {
		return false, service_errors.NewServiceErrorWithMessageAndError(400, "invalid log id", err)
	}

	isExists, err = l.logRepositories.IsExists(ctx, &domains.LogFilter{ID: logIDU})
	if err != nil {
		return
	}

	return
}
