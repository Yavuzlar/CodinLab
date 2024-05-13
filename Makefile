.PHONY: help

dev:
	@echo "[i] Project is starting in development mode...\n"
	@docker-compose -f docker/dev.docker-compose.yaml -p codinlab up -d
	@echo "\n[+] Project is started in development mode..."

build:
	@echo "[i] Production mode is not active for now.\n"

help:
	@echo "[i] Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  dev      	start project in development mode"
	@echo "  build    	build project in production mode"
	@echo "  help		Show this help"