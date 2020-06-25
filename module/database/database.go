package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

var (
	db *pg.DB
)

// Connect to database
func Connect() {
	// FIXME: change config
	db = pg.Connect(&pg.Options{
		User:     "namhq",
		Password: "123456",
		Database: "daily_toolkit",
	})
	// defer db.Close()

	// Check db is up and running
	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}
}

// GetConnection return a connection in pool
func GetConnection(ctx context.Context) *pg.Conn {
	return db.WithContext(ctx).Conn()
}
