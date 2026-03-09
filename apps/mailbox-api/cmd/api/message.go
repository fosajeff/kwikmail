package main

import (
	"mailbox-api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) getMailboxMessages(c *gin.Context) {
	mailBoxID := c.Param("id")

	// Update last activity timestamp
	if err := app.models.Mailbox.UpdateLastActivity(mailBoxID); err != nil {
		// Log error but don't fail the request
		c.Error(err)
	}

	messages, err := app.models.Message.GetByMailboxID(mailBoxID)

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
	var message database.Message
	mailboxID := c.Param("id")

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.MailBoxId = mailboxID

	err := app.models.Message.Insert(&message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to create message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}
