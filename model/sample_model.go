package model

type SampleInfo struct {
	ID         int64  `gorm:"column:id" json:"id" example:"1"`
	Md5        string `gorm:"column:md5" json:"md5" example:"2a57220fe8f64481b1311c892b788da5"`
	Sha1       string `gorm:"column:sha1" json:"sha1" example:"ff4cd7d8ee07f35037e834cc0f356f5fa159c871"`
	Level      int    `gorm:"column:level" json:"level" example:"70"`
	Operator   string `gorm:"column:operator" json:"operator" example:"admin"`
	CreateTime string `gorm:"column:create_time" json:"create_time" example:"2022-10-20 11:00:01"`
}

func (s *SampleInfo) TableName() string {
	return "sample_info"
}
