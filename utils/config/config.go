package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/wonderivan/logger"
	"os"
	"path"
	"strings"
)

var (
	projectPath = ""
	configPath = ""
	devConfigPath= ""
	section *ini.Section
	projectName="app_demo"
)

func init() {
	_path,err:=os.Getwd()
	if err != nil {
		logger.Fatal("get config fail,error=%+v",err)
	}
	if strings.Contains(_path,projectName) {
		projectPath=fmt.Sprintf("%s/%s",strings.Split(_path,projectName)[0],projectName)
	}
	configPath=path.Join(projectPath,"config")
	devConfigPath=path.Join(configPath,"config_dev.ini")
	Section,err:=ini.Load(devConfigPath)
	if err != nil {
		logger.Fatal("load config fail,error=%+v",err)
	}
	section=Section.Section("")
}

func String(key string) string {
	return section.Key(key).String()
}

