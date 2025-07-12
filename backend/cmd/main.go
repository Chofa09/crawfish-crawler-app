package main

import (
	"log"

	"github.com/Chofa09/crawfish-backend/internal/db"
	"github.com/Chofa09/crawfish-backend/server"
)

func main() {

	database := db.InitDB()

	defer database.Close()

	router := server.SetupRouter(database)

	log.Printf("🚀 Server running on port 8080")
	router.Run(":8080")
}
