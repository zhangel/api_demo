package mysql

import (
	"api_demo/utils/config"
	"github.com/wonderivan/logger"
	"testing"
)

func Test_MySQL(t *testing.T) {
	m := NewMySQL(config.String("mysql.db"))
	result := []map[string]interface{}{}
	m.DB.Table("sample_info").Find(&result)
	logger.Info("m=%+v", result)
}
