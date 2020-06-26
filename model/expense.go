package model

import (
	"time"
)

type (
	// Expense ...
	Expense struct {
		tableName struct{} `pg:"expenses"`

		ID         UUID             `sql:"_id,pk" pg:"_id,pk" json:"_id"`
		UserID     UUID             `pg:"user_id" json:"-"`
		User       *User            `pg:"fk:user_id" json:"user"`
		CategoryID UUID             `pg:"category_id" json:"-"`
		Category   *ExpenseCategory `sql:"fk:category_id" json:"category"`
		Amount     int              `pg:"amount" json:"amount"`
		Note       string           `pg:"note" json:"note"`
		CreatedAt  time.Time        `pg:"created_at" json:"createdAt"`
	}
)
