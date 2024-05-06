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

func (h *PublicHandler) initUserRoutes(root fiber.Router) {
	root.Post("/login", h.Login)
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
