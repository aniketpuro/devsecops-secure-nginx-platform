package main

import (
    "log"

    "github.com/aniketpuro/devsecops-secure-nginx-platform/microservice/internal/config"
    "github.com/aniketpuro/devsecops-secure-nginx-platform/microservice/internal/handler"
    "github.com/aniketpuro/devsecops-secure-nginx-platform/microservice/internal/middleware"
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    cfg := config.Load()

    r := gin.New()

    // Global Middlewares
    r.Use(gin.Recovery())
    r.Use(middleware.RequestID())
    r.Use(middleware.SecurityHeaders())
    r.Use(middleware.RateLimit())
    r.Use(middleware.AuditLog())

    // Public Routes
    r.GET("/health", handler.HealthCheck)
    r.GET("/ready", handler.ReadyCheck)
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // API v1 Routes
    v1 := r.Group("/api/v1")
    {
        auth := v1.Group("/auth")
        {
            auth.POST("/register", handler.Register)
            auth.POST("/login", handler.Login)
            auth.POST("/refresh", handler.RefreshToken)
        }
    }

    log.Printf("🚀 Secure Identity Service started successfully on port %s | Environment: %s", cfg.Port, cfg.Environment)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal("Failed to start server: ", err)
    }
}