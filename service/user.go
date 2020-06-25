package service

import (
	"context"

	"github.com/namhq1989/daily-toolkit/dao"
	"github.com/namhq1989/daily-toolkit/model"
)

// UserGetByID ...
func UserGetByID(ctx context.Context, userID model.UUID) (result model.User) {
	result = dao.UserFindByID(ctx, userID)
	return
}
