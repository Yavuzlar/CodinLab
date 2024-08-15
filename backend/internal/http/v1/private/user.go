package private

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initUserRoutes(root fiber.Router) {
	userRoute := root.Group("/user")
	userRoute.Get("/", h.GetProfile)
	userRoute.Put("/", h.UpdateUser)
	userRoute.Put("/password", h.UpdatePassword)
}

type UpdateUserDTO struct {
	Username      string `json:"username" validate:"omitempty,alphanum,min=3,max=30" `
	Name          string `json:"name"`
	Surname       string `json:"surname" `
	Password      string `json:"password" validate:"required"` //requires users password for update
	GithubProfile string `json:"githubProfile" validate:"omitempty,max=30"`
}

type UpdatePasswordDTO struct {
	Password        string `json:"password" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8"`
}

type UserDTO struct {
	Username      string `json:"username"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	GithubProfile string `json:"githubProfile"`
	BestLanguage  string `json:"bestLanguage"`
}

// @Tags User
// @Summary GetProfile
// @Description Gets users profile
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=UserDTO}
// @Router /private/user/ [get]
func (h *PrivateHandler) GetProfile(c *fiber.Ctx) error {
	session := session_store.GetSessionData(c)
	userID := session.UserID
	user, err := h.services.UserService.GetProfile(c.Context(), userID)
	if err != nil {
		return err
	}
	mostUsedLanguage, err := h.services.UserService.BestProgrammingLanguages(c.Context(), user.ID().String())
	if err != nil {
		return err
	}

	userDTO := UserDTO{
		Username:      user.Username(),
		Name:          user.Name(),
		Surname:       user.Surname(),
		GithubProfile: user.GithubProfile(),
		BestLanguage:  mostUsedLanguage,
	}

	return response.Response(200, "STATUS OK", userDTO)
}

// @Tags User
// @Summary UpdateUser
// @Description Updates user
// @Accept json
// @Produce json
// @Param update body UpdateUserDTO true "UpdateUser"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/ [put]
func (h *PrivateHandler) UpdateUser(c *fiber.Ctx) error {
	session := session_store.GetSessionData(c)
	userID := session.UserID
	var update UpdateUserDTO
	if err := c.BodyParser(&update); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(update); err != nil {
		return err
	}

	if err := h.services.UserService.UpdateUser(c.Context(), userID, update.Password, update.Username, update.GithubProfile, update.Name, update.Surname); err != nil {
		return err
	}

	//Update operation has been logged
	if err := h.services.LogService.Add(c.Context(), userID, domains.TypeUser, domains.ContentProfile, 0, 0); err != nil {
		return err
	}

	return response.Response(200, "User successfully updated", nil)
}

// @Tags User
// @Summary UpdatePassword
// @Description Updates users password
// @Accept json
// @Produce json
// @Param update body UpdatePasswordDTO true "UpdatePassword"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/password [put]
func (h *PrivateHandler) UpdatePassword(c *fiber.Ctx) error {
	session := session_store.GetSessionData(c)
	userID := session.UserID
	var update UpdatePasswordDTO
	if err := c.BodyParser(&update); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(update); err != nil {
		return err
	}

	if err := h.services.UserService.UpdatePassword(c.Context(), userID, update.Password, update.NewPassword, update.ConfirmPassword); err != nil {
		return err
	}

	//Update operation has been logged
	if err := h.services.LogService.Add(c.Context(), userID, domains.TypeUser, domains.ContentProfile, 0, 0); err != nil {
		return err
	}

	return response.Response(200, "Password successfully updated", nil)
}
