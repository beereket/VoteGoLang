package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Now include the database name vote_database
	dsn := "root:SECRET_PASSWORD@tcp(127.0.0.1:3306)/vote_database?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Error opening database: %s", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Error pinging database: %s", err)
	}

	fmt.Println("✅ Connected to MySQL vote_database successfully!")
}
