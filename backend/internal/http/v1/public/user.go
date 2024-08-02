package public

import (
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

// User Handlers structs

type LoginDTO struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"` // Username is required, must be alphanumeric and between 3-30 characters
	Password string `json:"password" validate:"required,min=8"`                 // Password is required and must be at least 8 characters
}

type RegisterDTO struct {
	Username      string `json:"username" validate:"required,alphanum,min=3,max=30"` // Username is required, must be alphanumeric and between 3-30 characters
	Name          string `json:"name" validate:"required"`                           // Name is required
	Surname       string `json:"surname" validate:"required"`                        // Surname is required
	Password      string `json:"password" validate:"required,min=8"`                 // Password is required and must be at least 8 characters
	GithubProfile string `json:"githubProfile" validate:"max=30"`                    // Github Profile is must be max 30 characters long.
}

func (h *PublicHandler) initUserRoutes(root fiber.Router) {
	root.Post("/login", h.Login)
	root.Post("/register", h.Register)
	// initialize routes
	// Buraya yeni route'lar eklenecek lütfen Swagger'da belirtmeyi unutmayın
}

// @Tags Auth
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Param login body LoginDTO true "Login"
// @Success 200 {object} response.BaseResponse{}
// @Router /public/login [post]
func (h *PublicHandler) Login(c *fiber.Ctx) error {
	var login LoginDTO
	if err := c.BodyParser(&login); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(login); err != nil {
		return err
	}
	userdata, err := h.services.UserService.Login(c.Context(), login.Username, login.Password)
	if err != nil {
		return err
	}
	sess, err := h.session_store.Get(c)
	if err != nil {
		return err
	}
	sessionData := session_store.SessionData{}
	sessionData.ParseFromUser(userdata)
	sess.Set("user", sessionData)
	if err := sess.Save(); err != nil {
		return err
	}

	return response.Response(200, "Login successful", nil)
}

// @Tags Auth
// @Summary Register
// @Description Register
// @Accept json
// @Produce json
// @Param register body RegisterDTO true "Register"
// @Success 200 {object} response.BaseResponse{}
// @Router /public/register [post]
func (h *PublicHandler) Register(c *fiber.Ctx) error {
	var register RegisterDTO
	if err := c.BodyParser(&register); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(register); err != nil {
		return err
	}

	if err := h.services.UserService.Register(c.Context(), register.Username, register.Name, register.Surname, register.Password, register.GithubProfile); err != nil {
		return err
	}

	return response.Response(200, "Register successful", nil)
}
