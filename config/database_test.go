package config

import (
	_ "os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	db := Database{
		Username: "pguser",
		Password: "pgpassword",
		Host:     "pghost",
		Port:     "5432",
		DBName:   "pgdatabase",
	}

	t.Run("testing for invalid username",func(t *testing.T){
		db := db
		db.Username =""
		assert.EqualError(t, db.Validate(), "invalid username")
	})
	t.Run("testing for invalid password",func(t *testing.T){
		db := db
		db.Password =""
		assert.EqualError(t, db.Validate(), "invalid password")
	})
	t.Run("testing for invalid host",func(t *testing.T){
		db := db
		db.Host =""
		assert.EqualError(t, db.Validate(), "invalid host")
	})
	t.Run("testing for invalid port",func(t *testing.T){
		db := db
		db.Port = ""
		assert.EqualError(t, db.Validate(), "invalid port")
	})
	t.Run("testing for invalid database name",func(t *testing.T){
		db := db
		db.DBName =""
		assert.EqualError(t, db.Validate(), "invalid database name")
	})
}