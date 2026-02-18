package main

import (
	"net/http"

	"mailbox-api/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	g.Use(cors.Default())

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := g.Group("/api/v1")
	v1.Use(utils.AuthMiddleware())
	{
		v1.POST("/mailboxes", app.createMailbox)
		v1.GET("/mailboxes", app.getMailboxes)
		v1.GET("/mailboxes/:id/messages", app.getMailboxMessages)
		// v1.POST("/messages", app.createMessage)
	}

	return g
}
