package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	// User ...
	User struct {
		tableName struct{} `pg:"users"`

		ID        UUID      `sql:"_id,pk,type:uuid" pg:"_id,pk" json:"_id"`
		Name      string    `pg:"name" json:"name"`
		Phone     string    `pg:"phone" json:"phone"`
		CreatedAt time.Time `pg:"created_at" json:"-"`
	}
)

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
