package model

import (
	"time"
)

type (
	// Expense ...
	Expense struct {
		ID         UUID            `gorm:"primary_key,column:_id" json:"_id"`
		UserID     UUID            `gorm:"column:user_id" json:"-"`
		User       User            `json:"user"`
		CategoryID UUID            `gorm:"column:category_id" json:"-"`
		Category   ExpenseCategory `sql:"fk:category_id" json:"category"`
		Amount     int             `gorm:"column:amount" json:"amount"`
		Note       string          `gorm:"column:note" json:"note"`
		CreatedAt  time.Time       `gorm:"column:created_at" json:"createdAt"`
	}
)

// TableName ...
func (Expense) TableName() string {
	return "expenses"
}
