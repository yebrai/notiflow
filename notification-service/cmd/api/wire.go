//go:build wireinject
// +build wireinject

package main

import (
	"notification-service/internal/config"
	"notification-service/internal/infrastructure/api"
	"notification-service/internal/infrastructure/database"

	"github.com/google/wire"
)

// InitializeServer configura todas las dependencias y retorna una instancia de Server lista para usar
func InitializeServer(cfg config.Config) (*api.Server, error) {
	wire.Build(
		// Database
		database.NewGormDB,
		wire.FieldsOf(&cfg, "Database"),

		// Proporcionar la configuración del servidor
		wire.FieldsOf(&cfg, "Server"),

		// Crear el servidor
		api.NewServer,
	)

	return nil, nil
}
