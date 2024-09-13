package dto

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/google/uuid"
)

type AdminDTOManager struct{}

// NewAdminDTOManager creates a new instance of AdminDTOManager
func NewAdminDTOManager() AdminDTOManager {
	return AdminDTOManager{}
}

type AdminUserDTO struct {
	UserID       uuid.UUID `json:"userID"`
	Order        int       `json:"order"`
	Username     string    `json:"username"`
	Level        string    `json:"level"`
	BestLanguage string    `json:"bestLanguage"`
}

type GetUsersDTO struct {
	Username      string `json:"username"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Role          string `json:"role"`
	GithubProfile string `json:"githubProfile"`
}

func (m *AdminDTOManager) ToUserProfileDTO(user *domains.User) GetUsersDTO {
	return GetUsersDTO{
		Username:      user.Username(),
		Name:          user.Name(),
		Surname:       user.Surname(),
		Role:          user.Role(),
		GithubProfile: user.GithubProfile(),
	}
}

func (m *AdminDTOManager) ToUserAdminDTO(user *domains.AdminUserDetail) AdminUserDTO {
	return AdminUserDTO{
		UserID:       user.GetID(),
		Order:        user.GetOrder(),
		Username:     user.GetUsername(),
		Level:        user.GetLevel(),
		BestLanguage: user.GetBestLanguage(),
	}
}

func (m *AdminDTOManager) ToUserAdminDTOs(users []domains.AdminUserDetail) []AdminUserDTO {
	var usersDTOs []AdminUserDTO
	for _, user := range users {
		usersDTOs = append(usersDTOs, m.ToUserAdminDTO(&user))
	}
	return usersDTOs
}

type AdminUpdateUsersDTO struct {
	Username      string `json:"username" validate:"omitempty,max=30"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Role          string `json:"role" validate:"oneof=admin user ''"`
	GithubProfile string `json:"githubProfile" validate:"omitempty,max=30"`
}
