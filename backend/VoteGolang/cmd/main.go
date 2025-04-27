package main

import (
	"VoteGolang/internal/middleware"
	"log"
	"net/http"

	"VoteGolang/internal/database"
	"VoteGolang/internal/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.InitDB()

	r := routes.RegisterRoutes()

	handler := middleware.CORSMiddleware(r)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
