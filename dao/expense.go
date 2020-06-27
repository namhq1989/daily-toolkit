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
		db   = database.GetConnection(ctx)
		user model.User
	)
	defer db.Close()

	if err := db.Model(&result).
		Where("user_id = ?", userID).
		Limit(query.Limit).
		Association("users").
		Find(&result).Error; err != nil {
		pretty.Println("err", err)
	}

	pretty.Println(result)

	return result, 0
}
