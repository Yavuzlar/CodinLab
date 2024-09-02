package private

import (
	_ "github.com/Yavuzlar/CodinLab/internal/domains"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initAdminRoutes(root fiber.Router) {
	adminRoute := root.Group("/admin")
	adminRoute.Use(h.adminAuthMiddleware)
	adminRoute.Get("/user/:ID", h.GetUserProfile)
	adminRoute.Put("/user/:ID", h.UpdateUser)
	adminRoute.Get("/user", h.GetAllUsers)
	adminRoute.Post("/user", h.CreateUser)
	adminRoute.Post("/user/:userID", h.UpdateUserAdmin)
}

// @Tags Admin
// @Summary Get Profile
// @Description Retrieves User Profile
// @Accept json
// @Produce json
// @Param ID path string true "User ID"
// @Success 200 {object} response.BaseResponse{data=dto.UserDTO}
// @Failure 400 {object} response.BaseResponse
// @Router /private/admin/user/{ID} [get]
func (h *PrivateHandler) GetUserProfile(c *fiber.Ctx) error {
	userID := c.Params("ID")
	user, err := h.services.AdminService.GetProfile(c.Context(), userID)
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

// @Tags Admin
// @Summary Get Users
// @Description Retrieves All Users
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Failure 400 {object} response.BaseResponse
// @Router /private/admin/user [get]
func (h *PrivateHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.services.AdminService.GetAllUsers(c.Context())
	if err != nil {
		return err
	}
	userDTOs := h.dtoManager.AdminDTOManager.ToUserAdminDTOs(users)

	return response.Response(200, "STATUS OK", userDTOs)
}

// @Tags Admin
// @Summary Creates User
// @Description Creates User
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "User"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/user [post]
func (h *PrivateHandler) CreateUser(c *fiber.Ctx) error {
	var user dto.CreateUserDTO
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(user); err != nil {
		return err
	}
	if err := h.services.AdminService.CreateUser(c.Context(), user.Username, user.Name, user.Surname, user.Password, user.Role, user.GithubProfile); err != nil {
		return err
	}

	return response.Response(200, "Register successful", nil)
}

// @Tags Admin
// @Summary Updates User
// @Description Updates User
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param user body dto.AdminUpdateUsersDTO true "New User Creds"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/user/{userID} [post]
func (h *PrivateHandler) UpdateUserAdmin(c *fiber.Ctx) error {
	userID := c.Params("userID")
	var user dto.AdminUpdateUsersDTO
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(user); err != nil {
		return err
	}

	if err := h.services.AdminService.UpdateUser(c.Context(), userID, user.Password, user.Username, user.GithubProfile, user.Name, user.Surname); err != nil {
		return err
	}

	return response.Response(200, "Update successful", nil)
}
