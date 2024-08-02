package private

import (
	_ "github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *PrivateHandler) initAdminRoutes(root fiber.Router) {
	adminRoute := root.Group("/admin")
	//adminRoute.Use(h.adminAuthMiddleware)
	adminRoute.Get("/user/:id", h.GetUserProfile)
	adminRoute.Put("/user/:id", h.UpdateUser)
	adminRoute.Get("/users", h.GetAllUsers)
	adminRoute.Post("/user", h.CreateUser)
}

type GetUserDTO struct {
	Username      string    `json:"username"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	GithubProfile string    `json:"githubProfile"`
	BestLanguage  string    `json:"bestLanguage"`
	ID            uuid.UUID `json:"id"`
}

type CreateUserDTO struct {
	Username      string `json:"username" validate:"required,alphanum,min=3,max=30"` // Username is required, must be alphanumeric and between 3-30 characters
	Name          string `json:"name" validate:"required"  `                         // Name is required
	Surname       string `json:"surname" validate:"required"`                        // Surname is required
	Password      string `json:"password" validate:"required,min=8"`                 // Password is required and must be at least 8 characters
	Role          string `json:"role" validate:"required"`
	GithubProfile string `json:"githubProfile" validate:"max=30"` // Github Profile is must be max 30 characters long.
}

// GetUserProfile retrieves a user profile by ID
// @Tags Admin
// @Summary GetProfile
// @Description Gets user profile
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.BaseResponse{data=UserDTO}
// @Failure 400 {object} response.BaseResponse
// @Router /private/admin/user/{id} [get]
func (h *PrivateHandler) GetUserProfile(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := h.services.AdminService.GetProfile(c.Context(), userID)
	if err != nil {
		return err
	}
	//best language hangi serviceten alinacak?
	mostUsedLanguage, err := h.services.UserService.BestLanguage(c.Context(), user.ID().String())
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

// GetAllUsers retrieves all users
// @Tags Admin
// @Summary Get All Users
// @Description Retrieves all users
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=[]GetUserDTO}
// @Failure 400 {object} response.BaseResponse
// @Router /private/admin/users [get]
func (h *PrivateHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.services.AdminService.GetAllUsers(c.Context())
	if err != nil {
		return err
	}

	userDTOs := make([]UserDTO, len(users))
	for i, user := range users {
		//best language atama işlemini service katmanında mı yapsak?
		//adminModelUser structı olusturmak gerek domainsde
		mostUsedLanguage, err := h.services.AdminService.BestLanguage(c.Context(), user.ID().String())
		if err != nil {
			return err
		}

		userDTOs[i] = UserDTO{
			Username:      user.Username(),
			Name:          user.Name(),
			Surname:       user.Surname(),
			GithubProfile: user.GithubProfile(),
			BestLanguage:  mostUsedLanguage,
		}
	}

	return response.Response(200, "STATUS OK", userDTOs)
}

// @Tags Admin
// @Summary Creates User
// @Description User Creation
// @Accept json
// @Produce json
// @Param user body CreateUserDTO true "User"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/user [post]
func (h *PrivateHandler) CreateUser(c *fiber.Ctx) error {
	var user CreateUserDTO
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(user); err != nil {
		return err
	}
	// direkt user göndersek daha mantıklı değil mi?
	if err := h.services.AdminService.CreateUser(c.Context(), user.Username, user.Name, user.Surname, user.Password, user.Role, user.GithubProfile); err != nil {
		return err
	}

	return response.Response(200, "Register successful", nil)
}
