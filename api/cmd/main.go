package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port = ":3000"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	err := r.Run(port)
	if err != nil {
		log.Fatalf("error starting application: %v", err)
	}
}
