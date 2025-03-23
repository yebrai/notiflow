# Makefile simple para notification-service
.PHONY: build run docker-build docker-up migrations

build:
	@echo "Compilando aplicación..."
	go build -o notification-service ./cmd/api

run:
	@echo "Ejecutando aplicación..."
	go run ./cmd/api

# Docker
docker-build:
	@echo "Construyendo contenedores Docker..."
	docker-compose build

docker-up:
	@echo "Iniciando contenedores Docker..."
	docker-compose up

docker-down:
	@echo "Apagando contenedores Docker..."
	docker-compose down

# Migrations
migrations:
	@echo "Ejecutando migraciones de base de datos..."
	docker exec -i notification_postgres psql -U notifuser  -d notifications < scripts/schema.sql