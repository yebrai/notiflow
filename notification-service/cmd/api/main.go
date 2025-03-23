// cmd/api/main.go
package main

import (
	"context"
	"log"
	"notification-service/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Cargar configuración
	if err := config.LoadEnv(); err != nil {
		log.Println("Archivo .env no encontrado. Usando variables de entorno del sistema.")
	}

	cfg := config.LoadConfig()

	// Validar configuración
	if missingVars := cfg.Validate(); len(missingVars) > 0 {
		log.Fatalf("Variables de entorno requeridas no encontradas: %v", missingVars)
	}

	// Inicializar servidor usando Wire
	server, err := InitializeServer(cfg)
	if err != nil {
		log.Fatalf("Error al inicializar el servidor: %v", err)
	}

	// Iniciar servidor
	if err := server.Start(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	// Configurar graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error en el shutdown del servidor: %v", err)
	}

	log.Println("Servidor apagado correctamente")
}
