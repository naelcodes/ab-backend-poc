// database/postgres.go
package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Database struct {
	*xorm.Engine
}

var engine *Database
var once sync.Once

func GetDatabase() *Database {
	once.Do(func() {
		dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"))

		xormEngine, err := xorm.NewEngine("postgres", dbURL)
		if err != nil {
			log.Fatalf("Error creating XORM engine: %v", err)
		}

		if err := engine.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}

		log.Println("Connected to PostgreSQL database")
		engine = &Database{xormEngine}
	})
	return engine
}
