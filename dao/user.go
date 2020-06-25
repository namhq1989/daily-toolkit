package dao

import (
	"context"

	"github.com/kr/pretty"
	"github.com/namhq1989/daily-toolkit/model"
	"github.com/namhq1989/daily-toolkit/module/database"
)

// UserFindByID ...
func UserFindByID(ctx context.Context, userID model.UUID) model.User {
	var (
		db     = database.GetConnection(ctx)
		result = &model.User{ID: userID}
	)
	defer db.Close()

	err := db.Select(result)
	if err != nil {
		pretty.Println("Error when find expense by user id", err)
	}

	return *result
}
