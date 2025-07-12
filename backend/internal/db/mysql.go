package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// Load env from root-level .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Failed to ping DB: %v", err)
	}

	log.Println("✅ Connected to MySQL")
	DB = db
}
