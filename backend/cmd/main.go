package main

import (
	"net/http"

	"github.com/Chofa09/crawfish-backend/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
