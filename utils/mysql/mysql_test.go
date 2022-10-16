package mysql

import (
	"github.com/wonderivan/logger"
	"testing"
	"tip/utils/config"
)

func Test_MySQL(t *testing.T) {
	m := NewMySQL(config.String("mysql.db"))
	result := []map[string]interface{}{}
	m.DB.Table("sample_info").Find(&result)
	logger.Info("m=%+v", result)
}
