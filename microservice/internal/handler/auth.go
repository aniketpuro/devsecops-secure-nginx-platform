package handler

import (
    "net/http"
    "time"

    "secure-identity-service/internal/config"
    "secure-identity-service/pkg/jwt"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
    Name     string `json:"name,omitempty"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
    AccessToken string `json:"access_token"`
    ExpiresIn   int    `json:"expires_in"`
    Message     string `json:"message,omitempty"`
}

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "healthy",
        "service": "secure-identity-service",
        "time":    time.Now().UTC(),
    })
}

func ReadyCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := uuid.New().String()

    token, err := jwt.GenerateAccessToken(userID, req.Email, "user", config.Load().JWTSecret, 15*time.Minute)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusCreated, AuthResponse{
        AccessToken: token,
        ExpiresIn:   900,
        Message:     "User registered successfully",
    })
}

func Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := uuid.New().String()

    token, err := jwt.GenerateAccessToken(userID, req.Email, "user", config.Load().JWTSecret, 15*time.Minute)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, AuthResponse{
        AccessToken: token,
        ExpiresIn:   900,
        Message:     "Login successful",
    })
}

func RefreshToken(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Refresh token endpoint - will be implemented later"})
}