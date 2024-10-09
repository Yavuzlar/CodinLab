package domains

import (
	"context"
	"time"

	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/google/uuid"
)

// ILogRepository is the interface that provides the methods for the log repository.
type ILogRepository interface {
	Filter(ctx context.Context, filter LogFilter) (logs []Log, dataCount int64, err error)
	Add(ctx context.Context, log *Log) (err error)
	IsExists(ctx context.Context, log *Log) (exists bool, err error)
	CountSolutionsByDay(ctx context.Context, year string) (solutions []SolutionsByDay, err error)
	CountSolutionsByProgrammingLast7Days(ctx context.Context) (solutions []SolutionsByProgramming, err error)
}

// ILogService is the interface that provides the methods for the log service.
type ILogService interface {
	Add(ctx context.Context, userID, programmingID, labPathID, logType, content string) error
	GetAllLogs(ctx context.Context, userID, programmingID, labRoadID, logType, content string) (logs []Log, err error)
	GetByID(ctx context.Context, logID string) (log *Log, err error)
	GetByUserID(ctx context.Context, userID string) (logs []Log, err error)
	GetByType(ctx context.Context, logType string) (logs []Log, err error)
	GetByContent(ctx context.Context, content string) (logs []Log, err error)
	GetByProgrammingID(ctx context.Context, programmingID string) (logs []Log, err error)
	IsExists(ctx context.Context, userID, programmingID, labPathID, logType, content string) (isExists bool, err error)
	CountSolutionsByDay(ctx context.Context, year string) (solutions *[]SolutionsByDay, err error)
	CountSolutionsByProgrammingLast7Days(ctx context.Context) (solutions []SolutionsByProgramming, err error)
	LanguageUsageRates(ctx context.Context) (languageUsageRates []LanguageUsageRates, err error)
}

// LogFilter is the struct that represents the log filter.
type LogFilter struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	ProgrammingID int32
	LabPathID     int32 // Lab - Road title
	LType         string
	Content       string // Success etc.
}

// Log is the struct that represents the log.
type Log struct {
	id            uuid.UUID
	userId        uuid.UUID
	programmingID int32
	labPathID     int32
	logType       string
	content       string
	createdAt     time.Time
}

// lab and road numbers solved day by day
type SolutionsByDay struct {
	date  time.Time
	count int
	level int
}

func (s *SolutionsByDay) GetDate() time.Time {
	return s.date
}

func (s *SolutionsByDay) SetDate(date time.Time) {
	s.date = date
}

func (s *SolutionsByDay) GetCount() int {
	return s.count
}
func (s *SolutionsByDay) SetCount(count int) {
	s.count = count
}

func (s *SolutionsByDay) GetLevel() int {
	return s.level
}

func (s *SolutionsByDay) SetLevel(level int) {
	s.level = level
}

type SolutionsByProgramming struct {
	programmingID int32
	labCount      int
	roadCount     int
}

func (s *SolutionsByProgramming) GetProgrammingID() int32 {
	return s.programmingID
}

func (s *SolutionsByProgramming) SetProgrammingID(programmingID int32) {
	s.programmingID = programmingID
}

func (s *SolutionsByProgramming) GetLabCount() int {
	return s.labCount
}

func (s *SolutionsByProgramming) SetLabCount(labCount int) {
	s.labCount = labCount
}

func (s *SolutionsByProgramming) GetRoadCount() int {
	return s.roadCount
}

func (s *SolutionsByProgramming) SetRoadCount(roadCount int) {
	s.roadCount = roadCount
}

type LanguageUsageRates struct {
	iconPath        string
	name            string
	usagePercentage float32
}

func (s *LanguageUsageRates) GetIconPath() string {
	return s.iconPath
}

func (s *LanguageUsageRates) SetIconPath(iconPath string) {
	s.iconPath = iconPath
}
func (s *LanguageUsageRates) GetName() string { //language name
	return s.name
}

func (s *LanguageUsageRates) SetName(name string) {
	s.name = name
}
func (s *LanguageUsageRates) GetUsagePercentage() float32 {
	return s.usagePercentage
}

func (s *LanguageUsageRates) SetUsagePercentage(usagePercentage float32) {
	s.usagePercentage = usagePercentage
}

// NewLog creates a new log
func NewLog(userID, logType, content string, programmingID, labPathID int32) (*Log, error) {
	if userID == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "user id is required")
	}
	if content == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "content is required")
	}
	if logType == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "log type is required")
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid user id")
	}

	return &Log{
		id:            uuid.New(),
		userId:        userUUID,
		programmingID: programmingID,
		labPathID:     labPathID,
		logType:       logType,
		content:       content,
	}, nil
}

// Unmarshal unmarshals the log for database operations. It is used in the repository.
func (l *Log) Unmarshal(
	id, userId uuid.UUID,
	lType, content string,
	programmingID, labPathID int32,
	createdAt time.Time,
) {
	l.id = id
	l.userId = userId
	l.programmingID = programmingID
	l.labPathID = labPathID
	l.logType = lType
	l.content = content
	l.createdAt = createdAt
}

func (l *Log) ID() uuid.UUID {
	return l.id
}

func (l *Log) UserID() uuid.UUID {
	return l.userId
}

func (l *Log) ProgrammingID() int32 {
	return l.programmingID
}

func (l *Log) LabPathID() int32 {
	return l.labPathID
}

func (l *Log) Type() string {
	return l.logType
}

func (l *Log) Content() string {
	return l.content
}

func (l *Log) CreatedAt() time.Time {
	return l.createdAt
}
