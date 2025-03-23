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
	docker-compose up -d

# Migrations
migrations:
	@echo "Ejecutando migraciones de base de datos..."
	docker exec -i notification_postgres psql -U postgres -d notification_db < scripts/schema.sql