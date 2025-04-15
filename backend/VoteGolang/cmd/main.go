package main

import (
	"VoteGolang/config"
	"VoteGolang/internal/data"
	"VoteGolang/internal/router"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	// MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Auto-migrate your models
	err = db.AutoMigrate(
		&data.User{},
		&data.Candidate{},
		&data.GeneralNews{},
		&data.Petition{},
		&data.PetitionVote{},
		&data.Vote{},
	)
	if err != nil {
		log.Fatalf("AutoMigration error: %v", err)
	}

	// Start router
	r := router.SetupRouter(db)
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
