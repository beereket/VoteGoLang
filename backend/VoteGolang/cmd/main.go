package main

import (
	"log"
	"net/http"

	"VoteGolang/internal/database"
	"VoteGolang/internal/routes"
)

func main() {
	database.InitDB()
	r := routes.RegisterRoutes()

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
