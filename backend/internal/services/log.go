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

func (l *logService) GetAllLogs(ctx context.Context, userID, programmingID, labPathID, logType, content string) (logs []domains.Log, err error) {
	var programmingIDInt, labPathIDInt int
	var userIDU uuid.UUID
	if userID != "" {
		userIDU, err = uuid.Parse(userID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
		}
	}

	if programmingID != "" {
		programmingIDInt, err = strconv.Atoi(programmingID)
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
		UserID:        userIDU,
		ProgrammingID: int32(programmingIDInt),
		LType:         logType,
		LabPathID:     int32(labPathIDInt),
		Content:       content,
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
func (l *logService) GetByProgrammingID(ctx context.Context, programmingID string) (logs []domains.Log, err error) {
	programmingIDInt, err := strconv.Atoi(programmingID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid lab or road id")
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{ProgrammingID: int32(programmingIDInt)})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering logs", err)
	}

	return
}

// Adds log
func (l *logService) Add(ctx context.Context, userID, logType, content string, programmingID, labPathID int32) (err error) {
	// Creates new log
	newLog, err := domains.NewLog(userID, logType, content, programmingID, labPathID)
	if err != nil {
		return err
	}

	// We save the new log to the database
	if err = l.logRepositories.Add(ctx, newLog); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while adding the log", err)
	}

	return
}

func (l *logService) IsExists(ctx context.Context, userID, logType, content string, programmingID, labPathID int32) (isExists bool, err error) {
	log, err := domains.NewLog(userID, logType, content, programmingID, labPathID)
	if err != nil {
		return false, err
	}

	isExists, err = l.logRepositories.IsExists(ctx, log)
	if err != nil {
		return
	}

	return
}

func (l *logService) CountSolutionsByDay(ctx context.Context) (solutions []domains.SolutionsByDay, err error) {
	solutions, err = l.logRepositories.CountSolutionsByDay(ctx)
	if err != nil {
		return nil, err
	}

	return solutions, err
}

func (l *logService) CountSolutionsHoursByProgrammingLast7Days(ctx context.Context) (solutions []domains.SolutionsHoursByProgramming, err error) {
	solutions, err = l.logRepositories.CountSolutionsHoursByProgrammingLast7Days(ctx)
	if err != nil {
		return nil, err
	}

	return solutions, err
}
