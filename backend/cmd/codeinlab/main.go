package main

import (
	"github.com/Yavuzlar/CodinLab/internal/app"
	"github.com/Yavuzlar/CodinLab/internal/config"
)

// @title API Service
// @description API Service for CodinLab
// @version v1
// @host localhost
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name session_id
func main() {
	cfg, err := config.Init("./config")
	if err != nil {
		panic(err)
	}
	app.Run(cfg)
}
