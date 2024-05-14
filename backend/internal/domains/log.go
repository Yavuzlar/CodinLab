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
}

// ILogService is the interface that provides the methods for the log service.
type ILogService interface {
	Add(ctx context.Context, userId, title, ltype, content string) (err error)
	GetByUserID(ctx context.Context, userID string) (logs []Log, err error)
	GetByType(ctx context.Context, logType string) (logs []Log, err error)
	GetByContent(ctx context.Context, content string) (logs []Log, err error)
	GetByTitle(ctx context.Context, title string) (logs []Log, err error)
}

// LogFilter is the struct that represents the log filter.
type LogFilter struct {
	Id      uuid.UUID
	UserId  uuid.UUID
	Title   string // Lab - Road title
	LType   string
	Content string // Success etc.
}

// Log is the struct that represents the log.
type Log struct {
	id        uuid.UUID
	userId    uuid.UUID
	title     string
	lType     string
	content   string
	createdAt time.Time
}

// NewLog creates a new log
func NewLog(userId, title, lType, content string) (*Log, error) {
	if userId == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "user id is required")
	}
	if content == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "content is required")
	}
	if lType == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "log type is required")
	}
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid user id")
	}

	return &Log{
		id:      uuid.New(),
		userId:  userUUID,
		title:   title,
		lType:   lType,
		content: content,
	}, nil
}

// Unmarshal unmarshals the log for database operations. It is used in the repository.
func (l *Log) Unmarshal(
	id, userId uuid.UUID,
	title, lType, content string,
	createdAt time.Time,
) {
	l.id = id
	l.userId = userId
	l.title = title
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

func (l *Log) Title() string {
	return l.title
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
	Started   = "Started"
	Completed = "Completed"
	Profile   = "Profile Updated"
)
