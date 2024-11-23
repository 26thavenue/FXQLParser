package config

import (
	"fmt"
	"os"
)

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func (d *Database) URL() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.SSLMode,
	)
}

func NewDB() (*Database, error) {
	username, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_USER env variable set")
	}
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_PASSWORD env variable set")
	}
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_HOST env variable set")
	}
	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_PORT env variable set")
	}
	dbName, ok := os.LookupEnv("POSTGRES_DBNAME")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_PORT env variable set")
	}
	sslMode, ok := os.LookupEnv("POSTGRES_SSLMODE")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_SSLMODE env variable set")
	}
	db := &Database{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName,
		SSLMode:  sslMode,
	}

	err := db.Validate()

	if err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return db,nil

}

func (d *Database) Validate() error {
	if d.DBName == "" {
		return fmt.Errorf("invalid database name")
	}

	if d.Host == "" {
		return fmt.Errorf("invalid host")
	}

	if d.Username == "" {
		return fmt.Errorf("invalid username")
	}

	if d.Password == "" {
		return fmt.Errorf("invalid password")
	}

	if d.Port == "" {
		return fmt.Errorf("invalid port")
	}

	if d.SSLMode == "" {
		return fmt.Errorf("invalid sslmode")
	}

	return nil
}
