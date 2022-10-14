package common

import (
	"app_demo/utils/config"
	"app_demo/utils/mysql"
)
var MySQL *mysql.MySQL

func RegistryDB() {
	MySQL = mysql.NewMySQL(config.String("mysql.db"))
}
