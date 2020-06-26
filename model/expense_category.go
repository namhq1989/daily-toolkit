package model

type (
	// ExpenseCategory ...
	ExpenseCategory struct {
		tableName struct{} `pg:"expense_categories"`

		ID    UUID   `sql:"_id,pk" pg:"_id,pk" json:"_id"`
		Name  string `pg:"name" json:"name"`
		Color string `pg:"color" json:"color"`
		Icon  int    `pg:"icon" json:"icon"`
	}
)
