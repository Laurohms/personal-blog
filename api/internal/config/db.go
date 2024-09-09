package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_DATABASE = "DB_DATABASE"
	DB_PORT     = "DB_PORT"
)

var db *sql.DB

func InitDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUser := os.Getenv(DB_USERNAME)
	dbPassword := os.Getenv(DB_PASSWORD)
	dbHost := os.Getenv(DB_HOST)
	dbPort := os.Getenv(DB_PORT)
	dbDatabase := os.Getenv(DB_DATABASE)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error verifying connection: %v", err)
	}

	log.Println("Conexão bem-sucedida com o banco de dados!")

}
