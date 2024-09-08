package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Laurohms/blog-api/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	port = ":3000"
	TEST = "TEST"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("fail loading .env file: %v", err)
	}
	envTest := os.Getenv(TEST)
	fmt.Println(envTest)

	r := gin.Default()

	err = routes.Start(r)
	if err != nil {
		log.Fatalf("Falha ao carregar as rotas: %v", err)
	}

	err = r.Run(port)
	if err != nil {
		log.Fatalf("error starting application: %v", err)
	}
}
