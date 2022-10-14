package common

import (
	"api_demo/utils/config"
	"api_demo/utils/mysql"
)

var MySQL *mysql.MySQL

func RegistryDB() {
	MySQL = mysql.NewMySQL(config.String("mysql.db"))
}
