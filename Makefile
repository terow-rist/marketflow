.PHONY: up build down logs

build:
	docker compose build
up:
	docker compose up -d

down:
	docker-compose down --volumes

logs:
	docker-compose logs
