package dao

import (
	"fmt"
	"github.com/wonderivan/logger"
	"tip/common"
	"tip/model"
	"tip/utils/sql_parse"
)

type InformationDao struct {
}

var (
	Table string = "sample_info"
)

func (i *InformationDao) GetDataList(page, limit int, search map[string]interface{}) ([]model.SampleInfo,error) {
	dataList := []model.SampleInfo{}
	offset := (page - 1) * limit
	sqlParse := sql_parse.NewSqlParse()
	sqlDB := sqlParse.Table(Table+" as a")
	for k, v := range search {
		sqlDB = sqlDB.Where(k+"=", fmt.Sprintf("'%v'", v))
	}
	sqlDB.Select([]string{"a.*","b.sha1_number as sha1_number","b.md5_number as md5_number","(md5_number+sha1_number) as md5_sha1_number"})
	sqlDB.Join("left join sample_stat as b on a.md5=b.md5 and a.sha1=b.sha1")
	sqlDB = sqlDB.Limit(limit).Offset(offset)
	sql := sqlDB.Get()
	db:=common.MySQL.DB.Raw(sql).Debug().Scan(&dataList)
	if db.Error != nil {
		return dataList,db.Error
	}
	return dataList,nil
}

func (i *InformationDao) Insert(insertMap map[string]interface{}) error {
	var count int64
	sqlParse := sql_parse.NewSqlParse()
	countSql := sqlParse.Table(Table).Where("md5=", fmt.Sprintf("'%s'", insertMap["md5"].(string))).Count()
	common.MySQL.DB.Raw(countSql).Debug().Scan(&count)
	if count > 0 {
		return fmt.Errorf("md5 %s already exists", insertMap["md5"])
	}
	insertSql := sqlParse.Table(Table).Insert(insertMap)
	common.MySQL.DB.Exec(insertSql)
	return nil
}

func (i *InformationDao) Update(updateMap map[string]interface{}) error {
	var count int64
	sqlParse := sql_parse.NewSqlParse()
	countSql := sqlParse.Table(Table).Where("id=", fmt.Sprintf("'%s'", updateMap["id"])).Count()
	common.MySQL.DB.Raw(countSql).Scan(&count)
	if count == 0 {
		return fmt.Errorf("id %s not exists", updateMap["id"])
	}
	updateSql := sqlParse.Table(Table).Where("id=", updateMap["id"].(string)).Update(updateMap)
	logger.Debug("sql: %s", updateSql)
	common.MySQL.DB.Exec(updateSql)
	return nil
}

func (i *InformationDao) Delete(id string) error {
	var count int64
	sqlParse := sql_parse.NewSqlParse()
	countSql := sqlParse.Table(Table).Where("id=", sqlParse.WrapCharacter(id)).Count()
	common.MySQL.DB.Raw(countSql).Scan(&count)
	if count == 0 {
		return fmt.Errorf("id %s not exists", id)
	}
	delSql := sqlParse.Table(Table).Where("id=", id).Delete()
	logger.Debug("sql: %s", delSql)
	common.MySQL.DB.Exec(delSql)
	return nil
}
