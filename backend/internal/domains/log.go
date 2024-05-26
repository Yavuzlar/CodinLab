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
	IsExists(ctx context.Context, log *LogFilter) (exists bool, err error)
}

// ILogService is the interface that provides the methods for the log service.
type ILogService interface {
	Add(ctx context.Context, userID, ltype, content string, languageID, labRoadID int32) (err error)
	GetAllLogs(ctx context.Context, userID, languageID, labRoadID, logType, content string) (logs []Log, err error)
	GetByID(ctx context.Context, logID string) (log *Log, err error)
	GetByUserID(ctx context.Context, userID string) (logs []Log, err error)
	GetByType(ctx context.Context, logType string) (logs []Log, err error)
	GetByContent(ctx context.Context, content string) (logs []Log, err error)
	GetByLabRoadID(ctx context.Context, labRoadID string) (logs []Log, err error)
	IsExists(ctx context.Context, logID string) (isExists bool, err error)
}

// LogFilter is the struct that represents the log filter.
type LogFilter struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	LanguageID int32
	LabRoadID  int32 // Lab - Road title
	LType      string
	Content    string // Success etc.
}

// Log is the struct that represents the log.
type Log struct {
	id         uuid.UUID
	userId     uuid.UUID
	languageID int32
	labRoadID  int32
	lType      string
	content    string
	createdAt  time.Time
}

// NewLog creates a new log
func NewLog(userID, lType, content string, languageID, labRoadID int32) (*Log, error) {
	if userID == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "user id is required")
	}
	if content == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "content is required")
	}
	if lType == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "log type is required")
	}
	if languageID == 0 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "log language is required")
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid user id")
	}

	return &Log{
		id:         uuid.New(),
		userId:     userUUID,
		languageID: languageID,
		labRoadID:  labRoadID,
		lType:      lType,
		content:    content,
	}, nil
}

// Unmarshal unmarshals the log for database operations. It is used in the repository.
func (l *Log) Unmarshal(
	id, userId uuid.UUID,
	lType, content string,
	languageID, labRoadID int32,
	createdAt time.Time,
) {
	l.id = id
	l.userId = userId
	l.languageID = languageID
	l.labRoadID = labRoadID
	l.lType = lType
	l.content = content
	l.createdAt = createdAt
}

func (l *Log) ID() uuid.UUID {
	return l.id
}

func (l *Log) UserID() uuid.UUID {
	return l.userId
}

func (l *Log) LanguageID() int32 {
	return l.languageID
}

func (l *Log) LabRoadID() int32 {
	return l.labRoadID
}

func (l *Log) Type() string {
	return l.lType
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
)
