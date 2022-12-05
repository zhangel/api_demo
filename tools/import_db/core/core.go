package core

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/zhangel/logger"
	"github.com/zhangel/gpool"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"tip/tools/import_db/model"
	"tip/utils/mysql"
	"time"
)

type Core struct {
	DB *gorm.DB
	path string
	second  int
	goNum 	int
	isDir   bool
	fileList []string
}

var (
	instance *Core
	once sync.Once
	src *string
	tableName = "vul_infos"
	fileSuffix = ".xml"
)

func NewCore(URL string) *Core {
	once.Do(func() {
		instance = new(Core)
		instance.init(URL)
	})
	return instance
}

func (c *Core) createTable() {
	if !c.DB.Migrator().HasTable(&model.VulInfo{}) {
		c.DB.Migrator().CreateTable(&model.VulInfo{})
	} else {
		//logger.Warn("vul_infos already exists")
	}
}

func (c *Core) CheckIsDir() {

}

func (c *Core) init(URL string) {
	conn :=mysql.NewMySQL(URL)
	c.DB = conn.DB
	c.createTable()
	c.Parse()

}

func (c *Core) Parse() {
	src = flag.String("path","","Please enter the file path, for example: -- path=./xml_ file")
	second:=flag.Int("second",0,"Please enter a few seconds to write a piece of data, for example: -- second=1, default is 0 seconds")
	goNum:=flag.Int("go_num",1,"Please enter --go_num=, default to 1")
	flag.Parse()
	if *src == "" {
		fmt.Println("please enter file path,example --path=./xml_file")
		os.Exit(0)
	}
	c.second = *second
	c.path = *src
	c.goNum = *goNum
}

func (c *Core) ReplaceXML(respByte []byte) string {
	respStr:=string(respByte)
	respStr = strings.Replace(respStr,"<=","&lt;=",-1)
	respStr = strings.Replace(respStr,"< ","&lt;",-1)
	return respStr
}

func (c *Core) WriteRecord(item model.Entry) {
	insert:=model.VulInfo{
		Name:item.Name,
		VulnID:item.VulnId,
		Published:item.Published,
		Modified:item.Modified,
		Source:item.Source,
		Severity:item.Severity,
		VulnType:item.VulnType,
		VulnDescript:item.VulnDescript,
		CveId:item.OtherId.CveId,
		BugtraqId: item.OtherId.BugtraqId,
		VulnSolution:item.VulnSolution,
	}
	c.DB.Create(&insert)
}

func (c *Core) Timer() {
	if c.second != 0 {
		time.Sleep(time.Second*time.Duration(c.second))
	}
}

func (c *Core) ParseXml(path string) {
	respByte,err:=os.ReadFile(path)
	if err != nil {
		logger.Error("read file %s fail,error=%+v",path,err)
		return
	}
	respStr:=c.ReplaceXML(respByte)
	Cnnvd:=model.Cnnvd{}
	decoder:=xml.NewDecoder(bytes.NewReader([]byte(respStr)))
	decoder.Strict=false
	err=decoder.Decode(&Cnnvd)
	if err != nil {
		logger.Error("parse file %s fail,error=%+v",path,err)
		return
	}
	g:=gpool.New(c.goNum)
	for _,item:=range Cnnvd.Entries {
		c.Timer()
		g.Add(1)
		go func() {
			defer g.Done()
			c.WriteRecord(item)
		}()
	}
	g.Wait()
}

func (c *Core) ScanDir() {
	entry,err:=os.ReadDir(c.path)
	if err != nil {
		logger.Error("scan dir %s fail,error=%+v",c.path,err)
		return
	}
	if len(entry) ==0 {
		logger.Error("%s entry is empty",c.path)
		return
	}
	xmlEntry:=make([]string,0)
	for _,file:=range entry {
		suffix:=filepath.Ext(file.Name())
		if !file.IsDir() && suffix == fileSuffix {
			xmlFile:=filepath.Join(c.path,file.Name())
			xmlEntry=append(xmlEntry,xmlFile)
		}
	}
	c.fileList = xmlEntry
}

func (c *Core) importData() {
	for _,file:=range c.fileList {
		c.ParseXml(file)
	}
}

func (c *Core) ParseDir() {
	c.ScanDir()
	c.importData()
}

func (c *Core) ReadXml() {
	fi,err:=os.Stat(c.path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s file does not exist\n",c.path)
			return
		}
	}
	if fi.IsDir() {
		c.isDir = true
		c.ParseDir()
	} else {
		c.ParseXml(c.path)
	}
}

func (c *Core) Run() {
	c.ReadXml()
}