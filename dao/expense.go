package dao

import (
	"context"

	"github.com/kr/pretty"
	"github.com/namhq1989/daily-toolkit/model"
	"github.com/namhq1989/daily-toolkit/module/database"
)

// ExpenseFindByUserID ...
func ExpenseFindByUserID(ctx context.Context, userID model.UUID, query model.CommonQuery) (result []model.Expense, total int) {
	var (
		db = database.GetConnection(ctx)
	)
	defer db.Close()

	err := db.Model(&result).Where("user_id = ?", userID).Join("JOIN users as u").JoinOn("u._id = ?", userID).Limit(query.Limit).Select()
	if err != nil {
		pretty.Println("Error when find expense by user id", err)
	}

	return result, 0
}
