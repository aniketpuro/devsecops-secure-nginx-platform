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

// Config hold karne ke liye struct
type AuthHandler struct {
	cfg *config.Config
}

// Constructor function
func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
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

// Register method
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uuid.New().String()

	// Config se dynamically values fetch ho rahi hain
	token, err := jwt.GenerateAccessToken(userID, req.Email, "user", h.cfg.JWTSecret, h.cfg.JWTAccessExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		AccessToken: token,
		ExpiresIn:   int(h.cfg.JWTAccessExpiry.Seconds()), 
		Message:     "User registered successfully",
	})
}

// Login method
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uuid.New().String()

	// Yahan bhi config use kiya
	token, err := jwt.GenerateAccessToken(userID, req.Email, "user", h.cfg.JWTSecret, h.cfg.JWTAccessExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		AccessToken: token,
		ExpiresIn:   int(h.cfg.JWTAccessExpiry.Seconds()),
		Message:     "Login successful",
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Refresh token endpoint - will be implemented later"})
}