package services

import (
	"context"

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

func (l *logService) GetByUserID(ctx context.Context, userID string) (logs []domains.Log, err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{UserId: userIDU})
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
func (l *logService) GetByTitle(ctx context.Context, title string) (logs []domains.Log, err error) {
	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{Title: title})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Adds log
func (l *logService) Add(ctx context.Context, userID, title, ltype, content string) (err error) {
	// Creates new log
	newLog, err := domains.NewLog(userID, title, ltype, content)
	if err != nil {
		return err
	}

	// We save the new log to the database
	if err = l.logRepositories.Add(ctx, newLog); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while adding the log", err)
	}

	return
}
