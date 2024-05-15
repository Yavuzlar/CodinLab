package private

import (
	"fmt"

	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/Yavuzlar/CodinLab/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type PrivateHandler struct {
	services   *services.Services
	sess_store *session.Store
}

func NewPrivateHandler(
	service *services.Services,
	sessStore *session.Store) *PrivateHandler {
	return &PrivateHandler{
		services:   service,
		sess_store: sessStore,
	}
}

func (h *PrivateHandler) Init(router fiber.Router) {
	root := router.Group("/private")
	root.Use(h.authMiddleware)

	root.Get("/", func(c *fiber.Ctx) error {
		data := session_store.GetSessionData(c)
		return response.Response(200, fmt.Sprintf("Dear %s %s Welcome to CodinLab API (Private Zone)", data.Name, data.Surname), nil)
	})
	h.initRoadRoutes(root)
	// initialize routes
	// Buraya yeni route'lar eklenecek

}

func (h *PrivateHandler) authMiddleware(c *fiber.Ctx) error {
	session, err := h.sess_store.Get(c)
	if err != nil {
		return err
	}
	user := session.Get("user")
	if user == nil {
		return service_errors.NewServiceErrorWithMessage(401, "unauthorized")
	}
	session_data, ok := user.(session_store.SessionData)
	if !ok {
		return service_errors.NewServiceErrorWithMessage(500, "session data error")
	}
	if session_data.Role == "banned" {
		return service_errors.NewServiceErrorWithMessage(403, "banned")
	}
	c.Locals("user", session_data)
	return c.Next()
}
