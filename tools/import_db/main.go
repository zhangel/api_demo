package main

import (
	"github.com/zhangel/logger"
	"tip/tools/import_db/core"
	"github.com/go-ini/ini"
)

func String(key string) string {
	configPath:="./config.ini"
	cfg,err:=ini.Load(configPath)
	if err != nil {
		logger.Fatal("load config file fail,error=%+v",err)
	}
	s:=cfg.Section("")
	return s.Key(key).String()
}

func main() {
	c:=core.NewCore()
	c.Run()
}
