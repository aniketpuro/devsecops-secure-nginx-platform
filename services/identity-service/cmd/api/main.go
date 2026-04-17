package main

import (
	"log"

	"github.com/gin-gonic/gin"
	
	"secure-identity-service/internal/config"
	"secure-identity-service/internal/handler"
	"secure-identity-service/internal/middleware"
)

func main() {
	cfg := config.Load()
	authHandler := handler.NewAuthHandler(cfg)

	router := gin.Default()

	// PUBLIC ROUTES
	router.GET("/health", handler.HealthCheck)
	router.GET("/ready", handler.ReadyCheck)
	
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// PROTECTED ROUTES
	protected := router.Group("/api")
	protected.Use(middleware.RequireAuth(cfg.JWTSecret)) 
	{
		protected.POST("/convert", handler.ConvertToMP3) 
	}

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting secure-identity-service on port %s", port)
	router.Run(":" + port)
}