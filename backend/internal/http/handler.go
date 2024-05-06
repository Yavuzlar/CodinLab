package http

import (
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	v1 "github.com/Yavuzlar/CodinLab/internal/http/v1"
	"github.com/Yavuzlar/CodinLab/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(devMode bool, middlewares ...func(*fiber.Ctx) error) *fiber.App {
	app := fiber.New()
	//init middlewares
	for i := range middlewares {
		app.Use(middlewares[i])
	}
	if devMode {
		app.Static("/docs/", "./docs")                // default // air ile çalışıldığından ./docs/ olarak değiştirildi.
		app.Get("/dev/*", swagger.New(swagger.Config{ // custom
			URL:          "/docs/swagger.yaml",
			DocExpansion: "none",
		}))
	}

	root := app.Group("/api")
	// init routes
	sessionStore := session_store.NewSessionStore()

	v1.NewV1Handler(h.services).Init(root, sessionStore)

	return app
}
