package model

import "encoding/xml"

//https://gorm.io/zh_CN/docs/models.html

/*
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
*/

type VulInfo struct {
	Id 				int				`gorm:"primaryKey;type:int"`
	Name 			string			`gorm:"column:name;type:varchar(255)"`
	VulnID 			string			`gorm:"column:vuln-id;type:varchar(255)"`
	Published		string			`gorm:"column:published;type:date"`
	Modified		string			`gorm:"column:modified;type:date"`
	Source 			string			`gorm:"column:source;type:varchar(200)"`
	Severity		string			`gorm:"column:severity;type:varchar(255)"`
	VulnType		string			`gorm:"column:vuln-type;type:varchar(255)"`
	VulnDescript 	string			`gorm:"column:vuln-descript;type:text"`
	CveId			string			`gorm:"column:cve-id;type:varchar(255);unique_index"`
	BugtraqId		string			`gorm:"column:bugtraq-id;type:varchar(255)"`
	VulnSolution	string			`gorm:"column:vuln-solution;type:text"`
}

//xml structure

type OtherId struct {
	CveId 	   string 	`xml:"cve-id"`
	BugtraqId  string	`xml:"bugtraq-id"`
}

type Entry struct {
	Name 	string	`xml:"name"`
	VulnId 	string	`xml:"vuln-id"`
	Published string `xml:"published"`
	Modified string	`xml:"modified"`
	Source string	`xml:"source"`
	Severity	string	`xml:"severity"`
	VulnType string `xml:"vuln-type"`
	VulnDescript	string	`xml:"vuln-descript"`
	VulnSolution string `xml:"vuln-solution"`
	OtherId OtherId	`xml:"other-id"`
}

type Cnnvd struct {
	XmlName 	xml.Name	`xml:"cnnvd"'`
	Version 	string		`xml:"cnnvd_xml_version"`
	Entries     []Entry		`xml:"entry"`
}