package domains

import (
	"context"
	"time"

	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/google/uuid"
)

// IUserRepository is the interface that provides the methods for the user repository.
type IUserRepository interface {
	Filter(ctx context.Context, filter UserFilter, limit, page int64) (users []User, dataCount int64, err error)
	Add(ctx context.Context, user *User) (err error)
	Update(ctx context.Context, user *User) (err error)
	Delete(ctx context.Context, userID uuid.UUID) (err error)
	// Devamı gelecek...
}

// IUserService is the interface that provides the methods for the user service.
type IUserService interface {
	Login(ctx context.Context, username, password string) (user *User, err error)
	Register(ctx context.Context, username, name, surname, password, githubProfile string) (err error)
	CreateUser(ctx context.Context, username, name, surname, password, githubProfile string) (err error)
	GetAllUsers(ctx context.Context) (users []User, err error)
	GetProfile(ctx context.Context, userID string) (user *User, err error)
	UpdateUser(ctx context.Context, userID, password, newPassword, username, githubProfile, name, surname string) (err error)
	DeleteUser(ctx context.Context, userID string) (err error)
	BestLanguage(ctx context.Context, userID string) (bestLanguage string, err error)
	// Devamı gelecek...
}

// UserFilter is the struct that represents the user filter.
type UserFilter struct {
	Id       uuid.UUID
	Username string
	Name     string
	Surname  string
	Role     string
}

// User is the struct that represents the user.
type User struct {
	id            uuid.UUID
	username      string
	password      string
	name          string
	surname       string
	role          string
	githubProfile string
	totalPoints   int32
	createdAt     time.Time
}

// NewUser creates a new user.
func NewUser(username, password, name, surname, role, githubProfile string, totalPoints int32) (*User, error) {
	if username == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "username is required")
	}
	if len(username) < 3 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "username must be at least 3 characters")
	} else if len(username) > 30 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "username must be at most 30 characters")
	}
	if password == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "password is required")
	}
	if len(password) < 8 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "password must be at least 8 characters")
	}
	if name == "" || surname == "" {
		return nil, service_errors.NewServiceErrorWithMessage(400, "name and surname are required")
	}
	if role == "" {
		role = "user"
	}
	return &User{
		id:            uuid.New(),
		username:      username,
		password:      password,
		name:          name,
		surname:       surname,
		role:          role,
		githubProfile: githubProfile,
		totalPoints:   totalPoints,
	}, nil
}

// Unmarshal unmarshals the user for database operations. It is used in the repository.
func (u *User) Unmarshal(
	id uuid.UUID,
	username, password, name, surname, role, githubProfile string,
	totalPoints int32,
	createdAt time.Time,
) {
	u.id = id
	u.username = username
	u.password = password
	u.name = name
	u.surname = surname
	u.role = role
	u.githubProfile = githubProfile
	u.createdAt = createdAt
	u.totalPoints = totalPoints
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Surname() string {
	return u.surname
}

func (u *User) Role() string {
	return u.role
}

func (u *User) GithubProfile() string {
	return u.githubProfile
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) TotalPoints() int32 {
	return u.totalPoints
}

// Setters
func (u *User) SetTotalPoints(totalPoints int32) {
	if totalPoints >= 0 {
		u.totalPoints = totalPoints
	}
}

func (u *User) SetRole(role string) {
	if role != "" {
		u.role = role
	}
}

func (u *User) SetGithubProfile(githubProfile string) {
	if githubProfile != "" {
		u.githubProfile = githubProfile
	}
}

func (u *User) SetPassword(password, hashedPassword string) error {
	if password != "" {
		if len(password) < 8 {
			return service_errors.NewServiceErrorWithMessage(400, "password must be at least 8 characters")
		}
		u.password = hashedPassword
	}
	return nil
}

func (u *User) SetName(name string) {
	if name != "" {
		u.name = name
	}
}

func (u *User) SetSurname(surname string) {
	if surname != "" {
		u.surname = surname
	}
}

func (u *User) SetUsername(username string) error {
	if username != "" {
		if len(username) < 3 {
			return service_errors.NewServiceErrorWithMessage(400, "username must be at least 3 characters")
		} else if len(username) > 30 {
			return service_errors.NewServiceErrorWithMessage(400, "username must be at most 30 characters")
		}
		u.username = username
	}
	return nil
}
func (u *User) SetID(id uuid.UUID) {
	u.id = id
}
