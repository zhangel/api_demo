package information

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"strconv"
	"tip/common"
	"tip/dao"
)

type InformationController struct {
	Dao *dao.InformationDao
}

// @Summary 获取样本列表数据
// @Tags 标签
// @Router /api/v1/information/sample/get [get]
func (i InformationController) Get(c *gin.Context) {
	searchMap := map[string]interface{}{}
	limitNum := 25
	var err error
	limit := c.Query("limit")
	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil {
			logger.Error("convert limit to int fail,error=%+v", err)
		}
	}
	pageNum := 1
	page := c.Query("page")
	if page != "" {
		pageNum, err = strconv.Atoi(page)
		if err != nil {
			logger.Error("convert page to int fail,error=%+v", err)
		}
	}
	md5 := c.Query("md5")
	if md5 != "" {
		searchMap["md5"] = md5
	}
	sha1 := c.Query("sha1")
	if sha1 != "" {
		searchMap["sha1"] = sha1
	}
	level := c.Query("level")
	if level != "" {
		levelInt, err := strconv.Atoi(level)
		if err != nil {
			logger.Error("convert page to int fail,error=%+v", err)
		}
		searchMap["level"] = levelInt
	}
	dataList := i.Dao.GetDataList(pageNum, limitNum, searchMap)
	common.Json(http.StatusOK, "OK", dataList, c)
}

// @Summary 新增样本信息
// @Tags 标签
// @Router /api/v1/information/sample/insert [post]
func (i InformationController) Insert(c *gin.Context) {
	md5 := c.PostForm("md5")
	if common.CheckArgs(md5, "md5", c) {
		return
	}
	sha1 := c.PostForm("sha1")
	if common.CheckArgs(sha1, "sha1", c) {
		return
	}
	level := c.PostForm("level")
	if common.CheckArgs(level, "level", c) {
		return
	}
	operator := c.PostForm("operator")
	if common.CheckArgs(operator, "operator", c) {
		return
	}
	insertMap := map[string]interface{}{"md5": md5, "sha1": sha1, "level": level, "operator": operator}
	err := i.Dao.Insert(insertMap)
	if err != nil {
		common.Json(http.StatusInternalServerError, err.Error(), false, c)
		return
	}
	common.Json(http.StatusOK, "OK", true, c)
}

// @Summary 更新样本信息
// @Tags 标签
// @Router /api/v1/information/sample/update [get]
func (i InformationController) Update(c *gin.Context) {
	updateMap := map[string]interface{}{}
	id := c.PostForm("id")
	if common.CheckArgs(id, "id", c) {
		return
	}
	level := c.PostForm("level")
	if common.CheckArgs(level, "level", c) {
		return
	}
	updateMap["level"] = level
	updateMap["id"] = id
	err := i.Dao.Update(updateMap)
	if err != nil {
		common.Json(http.StatusInternalServerError, err.Error(), false, c)
		return
	}
	common.Json(http.StatusOK, "OK", true, c)
}

// @Summary 删除样本列表数据
// @Tags 标签
// @Router /api/v1/information/sample/delete [post]
func (i InformationController) Delete(c *gin.Context) {
	id := c.PostForm("id")
	if common.CheckArgs(id, "id", c) {
		return
	}
	err := i.Dao.Delete(id)
	if err != nil {
		common.Json(http.StatusInternalServerError, err.Error(), false, c)
		return
	}
	common.Json(http.StatusOK, "OK", true, c)
}
