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
	// Devamı gelecek...
}

// IUserService is the interface that provides the methods for the user service.
type IUserService interface {
	Login(ctx context.Context, username, password string) (user *User, err error)
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
	createdAt     time.Time
}

// NewUser creates a new user.
func NewUser(username, password, name, surname, role, githubProfile string, createdAt time.Time) (*User, error) {

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
		createdAt:     createdAt,
	}, nil
}

// Unmarshal unmarshals the user for database operations. It is used in the repository.
func (u *User) Unmarshal(
	id uuid.UUID,
	username, password, name, surname, role, githubProfile string,
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
