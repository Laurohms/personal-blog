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

var DB *sql.DB

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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error verifying connection: %v", err)
	}

	log.Println("Conexão bem-sucedida com o banco de dados!")

	DB = db
	createUsersTable()
}

func createUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(255) NOT NULL,
		role ENUM('BASIC', 'ADMIN') NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("erro ao criar a tabela users: %v", err)
	} else {
		log.Println("Tabela users verificada/criada com sucesso!")
	}
}
