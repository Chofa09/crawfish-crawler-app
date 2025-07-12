package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Chofa09/crawfish-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	DB *sql.DB
}

func NewURLHandler(db *sql.DB) *URLHandler {
	return &URLHandler{DB: db}
}

// GET /urls
func (h *URLHandler) GetAll(c *gin.Context) {
	urls, err := models.GetAllURLs(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch URLs"})
		return
	}
	c.JSON(http.StatusOK, urls)
}

// GET /urls/:id
func (h *URLHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	url, err := models.GetURLById(h.DB, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.JSON(http.StatusOK, url)
}
