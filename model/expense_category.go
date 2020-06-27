package model

type (
	// ExpenseCategory ...
	ExpenseCategory struct {
		ID    UUID   `gorm:"primary_key,column:_id" json:"_id"`
		Name  string `gorm:"column:name" json:"name"`
		Color string `gorm:"column:color" json:"color"`
		Icon  int    `gorm:"column:icon" json:"icon"`
	}
)

// TableName ...
func (ExpenseCategory) TableName() string {
	return "expense_categories"
}
