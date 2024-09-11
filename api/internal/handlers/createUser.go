package handlers

import (
	"net/http"
	"time"

	"github.com/Laurohms/blog-api/internal/config"
	"github.com/Laurohms/blog-api/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "bcrypt error"})
	}

	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()

	stmt, err := config.DB.Prepare("INSERT INTO users (username, password, role, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao preparar a inserção no banco de dados"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.Role, user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir usuário no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso!"})
}
