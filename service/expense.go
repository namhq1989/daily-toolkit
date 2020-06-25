package service

import (
	"context"

	"github.com/namhq1989/daily-toolkit/dao"
	"github.com/namhq1989/daily-toolkit/model"
)

// ExpenseGetByUserID ...
func ExpenseGetByUserID(ctx context.Context, userID model.UUID, query model.CommonQuery) (result []model.Expense, total int) {
	result, total = dao.ExpenseFindByUserID(ctx, userID, query)
	return
}
