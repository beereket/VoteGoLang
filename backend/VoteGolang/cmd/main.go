package main

import (
	"VoteGolang/config"
	"VoteGolang/internal/routes"
	"log"
	"net/http"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	defer db.Close()

	router := routes.SetupRoutes(db)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
