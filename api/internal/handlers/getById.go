package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Laurohms/blog-api/internal/config"
	"github.com/Laurohms/blog-api/internal/models"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id", "details": err.Error()})
		return
	}

	var user models.User
	var createdAt []uint8

	err = config.DB.QueryRow("SELECT id, username, role, created_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Role, &createdAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found", "details": err.Error()})
		return
	}

	createdAtStr := string(createdAt)
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr) // Ajuste o formato conforme necessário
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fail parsing created_at", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
