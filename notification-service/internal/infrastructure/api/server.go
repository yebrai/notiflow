package api

import (
	"context"
	"log"
	"net/http"
	"notification-service/internal/config"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	config     config.ServerConfig
	httpServer *http.Server
	db         *gorm.DB
}

func NewServer(config config.ServerConfig, db *gorm.DB) *Server {
	gin.SetMode(config.GinMode)
	router := gin.Default()

	return &Server{
		router: router,
		config: config,
		db:     db,
		httpServer: &http.Server{
			Addr:         ":" + config.Port,
			Handler:      router,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
			IdleTimeout:  config.IdleTimeout,
		},
	}
}

func (s *Server) SetupRoutes() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "notification-service",
			"version": "0.1.0",
			"mode":    s.config.GinMode,
		})
	})

}

func (s *Server) Start() error {
	s.SetupRoutes()

	go func() {
		log.Printf("Iniciando servidor en http://localhost:%s en modo %s\n",
			s.config.Port, s.config.GinMode)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor: %s\n", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
