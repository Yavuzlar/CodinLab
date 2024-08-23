package v1

import (
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/v1/private"
	"github.com/Yavuzlar/CodinLab/internal/http/v1/public"
	"github.com/Yavuzlar/CodinLab/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type V1Handler struct {
	services   *services.Services
	dtoManager *dto.DTOManager
}

func NewV1Handler(services *services.Services, dtoManager *dto.DTOManager) *V1Handler {
	return &V1Handler{
		services:   services,
		dtoManager: dtoManager,
	}
}

func (h *V1Handler) Init(router fiber.Router, sessionStore *session.Store) {
	root := router.Group("/v1")
	root.Get("/", func(c *fiber.Ctx) error {
		return response.Response(200, "Welcome to CodinLab API (Root Zone)", nil)
	})
	// Init Fiber Session Store
	//---------------------------
	private.NewPrivateHandler(h.services, sessionStore, h.dtoManager).Init(root)
	public.NewPublicHandler(h.services, sessionStore, h.dtoManager).Init(root)
}
