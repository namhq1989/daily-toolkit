package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/kr/pretty"
)

// Actors ...
type Actors struct {
	tableName struct{} `pg:"actors"`

	ID        int    `pg:"actor_id,pk"`
	FirstName string `pg:"first_name"`
	LastName  string `pg:"last_name"`
}

// Connect to database
func Connect() {
	db := pg.Connect(&pg.Options{
		User:     "namhq",
		Password: "123456",
		Database: "movie_data",
	})
	defer db.Close()

	// Check db is up and running
	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	actors := make([]Actors, 0)
	err := db.Model(&actors).Where("actor_id > ?", 1).Limit(5).Select()
	pretty.Println("actors", actors)
	fmt.Println("err", err)
}
