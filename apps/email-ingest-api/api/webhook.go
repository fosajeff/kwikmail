package main

import (
	redisclient "email-ingest-api/internal/redis"
	"email-ingest-api/internal/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type IngestRequest struct {
	To      string `json:"to" binding:"required,email"`
	From    string `json:"from" binding:"required,email"`
	Subject string `json:"subject"`
	Body    string `json:"body" binding:"required"`
}

var streamName = utils.GetEnvString("STREAM_NAME", "email.received")
var rdb = redisclient.New()

func (app *application) webhook(c *gin.Context) {
	var req IngestRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	localPart := strings.Split(req.To, "@")[0]

	id, err := rdb.XAdd(redisclient.Ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"to":         req.To,
			"from":       req.From,
			"subject":    req.Subject,
			"body":       req.Body,
			"local_part": localPart,
			"receivedAt": time.Now().UTC().Format(time.RFC3339),
		},
	}).Result()

	if err != nil {
		log.Printf("Error adding to Redis stream: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process webhook",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "Webhook received", "eventId": id})
}
