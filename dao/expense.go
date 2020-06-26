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

	err := db.Model(&result).
		Where("user_id = ?", userID).
		Relation("User").
		Relation("Category").
		// ColumnExpr("expense.*").
		// ColumnExpr("u.name AS user__name, u._id AS user___id, u.phone AS user__phone").
		// Join("JOIN users AS u ON u._id = expense.user_id").
		// ColumnExpr("c.name AS category__name, c._id AS category___id, c.color AS category__color, c.icon AS category__icon").
		// Join("JOIN expense_categories AS c ON c._id = expense.category_id").
		Limit(query.Limit).
		Select()
	if err != nil {
		pretty.Println("Error when find expense by user id", err)
	}

	return result, 0
}
