package handlers

import (
	"net/http"
	"time"

	"github.com/Laurohms/blog-api/internal/config"
	"github.com/Laurohms/blog-api/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	rows, err := config.DB.Query("SELECT id, username, role, created_at FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fail listing users", "details": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var createdAt []uint8

		if err := rows.Scan(&user.ID, &user.Username, &user.Role, &createdAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail listing users in a row", "details": err.Error()})
			return
		}

		createdAtStr := string(createdAt)
		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail parsing created_at", "details": err.Error()})
			return
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error processing user data", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
