package public

import (
	"sync"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type PublicHandler struct {
	services      *services.Services
	session_store *session.Store
	dtoManager    *dto.DTOManager
	clients       map[*domains.Client]bool
	mu            sync.Mutex
}

func NewPublicHandler(
	service *services.Services,
	sessionStore *session.Store,
	dtoManager *dto.DTOManager,
	clients map[*domains.Client]bool,

) *PublicHandler {
	return &PublicHandler{
		services:      service,
		session_store: sessionStore,
		dtoManager:    dtoManager,
		clients:       clients,
	}
}

func (h *PublicHandler) Init(router fiber.Router) {
	root := router.Group("/public")

	root.Get("/", func(c *fiber.Ctx) error {
		return response.Response(200, "Welcome to CodinLab API (Public Zone)", nil)
	})
	// initialize routes
	h.initUserRoutes(root)
}
