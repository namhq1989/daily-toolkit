package database

import (
	"context"

	"github.com/jinzhu/gorm"

	// For postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

// Connect to database
func Connect() {
	// FIXME: change config
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=namhq dbname=daily_toolkit password=123456 sslmode=disable")
	// defer db.Close()

	// Check db is up and running
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
}

// GetConnection return a connection in pool
func GetConnection(ctx context.Context) *gorm.DB {
	return db.New()
}
