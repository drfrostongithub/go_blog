package main

import (
	"net/http"

	"hello-go/database"
	"hello-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "se",
		})
	})

	database.DB.AutoMigrate(
		&models.Portfolio{},
		&models.BlogPost{},
		&models.Resume{},
		&models.SocialMedia{},
	)

	r.Run(":8080")
}
