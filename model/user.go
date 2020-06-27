package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	// User ...
	User struct {
		ID        UUID      `gorm:"primary_key,column:_id" json:"_id"`
		Name      string    `gorm:"column:name" json:"name"`
		Phone     string    `gorm:"column:phone" json:"phone"`
		CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	}
)

// TableName ...
func (User) TableName() string {
	return "users"
}

// GenerateToken generate token for authentication
func (user *User) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":  user.ID,
		"name": user.Name,
		"exp":  time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte("app_secret"))
	return tokenString
}
