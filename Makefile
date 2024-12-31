.PHONY: help

dev:
	@echo "[i] Project is starting in development mode...\n"
	@if command -v docker-compose > /dev/null; then \
		echo "[i] Using docker-compose..."; \
		docker-compose -f docker/dev.docker-compose.yaml -p codinlab up -d; \
	else \
		echo "[i] Using docker compose..."; \
		docker compose -f docker/dev.docker-compose.yaml -p codinlab up -d; \
	fi
	@echo "\n[+] Project is started in development mode..."
	
down:
	@echo "[i] Stopping and removing containers...\n"
	@if command -v docker-compose > /dev/null; then \
		echo "[i] Using docker-compose..."; \
		docker-compose -f docker/dev.docker-compose.yaml -p codinlab down; \
	else \
		echo "[i] Using docker compose..."; \
		docker compose -f docker/dev.docker-compose.yaml -p codinlab down; \
	fi
	@echo "\n[+] Project is stopped and containers are removed..."

dev-build:
	@echo "[i] Production mode is not active for now.\n"
	@if command -v docker-compose > /dev/null; then \
		echo "[i] Using docker-compose..."; \
		docker-compose -f docker/dev.docker-compose.yaml -p codinlab build; \
	else \
		echo "[i] Using docker compose..."; \
		docker compose -f docker/dev.docker-compose.yaml -p codinlab build; \
	fi

help:
	@echo "[i] Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  dev      	start project in development mode"
	@echo "  down     	stop and remove containers"
	@echo "  build    	build project in production mode"
	@echo "  dev-build	build project in development mode"
	@echo "  help		Show this help"
