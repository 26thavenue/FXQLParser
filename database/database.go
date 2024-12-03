package database

import (
	"database/sql"
	_ "fmt"
	"log"

	"github.com/26thavenue/FXQLParser/config"
	_ "github.com/lib/pq"
)

type DB struct {
	Instance *sql.DB
}

var DBInstance *DB

func Connect() {
	dbConfig, err := config.NewDB()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	db, err := sql.Open("postgres", dbConfig.URL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}

	log.Println("Successfully connected to the database")

	DBInstance = &DB{Instance: db}
}
