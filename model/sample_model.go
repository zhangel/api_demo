package model

type SampleInfo struct {
	ID         int64  `gorm:"column:id" json:"id"`
	Md5        string `gorm:"column:md5" json:"md5"`
	Sha1       string `gorm:"column:sha1" json:"sha1"`
	Level      int    `gorm:"column:level" json:"level"`
	Operator   string `gorm:"column:operator" json:"operator"`
	CreateTime string `gorm:"column:create_time" json:"create_time"`
}
