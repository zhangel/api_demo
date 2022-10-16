package common

import (
	"tip/utils/config"
	"tip/utils/mysql"
)

var MySQL *mysql.MySQL

func RegistryDB() {
	MySQL = mysql.NewMySQL(config.String("mysql.db"))
}
