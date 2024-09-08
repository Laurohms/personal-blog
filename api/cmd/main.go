package main

import (
	"log"

	"github.com/Laurohms/blog-api/internal/routes"
	"github.com/gin-gonic/gin"
)

const (
	port = ":3000"
)

func main() {
	r := gin.Default()

	err := routes.Start(r)
	if err != nil {
		log.Fatalf("Falha ao carregar as rotas: %v", err)
	}

	err = r.Run(port)
	if err != nil {
		log.Fatalf("error starting application: %v", err)
	}
}
