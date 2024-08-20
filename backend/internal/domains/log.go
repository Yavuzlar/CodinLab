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
	CountSolutionsByDay(ctx context.Context) (solutions []SolutionsByDay, err error)
	CountSolutionsHoursByProgrammingLast7Days(ctx context.Context) (solutions []SolutionsHoursByProgramming, err error)
}

// ILogService is the interface that provides the methods for the log service.
type ILogService interface {
	Add(ctx context.Context, userID, logType, content string, programmingID, labRoadID int32) (err error)
	GetAllLogs(ctx context.Context, userID, programmingID, labRoadID, logType, content string) (logs []Log, err error)
	GetByID(ctx context.Context, logID string) (log *Log, err error)
	GetByUserID(ctx context.Context, userID string) (logs []Log, err error)
	GetByType(ctx context.Context, logType string) (logs []Log, err error)
	GetByContent(ctx context.Context, content string) (logs []Log, err error)
	GetByProgrammingID(ctx context.Context, programmingID string) (logs []Log, err error)
	IsExists(ctx context.Context, userID, logType, content string, programmingID, labPathID int32) (isExists bool, err error)
	CountSolutionsByDay(ctx context.Context) (solutions []SolutionsByDay, err error)
	CountSolutionsHoursByProgrammingLast7Days(ctx context.Context) (solutions []SolutionsHoursByProgramming, err error)
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
	Date      time.Time
	RoadCount int
	LabCount  int
}

// SolutionsHoursByProgramming represents the total hours spent on lab and road solutions for each programming language.
type SolutionsHoursByProgramming struct {
	ProgrammingID int32
	LabHours      float64
	RoadHours     float64
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

// Log Types
var (
	TypeRoad = "Road"
	TypePath = "Path"
	TypeLab  = "Lab"
	TypeUser = "User"
)

// Log Content
var (
	ContentStarted   = "Started"
	ContentCompleted = "Completed"
	ContentProfile   = "Profile Updated"
	ContentLevelUp   = "Level Up"
)
