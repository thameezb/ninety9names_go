package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //pg driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	m, err := migrate.New(
		"file://./migrations",
		dbURL)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Steps(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Migrations complete")
}
