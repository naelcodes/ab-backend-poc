// database/postgres.go
package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"gocloud.dev/postgres"
	_ "gocloud.dev/postgres/awspostgres"
	//"gorm.io/driver/postgres"
)

type Database struct {
	*xorm.Engine
}

var engine *Database
var once sync.Once

func GetDatabase() *Database {
	once.Do(func() {
		dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"))

		db, err := postgres.Open(context.Background(), dbURL)
		if err != nil {
			return
		}
		defer db.Close()

		xormEngine, err := xorm.NewEngine("postgres", dbURL)
		if err != nil {
			log.Fatalf("Error creating XORM engine: %v", err)
		}
		xormEngine.DB().DB = db

		if err := xormEngine.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}

		log.Println("Connected to PostgreSQL database")
		engine = &Database{xormEngine}
	})
	return engine
}
