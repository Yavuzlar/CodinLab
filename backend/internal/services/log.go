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
	parserService   domains.IParserService
	utils           IUtilService
}

func newLogService(
	logRepositories domains.ILogRepository,
	utils IUtilService,
	parserService domains.IParserService,
) domains.ILogService {
	return &logService{
		logRepositories: logRepositories,
		utils:           utils,
		parserService:   parserService,
	}
}

func (l *logService) GetAllLogs(ctx context.Context, userID, programmingID, labPathID, logType, content string) (logs []domains.Log, err error) {
	var programmingIDInt, labPathIDInt int
	var userIDU uuid.UUID
	if userID != "" {
		userIDU, err = uuid.Parse(userID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessageAndError(400, domains.ErrInvalidUserID, err)
		}
	}

	if programmingID != "" {
		programmingIDInt, err = strconv.Atoi(programmingID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
		}
	}

	if labPathID != "" {
		labPathIDInt, err = strconv.Atoi(labPathID)
		if err != nil {
			return nil, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidLabOrPathID)
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
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrFilteringLogs, err)
	}

	return
}

func (l *logService) GetByID(ctx context.Context, logID string) (log *domains.Log, err error) {
	logIDU, err := uuid.Parse(logID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, domains.ErrInvalidLogID, err)
	}

	logs, _, err := l.logRepositories.Filter(ctx, domains.LogFilter{ID: logIDU})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrFilteringLogs, err)
	}
	log = &logs[0]

	return
}

func (l *logService) GetByUserID(ctx context.Context, userID string) (logs []domains.Log, err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, domains.ErrInvalidUserID, err)
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{UserID: userIDU})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrFilteringLogs, err)
	}

	return
}

// Recives Logs with Spesific Type
func (l *logService) GetByType(ctx context.Context, logType string) (logs []domains.Log, err error) {
	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{LType: logType})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrFilteringLogs, err)
	}

	return
}

// Recives Logs with Spesific Content
func (l *logService) GetByContent(ctx context.Context, content string) (logs []domains.Log, err error) {
	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{Content: content})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrFilteringLogs, err)
	}

	return
}

// Recives Logs with Spesific Title
func (l *logService) GetByProgrammingID(ctx context.Context, programmingID string) (logs []domains.Log, err error) {
	programmingIDInt, err := strconv.Atoi(programmingID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidLabOrPathID)
	}

	logs, _, err = l.logRepositories.Filter(ctx, domains.LogFilter{ProgrammingID: int32(programmingIDInt)})
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrFilteringLogs, err)
	}

	return
}

// Adds log
func (l *logService) Add(ctx context.Context, userID, programmingID, labPathID, logType, content string) error {
	// Error Control
	var err error
	var intProgrammingID, intLabPathID int
	if programmingID != "" {
		intProgrammingID, err = strconv.Atoi(programmingID)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
		}
	}

	if labPathID != "" {
		intLabPathID, err = strconv.Atoi(labPathID)
		if err != nil {
			return service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidLabOrPathID)
		}
	}

	// Creates new log
	newLog, err := domains.NewLog(userID, logType, content, int32(intProgrammingID), int32(intLabPathID))
	if err != nil {
		return err
	}

	// We save the new log to the database
	if err = l.logRepositories.Add(ctx, newLog); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, domains.ErrAddingLog, err)
	}

	return nil
}

func (l *logService) IsExists(ctx context.Context, userID, programmingID, labPathID, logType, content string) (isExists bool, err error) {
	var intProgrammingID, intLabPathID int

	if programmingID != "" {
		intProgrammingID, err = strconv.Atoi(programmingID)
		if err != nil {
			return false, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
		}
	}
	if labPathID != "" {
		intLabPathID, err = strconv.Atoi(labPathID)
		if err != nil {
			return false, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
		}
	}

	log, err := domains.NewLog(userID, logType, content, int32(intProgrammingID), int32(intLabPathID))
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
func (s *logService) LanguageUsageRates(ctx context.Context) (languageUsageRates []domains.LanguageUsageRates, err error) {
	var rate int
	programmingLanguages, _ := s.parserService.GetInventory()

	roadLogs, err := s.GetByType(ctx, domains.TypePath)
	if err != nil {
		return nil, err
	}

	labLogs, err := s.GetByType(ctx, domains.TypeLab)
	if err != nil {
		return nil, err
	}

	roads, err := s.parserService.GetRoads()
	if err != nil {
		return nil, err
	}

	labs, err := s.parserService.GetLabs()
	if err != nil {
		return nil, err
	}

	total := len(roads) + len(labs)

	var rates domains.LanguageUsageRates
	for _, pl := range programmingLanguages {
		rates.SetIconPath(pl.IconPath)
		rates.SetName(pl.Name)
		for _, path := range roadLogs {
			if pl.ID == int(path.ProgrammingID()) {
				if path.Content() == domains.ContentStarted {
					rate++
				}
			}
		}
		for _, labs := range labLogs {
			if pl.ID == int(labs.ProgrammingID()) {
				if labs.Content() == domains.ContentStarted {
					rate++
				}
			}
		}
		totalUsage := float32(float32(rate)/float32(total)) * 100
		rates.SetUsagePercentage(totalUsage)
		languageUsageRates = append(languageUsageRates, rates)
	}
	return
}
