package server

import (
	"database/sql"

	"github.com/Chofa09/crawfish-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// URL routes
	urlHandler := handlers.NewURLHandler(db)
	r.GET("/urls", urlHandler.GetAll)
	r.GET("/urls/:id", urlHandler.GetByID)

	return r
}
