// cmd/api/main.go
package main

import (
	"context"
	"log"
	"net/http"
	"notifyflow/notification-service/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Println("Archivo .env no encontrado. Usando variables de entorno del sistema.")
	}

	cfg := config.LoadConfig()

	if missingVars := cfg.Validate(); len(missingVars) > 0 {
		log.Fatalf("Variables de entorno requeridas no encontradas: %v", missingVars)
	}

	gin.SetMode(cfg.Server.GinMode)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "notification-service",
			"version": "0.1.0",
			"mode":    cfg.Server.GinMode,
		})
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		log.Printf("Iniciando servidor en http://localhost:%s en modo %s\n",
			cfg.Server.Port, cfg.Server.GinMode)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error en el shutdown del servidor: %s\n", err)
	}
	log.Println("Servidor apagado correctamente")
}
