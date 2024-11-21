package config

type Database struct {
	Username string
	Password string
	Host     string
	Port     uint16
	DBName   string
	SSLMode  string
}