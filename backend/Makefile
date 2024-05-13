.PHONY: help

dev.migrate.up:
	@echo "Migrating up..."
	@goose -dir migrations sqlite3 ./data.db up

dev.migrate.reset:
	@echo "Migrating down..."
	@goose -dir migrations sqlite3 ./data.db reset

dev.swagger.init:
	@echo "Generating swagger..."
	@swag init  --parseVendor  -d . -g ./cmd/codeinlab/main.go 

dev.docker.start:
	@echo "Starting docker..."
	@docker compose -f ../docker/dev.docker-compose.yaml up --build

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  dev.migrate.up      Migrate up"
	@echo "  dev.migrate.reset   Migrate down"
	@echo "  dev.swagger.init    Generate swagger"
	@echo "  help                Show this help"