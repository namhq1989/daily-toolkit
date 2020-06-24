package initialize

import "github.com/namhq1989/daily-toolkit/module/database"

// Init data
func Init() {
	database.Connect()
}
