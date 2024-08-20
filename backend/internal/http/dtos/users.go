package dto

import "github.com/Yavuzlar/CodinLab/internal/domains"

// UserDTOManager handles the conversion of domain users to DTOs
type UserDTOManager struct{}

// NewUserDTOManager creates a new instance of UserDTOManager
func NewUserDTOManager() UserDTOManager {
	return UserDTOManager{}
}

type LoginDTO struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterDTO struct {
	Username      string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Name          string `json:"name" validate:"required"`
	Surname       string `json:"surname" validate:"required"`
	Password      string `json:"password" validate:"required,min=8"`
	GithubProfile string `json:"githubProfile" validate:"max=30"`
}

type CreateUserDTO struct {
	Username      string `json:"username" validate:"required,alphanum,min=3,max=30"` // Username is required, must be alphanumeric and between 3-30 characters
	Name          string `json:"name" validate:"required"  `                         // Name is required
	Surname       string `json:"surname" validate:"required"`                        // Surname is required
	Password      string `json:"password" validate:"required,min=8"`                 // Password is required and must be at least 8 characters
	Role          string `json:"role" validate:"required"`
	GithubProfile string `json:"githubProfile" validate:"max=30"`
}

type UserDTO struct {
	Username      string `json:"username"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	GithubProfile string `json:"githubProfile"`
	BestLanguage  string `json:"bestLanguage"`
}

func (m *UserDTOManager) ToUserDTO(user *domains.User, bestProgrammingLanguage string) UserDTO {
	return UserDTO{
		Username:      user.Username(),
		Name:          user.Name(),
		Surname:       user.Surname(),
		GithubProfile: user.GithubProfile(),
		BestLanguage:  bestProgrammingLanguage,
	}
}

type UpdateUserDTO struct {
	Username      string `json:"username" validate:"omitempty,alphanum,min=3,max=30" `
	Name          string `json:"name"`
	Surname       string `json:"surname" `
	Password      string `json:"password" validate:"required"`
	GithubProfile string `json:"githubProfile" validate:"omitempty,max=30"`
}

type UpdatePasswordDTO struct {
	Password        string `json:"password" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8"`
}
