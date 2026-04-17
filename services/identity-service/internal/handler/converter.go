package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// User se jo data aayega (jaise YouTube/Video link)
type ConvertRequest struct {
	VideoURL string `json:"video_url" binding:"required"`
}

func ConvertToMP3(c *gin.Context) {
	// 1. Security Check: Middleware se User ID nikalo
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Access"})
		return
	}

	// 2. Request body check karo
	var req ConvertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video_url is required"})
		return
	}

	// 3. RabbitMQ se connect karo
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://admin:securepass123@rabbitmq:5672/"
	}
	
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Message Queue"})
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open a channel"})
		return
	}
	defer ch.Close()

	// 4. Task ID generate karo aur RabbitMQ queue 'mp3_tasks' me daal do
	q, _ := ch.QueueDeclare("mp3_tasks", true, false, false, false, nil)
	taskID := uuid.New().String()
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(taskID), // Yahan hum apna Task ID bhej rahe hain
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push task to queue"})
		return
	}

	// 5. User ko success message bhej do
	c.JSON(http.StatusOK, gin.H{
		"message":   "Conversion task queued successfully!",
		"task_id":   taskID,
		"user_id":   userID,
		"video_url": req.VideoURL,
	})
}