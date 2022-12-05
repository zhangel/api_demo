package mysql

import (
	"fmt"
	"github.com/zhangel/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
	"net/url"
	"sync"
)

type MySQL struct {
	url string
	DB  *gorm.DB
}

var (
	instance *MySQL
	once     sync.Once
)

func NewMySQL(dbURL string) *MySQL {
	once.Do(func() {
		instance = new(MySQL)
		instance.Init(dbURL)
	})
	return instance
}

func (m *MySQL) Parse() string {
	cfg, err := url.Parse(m.url)
	if err != nil {
		logger.Fatal("parse mysql url fail,error=%+v", err)
	}
	username := cfg.User.Username()
	password, isPassword := cfg.User.Password()
	if isPassword == false {
		logger.Fatal("get mysql url password fail")
	}
	dbLink := fmt.Sprintf("%s:%s@tcp(%s)%s?autocommit=true&charset=utf8", username, password, cfg.Host, cfg.Path)
	return dbLink
}

func (m *MySQL) Init(dbURL string) {
	m.url = dbURL
	dsnURL := m.Parse()
	dbConfig := &gorm.Config{PrepareStmt: false,Logger:dbLogger.Default.LogMode(dbLogger.Silent)}

	db, err := gorm.Open(mysql.Open(dsnURL), dbConfig)
	if err != nil {
		logger.Fatal("connection mysql fail,error=%+v,db_url=%s", err, dsnURL)
	}
	m.DB = db
}
