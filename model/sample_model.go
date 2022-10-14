package model

type SampleInfo struct {
	Md5 		string	`gorm:"column:md5"`
	Sha1 		string	`gorm:"column:sha1"`
	Level   	int		`gorm:"column:level"`
	Operator 	string	`gorm:"column:operator"`
	CreateTime 	string	`gorm:"column:create_time"`
}
