package server

import (
	"context"
	"time"

	"github.com/Yavuzlar/CodinLab/internal/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	port string
}

func NewServer(cfg *config.Config, errHandler ...func(c *fiber.Ctx, err error) error) *Server {
	fiberConfig := fiber.Config{
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ServerHeader: "CodeinLab " + cfg.Application.Version,
		ProxyHeader:  cfg.HTTP.ProxyHeader,
	}
	if len(errHandler) > 0 {
		fiberConfig.ErrorHandler = errHandler[0]
	}
	return &Server{
		app:  fiber.New(fiberConfig),
		port: cfg.HTTP.Port,
	}
}

func (s *Server) Run(apiApp *fiber.App) error {
	s.app.Mount("/", apiApp)
	return s.app.Listen(":" + s.port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx2, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	return s.app.ShutdownWithContext(ctx2)
}
