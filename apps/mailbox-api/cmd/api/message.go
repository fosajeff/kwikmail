package main

import (
	"encoding/json"
	"mailbox-api/internal/database"
	"mailbox-api/internal/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) getMailboxMessages(c *gin.Context) {
	mailBoxAddress := c.Param("address")

	// Update last activity timestamp
	if err := app.models.Mailbox.UpdateLastActivity(mailBoxAddress); err != nil {
		// Log error but don't fail the request
		c.Error(err)
	}

	messages, err := app.models.Message.GetByMailboxAddress(mailBoxAddress)

	if messages == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mailbox ID not provided"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get messages": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (app *application) createMailboxMessage(c *gin.Context) {
	var req struct {
		To      string `json:"to" binding:"required,email"`
		From    string `json:"from" binding:"required,email"`
		Subject string `json:"subject"`
		Body    string `json:"body" binding:"required"`
	}

	headerWorkerSecret := c.GetHeader("X-Internal-Key")

	workerSecret := env.GetEnvString("WORKER_SECRET", "default_secret")

	if headerWorkerSecret != workerSecret {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - missing header"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mailBoxId, err := app.models.Message.GetMailBoxIdByAddress(req.To)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Invalid mailbox address": err.Error()})
		return
	}

	rawPayload, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal payload"})
		return
	}

	message := database.Message{
		MailBoxId:  mailBoxId,
		FromEmail:  req.From,
		Subject:    req.Subject,
		Body:       req.Body,
		RawPayload: string(rawPayload),
	}

	if err := app.models.Message.Insert(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to create message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}
