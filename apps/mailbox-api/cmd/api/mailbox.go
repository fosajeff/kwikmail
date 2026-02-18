package main

import (
	"mailbox-api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createMailbox(c *gin.Context) {
	var mailbox database.Mailbox
	userId := c.MustGet("user_id").(string)

	mailbox.UserId = userId

	if err := c.ShouldBindJSON(&mailbox); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Mailbox.Insert(&mailbox)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to create mailbox": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mailbox)
}

func (app *application) getMailboxes(c *gin.Context) {
	var mailboxes []*database.Mailbox
	userId := c.MustGet("user_id").(string)

	mailboxes, err := app.models.Mailbox.GetAll(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to get mailboxes": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mailboxes)
}
