package main

import (
	"mailbox-api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) getMailboxMessages(c *gin.Context) {
	id := c.Param("id")

	// Update last activity timestamp
	if err := app.models.Mailbox.UpdateLastActivity(id); err != nil {
		// Log error but don't fail the request
		c.Error(err)
	}

	messages, err := app.models.Message.GetByMailboxID(id)

	if messages == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mailbox ID not provided"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get messages": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (app *application) ingestMail(c *gin.Context) {
	var message database.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Message.Insert(&message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to create message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}
