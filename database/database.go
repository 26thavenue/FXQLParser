package database

import (
	_ "fmt"
	"log"

	"github.com/26thavenue/FXQLParser/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Instance *gorm.DB
}

type Transaction struct {
    gorm.Model
    SourceCurrency     string
    DestinationCurrency string
    SellPrice          int
    BuyPrice           int
    CapAmount          int
}

var DBInstance *DB

func Connect() {
	dbConfig, err := config.NewDB()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	db, err := gorm.Open(postgres.Open(dbConfig.URL()),&gorm.Config{} )
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&Transaction{})
	if err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}

	log.Println("Successfully connected to the database")

	DBInstance = &DB{Instance: db}
}
