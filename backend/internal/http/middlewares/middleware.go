package middlewares

import (
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/config"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func InitMiddlewares(cfg *config.Config) (mws []func(*fiber.Ctx) error) {
	// Fiber storage bağlantısı yapacağım unutma
	cors := cors.New(
		cors.Config{
			AllowOrigins:     strings.Join(cfg.HTTP.AllowedOrigins, ","),
			AllowMethods:     strings.Join(cfg.HTTP.AllowedMethods, ","),
			AllowHeaders:     strings.Join(cfg.HTTP.AllowedHeaders, ","),
			AllowCredentials: cfg.HTTP.AllowCredentials,
			ExposeHeaders:    strings.Join(cfg.HTTP.ExposedHeaders, ","),
		},
	)

	helmetMid := helmet.New(helmet.ConfigDefault) // helmet configleri yazılacak

	mws = append(mws, cors, helmetMid)

	if !cfg.Application.DevMode {
		limitter := limiter.New(limiter.Config{
			Max: 50,
			LimitReached: func(c *fiber.Ctx) error {
				return service_errors.NewServiceErrorWithMessage(429, "too many requests")
			},
		})
		mws = append(mws, limitter)
	}
	return
}
