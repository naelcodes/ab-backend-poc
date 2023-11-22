// database/postgres.go
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

func InitDB() *xorm.Engine {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	var err error
	engine, err = xorm.NewEngine("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error creating XORM engine: %v", err)
	}

	if err := engine.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Connected to PostgreSQL database")
	return engine
}
