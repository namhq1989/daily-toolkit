package model

type (
	// ExpenseCategory ...
	ExpenseCategory struct {
		tableName struct{} `pg:"expense_categories"`

		ID    UUID   `pg:"_id,pk,type:uuid"`
		Name  string `pg:"name"`
		Color string `pg:"color"`
		Icon  int    `pg:"icon"`
	}
)
