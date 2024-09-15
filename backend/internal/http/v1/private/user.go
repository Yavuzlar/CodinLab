package private

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
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

// @Tags User
// @Summary GetProfile
// @Description Gets users profile
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=dto.UserDTO}
// @Router /private/user/ [get]
func (h *PrivateHandler) GetProfile(c *fiber.Ctx) error {
	session := session_store.GetSessionData(c)
	userID := session.UserID
	user, err := h.services.UserService.GetProfile(c.Context(), userID)
	if err != nil {
		return err
	}
	bestProgrammingLanguage, err := h.services.UserService.BestProgrammingLanguages(c.Context(), user.ID().String())
	if err != nil {
		return err
	}
	userDTO := h.dtoManager.UserDTOManager.ToUserDTO(user, bestProgrammingLanguage)

	return response.Response(200, "STATUS OK", userDTO)
}

// @Tags User
// @Summary UpdateUser
// @Description Updates user
// @Accept json
// @Produce json
// @Param update body dto.UpdateUserDTO true "UpdateUser"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/ [put]
func (h *PrivateHandler) UpdateUser(c *fiber.Ctx) error {
	var update dto.UpdateUserDTO
	if err := c.BodyParser(&update); err != nil {
		return err
	}
	session := session_store.GetSessionData(c)

	if err := h.services.UtilService.Validator().ValidateStruct(update); err != nil {
		return err
	}

	if err := h.services.UserService.UpdateUser(c.Context(), session.UserID, update.Password, update.Username, update.GithubProfile, update.Name, update.Surname); err != nil {
		return err
	}

	if err := h.services.LogService.Add(c.Context(), session.UserID, "", "", domains.TypeUser, domains.ContentProfile); err != nil {
		return err
	}

	return response.Response(200, "User successfully updated", nil)
}

// @Tags User
// @Summary UpdatePassword
// @Description Updates users password
// @Accept json
// @Produce json
// @Param update body dto.UpdatePasswordDTO true "UpdatePassword"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/password [put]
func (h *PrivateHandler) UpdatePassword(c *fiber.Ctx) error {
	session := session_store.GetSessionData(c)
	userID := session.UserID
	var update dto.UpdatePasswordDTO
	if err := c.BodyParser(&update); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(update); err != nil {
		return err
	}

	if err := h.services.UserService.UpdatePassword(c.Context(), userID, update.Password, update.NewPassword, update.ConfirmPassword); err != nil {
		return err
	}

	if err := h.services.LogService.Add(c.Context(), userID, "", "", domains.TypeUser, domains.ContentProfile); err != nil {
		return err
	}

	return response.Response(200, "Password successfully updated", nil)
}
