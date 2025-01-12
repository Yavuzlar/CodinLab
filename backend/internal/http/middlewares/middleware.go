package middlewares

import (
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func InitMiddlewares(cfg *config.Config) (mws []func(*fiber.Ctx) error) {
	// Don't forget that I will make a fiber storage connection.
	cors := cors.New(
		cors.Config{
			AllowOrigins:     strings.Join(cfg.HTTP.AllowedOrigins, ","),
			AllowMethods:     strings.Join(cfg.HTTP.AllowedMethods, ","),
			AllowHeaders:     strings.Join(cfg.HTTP.AllowedHeaders, ","),
			AllowCredentials: cfg.HTTP.AllowCredentials,
			ExposeHeaders:    strings.Join(cfg.HTTP.ExposedHeaders, ","),
		},
	)

	helmetMid := helmet.New(helmet.ConfigDefault) // helmet configurations will be written

	mws = append(mws, cors, helmetMid)

	// Limitter Stopped
	// if !cfg.Application.DevMode {
	// 	limitter := limiter.New(limiter.Config{
	// 		Max: 50,
	// 		LimitReached: func(c *fiber.Ctx) error {
	// 			return service_errors.NewServiceErrorWithMessage(429, "too many requests")
	// 		},
	// 	})
	// 	mws = append(mws, limitter)
	// }
	return
}
