package main

import (
	"api-golang-codebase/config"
	"api-golang-codebase/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := app.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close

	server := app.NewServer(cfg, db)
	if err := server
}
