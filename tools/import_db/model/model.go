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
	Name 			string			`gorm:"column:name;type:varchar(255)" json:"name"`
	VulnID 			string			`gorm:"column:vuln-id;type:varchar(255)" json:"vuln-id"`
	Published		string			`gorm:"column:published;type:date" json:"published"`
	Modified		string			`gorm:"column:modified;type:date" json:"modified"`
	Source 			string			`gorm:"column:source;type:varchar(200)" json:"source"`
	Severity		string			`gorm:"column:severity;type:varchar(255)" json:"severity"`
	VulnType		string			`gorm:"column:vuln-type;type:varchar(255) json:"vuln-type"`
	VulnDescript 	string			`gorm:"column:vuln-descript;type:text" json:"vuln-descript"`
	CveId			string			`gorm:"column:cve-id;not null;type:varchar(255);uniqueIndex" json:"cve-id"`
	BugtraqId		string			`gorm:"column:bugtraq-id;type:varchar(255)" json:"bugtraq-id"`
	VulnSolution	string			`gorm:"column:vuln-solution;type:text" json:"vuln-solution"`
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
