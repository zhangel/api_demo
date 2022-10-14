package dao

import (
	"app_demo/common"
	"app_demo/model"
	"fmt"
)

type InformationDao struct {
}

var (
	Table string = "sample_info"
)

func (i *InformationDao) GetDataList(page, limit int, search map[string]string) []model.SampleInfo {
	dataList := []model.SampleInfo{}
	offset := (page - 1) * limit
	db := common.MySQL.DB.Table(Table).Offset(offset).Limit(limit)
	for k, v := range search {
		db.Where(fmt.Sprintf("%s=?", k), v)
	}
	db.Find(&dataList)
	return dataList
}

func (i *InformationDao) Insert(insertMap map[string]interface{}) error {
	var count int64
	common.MySQL.DB.Table(Table).Where("md5=?", insertMap["md5"]).Count(&count)
	if count > 0 {
		return fmt.Errorf("md5 %s already exists", insertMap["md5"])
	}
	common.MySQL.DB.Table(Table).Create(insertMap)
	return nil
}

func (i *InformationDao) Update(updateMap map[string]interface{}) error {
	var count int64
	common.MySQL.DB.Table(Table).Where("id=?", updateMap["id"]).Count(&count)
	if count == 0 {
		return fmt.Errorf("id %s not exists", updateMap["id"])
	}
	common.MySQL.DB.Table(Table).Where(fmt.Sprintf("%s=?", "id"), updateMap["id"]).Updates(updateMap)
	return nil
}

func (i *InformationDao) Delete(id string) error {
	var count int64
	common.MySQL.DB.Table(Table).Where("id=?", id).Count(&count)
	if count == 0 {
		return fmt.Errorf("id %s not exists", id)
	}
	var sample model.SampleInfo
	common.MySQL.DB.Table(Table).Where("id=?", id).Delete(&sample)
	return nil
}
