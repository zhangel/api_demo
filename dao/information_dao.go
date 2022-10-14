package dao

import (
	"app_demo/common"
	"app_demo/model"
)

type InformationDao struct {
}

var (
	Table string = "sample_info"
)



func (i *InformationDao) GetDataList() []model.SampleInfo {
	dataList:=[]model.SampleInfo{}
	common.MySQL.DB.Table(Table).Find(&dataList)
	return dataList
}
