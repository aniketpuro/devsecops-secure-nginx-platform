package main

import (
	"log"

	"github.com/gin-gonic/gin"
	
	// Aapke internal packages zaroor import hone chahiye
	"secure-identity-service/internal/config"
	"secure-identity-service/internal/handler"
)

func main() {
	// 1. Load config
	cfg := config.Load()

	// 2. Initialize AuthHandler
	authHandler := handler.NewAuthHandler(cfg)

	// 3. Setup Gin router
	router := gin.Default()

	// 4. Routes
	router.GET("/health", handler.HealthCheck) // Normal func
	router.GET("/ready", handler.ReadyCheck)   // Normal func
	
	router.POST("/register", authHandler.Register) // Struct method
	router.POST("/login", authHandler.Login)       // Struct method

	// 5. Start Server
	port := cfg.Port
	if port == "" {
		port = "8080" // Default port agar env me nahi hai
	}
	
	log.Printf("Starting secure-identity-service on port %s", port)
	router.Run(":" + port)
}