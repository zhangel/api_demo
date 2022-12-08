package core

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/zhangel/gpool"
	"github.com/zhangel/logger"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"tip/tools/import_db/model"
	"tip/utils/mysql"
)

type Core struct {
	src string
	dest string
	fileType string
	second  int
	goNum 	int
	isDir   bool
	fileList map[string][]string
	lock 	sync.Mutex
	fileSuffix []string
	MySQL *mysql.MySQL
}

var (
	instance *Core
	once sync.Once
	src *string
	dest *string
	fileType *string
	processType map[string]func()
)

func NewCore() *Core {
	once.Do(func() {
		instance = new(Core)
		instance.init()
	})
	return instance
}

func (c *Core) PathExists(path string) (bool) {
	_,err:=os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (c *Core) init() {
	c.Parse()
	processType =make(map[string]func())
	processType["xml"] = c.XmlParse
	processType["json"] = c.JsonParse
	if !c.PathExists(c.dest) {
		err:=os.Mkdir(c.dest,0777)
		if err != nil {
			logger.Fatal("create dest dir fail,error=%+v",err)
		}
	}
}

func (c *Core) Parse() {
	src = flag.String("src","","Please enter the source path, for example: --src=./src_dir")
	dest = flag.String("dest","","Please enter the destination path, for example: --dest=./dest_dir")
	fileType = flag.String("file_type","","Please enter the file type, for example: --file_type=xml,json")
	second:=flag.Int("second",0,"Please enter a few seconds to write a piece of data, for example: --second=1, default is 0 seconds")
	goNum:=flag.Int("go_num",1,"Please enter --go_num=, default to 1")
	flag.Parse()
	if *src == "" {
		fmt.Println("please enter source path,example --src=./src_dir")
		os.Exit(0)
	}
	if *fileType == "" {
		fmt.Println("Please enter the file type, for example: --file_type=xml,json")
		os.Exit(0)
	}
	if *dest == "" {
		fmt.Println("Please enter the destination path, for example: --dest=./dest_dir")
		os.Exit(0)
	}
	c.second = *second
	c.src= *src
	c.dest = *dest
	c.fileType = *fileType
	c.goNum = *goNum
	c.fileSuffix = strings.Split(*fileType,",")
}

func (c *Core) ReplaceXML(respByte []byte) string {
	respStr:=string(respByte)
	respStr = strings.Replace(respStr,"<=","&lt;=",-1)
	respStr = strings.Replace(respStr,"< ","&lt;",-1)
	return respStr
}

func (c *Core) WriteRecord(item model.Entry,f *os.File) {
	fields:=[]string{
		item.Name,
		item.VulnId,
		item.Published,
		item.Modified,
		item.Source,
		item.Severity,
		item.VulnType,
		item.VulnDescript,
		item.OtherId.CveId,
		item.OtherId.BugtraqId,
		item.VulnSolution,
	}
	writeString:=strings.Join(fields,"|")
	f.WriteString(writeString+"\n")
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
	csvFileName:= strings.Replace(path,".xml",".csv",-1)
	csvFileName=filepath.Join(c.dest,csvFileName)
	f,err:= os.OpenFile(csvFileName,os.O_RDWR|os.O_CREATE,0777)
	if err != nil {
		logger.Fatal("open csv file fail,error=%+v",err)
	}
	defer f.Close()
	for _,item:=range Cnnvd.Entries {
		c.Timer()
		g.Add(1)
		go func() {
			defer g.Done()
			c.WriteRecord(item,f)
		}()
	}
	g.Wait()
}
func (c *Core) XmlParse() {
	for _,item:=range c.fileList["xml"] {
		c.ParseXml(item)
	}
}

func (c *Core) JsonProcess(src string) {
	resp,err:=os.ReadFile(src)
	if err != nil {
		logger.Fatal("read json file %s fail,error=%+v",src,err)
	}
	exploit:=make(map[string][]string)
	err=json.Unmarshal(resp,&exploit)
	if err != nil {
		logger.Fatal("unmarshal file %s fail,error=%+v",src,err)
	}
	csvFileName:=strings.Replace(src,".json",".csv",-1)
	csvFileName=filepath.Join(c.dest,csvFileName)
	f,err:=os.OpenFile(csvFileName,os.O_RDWR|os.O_CREATE,0777)
	if err != nil {
		logger.Fatal("create csv file %s fail,error=%+v",csvFileName,err)
	}
	defer f.Close()
	for key,list:=range exploit {
		for _,item:=range list {
			writeStr:=fmt.Sprintf("%s\t%s\n",key,item)
			f.WriteString(writeStr)
		}
	}
}

func (c *Core) JsonParse() {
	for _,item:=range c.fileList["json"] {
		c.JsonProcess(item)
	}
}

func (c *Core) ScanDir() {
	entry,err:=os.ReadDir(c.src)
	if err != nil {
		logger.Error("scan dir %s fail,error=%+v",c.src,err)
		return
	}
	if len(entry) ==0 {
		logger.Error("%s entry is empty",c.src)
		return
	}
	xmlEntry:=make(map[string][]string)
	for _,file:=range entry {
		suffix:=filepath.Ext(file.Name())
		for _,fType := range c.fileSuffix {
			fileSuffix:=fmt.Sprintf(".%s",fType)
			if suffix == fileSuffix {
				xmlFile:=filepath.Join(c.src,file.Name())
				xmlEntry[fType] = append(xmlEntry[fType],xmlFile)
			}
		}
	}
	c.fileList = xmlEntry
}

func (c *Core) importData() {
	for _type,_:=range c.fileList {
		if _, ok:= processType[_type]; ok {
			processType[_type]()
		}
	}
}

func (c *Core) ParseDir() {
	c.ScanDir()
	c.importData()
}

func (c *Core) ReadXml() {
	fi,err:=os.Stat(c.src)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s file does not exist\n",c.src)
			return
		}
	}
	if fi.IsDir() {
		c.isDir = true
		c.ParseDir()
	} else {
		c.ParseXml(c.src)
	}
}

func (c *Core) Run() {
	c.ReadXml()
}