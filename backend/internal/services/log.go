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

func (l *logService) GetAllLogs(ctx context.Context, userID, languageID, labPathID, logType, content string) (logs []domains.Log, err error) {
	var languageIDInt, labPathIDInt int
	var userIDU uuid.UUID
	if userID != "" {
		userIDU, err = uuid.Parse(userID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
		}
	}

	if languageID != "" {
		languageIDInt, err = strconv.Atoi(languageID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessage(400, "invalid language id")
		}
	}

	if labPathID != "" {
		labPathIDInt, err = strconv.Atoi(labPathID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessage(400, "invalid lab or path id")
		}
	}

	logFilter := domains.LogFilter{
		UserID:     userIDU,
		LanguageID: int32(languageIDInt),
		LType:      logType,
		LabPathID:  int32(labPathIDInt),
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
func (l *logService) GetByLanguageID(ctx context.Context, languageID string) (logs []domains.Log, err error) {
	labPathIDInt, err := strconv.Atoi(languageID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid lab or road id")
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{LanguageID: int32(labPathIDInt)})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Adds log
func (l *logService) Add(ctx context.Context, userID, ltype, content string, languageID, labPathID int32) (err error) {
	// Creates new log
	newLog, err := domains.NewLog(userID, ltype, content, languageID, labPathID)
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

// author: yasir
func (l *logService) CountSolutionsByDay(ctx context.Context) (solutions []domains.SolutionsByDay, err error) {
	solutions, err = l.logRepositories.CountSolutionsByDay(ctx)
	if err != nil {
		return nil, err
	}

	return solutions, err
}

// author: yasir
func (l *logService) CountSolutionsHoursByLanguageLast7Days(ctx context.Context) (solutions []domains.SolutionsHoursByLanguage, err error) {
	solutions, err = l.logRepositories.CountSolutionsHoursByLanguageLast7Days(ctx)
	if err != nil {
		return nil, err
	}

	return solutions, err
}
