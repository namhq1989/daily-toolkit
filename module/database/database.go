package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/kr/pretty"
)

var (
	db *pg.DB
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	data, _ := q.FormattedQuery()
	pretty.Println(string(data))
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}

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

	db.AddQueryHook(dbLogger{})

}

// GetConnection return a connection in pool
func GetConnection(ctx context.Context) *pg.Conn {
	return db.WithContext(ctx).Conn()
}
