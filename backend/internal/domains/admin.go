package domains

import (
	"context"

	"github.com/google/uuid"
)

// IAdminService is the interface that provides the methods for the admin service.
type IAdminService interface {
	CreateUser(ctx context.Context, username, name, surname, password, role, githubProfile string) (err error)
	GetAllUsers(ctx context.Context) (adminModelUsers []AdminUserDetail, err error)
	GetProfile(ctx context.Context, userID string) (user *User, err error)
	UpdateUser(ctx context.Context, userID, newPassword, username, githubProfile, name, surname string) (err error)
	DeleteUser(ctx context.Context, userID string) (err error)
	BestProgrammingLanguage(ctx context.Context, userID string) (bestLanguage string, err error)
}

// AdminUserDetail represents a user model in the admin panel
type AdminUserDetail struct {
	id           uuid.UUID
	username     string
	level        string
	bestLanguage string
	order        int
}

// NewAdminModelUser is a constructor function for AdminModelUser
func NewAdminUser(id uuid.UUID, order int, username, level, bestLanguage string) *AdminUserDetail {
	return &AdminUserDetail{
		id:           id,
		username:     username,
		level:        level,
		bestLanguage: bestLanguage,
		order:        order,
	}
}

func (u *AdminUserDetail) GetID() uuid.UUID {
	return u.id
}

// GetOrder returns the order of the user
func (u *AdminUserDetail) GetOrder() int {
	return u.order
}

// SetOrder sets the order of the user
func (u *AdminUserDetail) SetOrder(order int) {
	u.order = order
}

// GetUsername returns the username of the user
func (u *AdminUserDetail) GetUsername() string {
	return u.username
}

// SetUsername sets the username of the user
func (u *AdminUserDetail) SetUsername(username string) {
	u.username = username
}

// GetLevel returns the level of the user
func (u *AdminUserDetail) GetLevel() string {
	return u.level
}

// SetLevel sets the level of the user
func (u *AdminUserDetail) SetLevel(level string) {
	u.level = level
}

// GetBestLanguage returns the best language of the user
func (u *AdminUserDetail) GetBestLanguage() string {
	return u.bestLanguage
}

// SetBestLanguage sets the best language of the user
func (u *AdminUserDetail) SetBestLanguage(bestLanguage string) {
	u.bestLanguage = bestLanguage
}
