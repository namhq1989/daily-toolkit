package model

import (
	"time"
)

type (
	// Expense ...
	Expense struct {
		tableName struct{} `pg:"expenses"`

		ID        UUID            `pg:"_id,pk,type:uuid"`
		User      User            `pg:"user_id"`
		Category  ExpenseCategory `pg:"category_id"`
		Amount    int             `pg:"amount"`
		Note      string          `pg:"note"`
		CreatedAt time.Time       `pg:"created_at"`
	}
)
