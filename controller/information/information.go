package information

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"strconv"
	"tip/common"
	"tip/dao"
	_ "tip/model"
)

type InformationController struct {
	Dao *dao.InformationDao
}

// ShowAccount godoc
// @Summary 获取样本列表数据
// @Tags 标签
// @Produce json
// @Description 获取样本列表数据
// @Param page query int false "默认为 1"   default(1)
// @Param limit query int false "默认为 25" default(25)
// @Param level query int false "默认为 空"
// @Param token query string false "默认为 空"
// @Param md5 query string false "默认为 空"
// @Param sha1 query string false "默认为 空"
// @Success  200 {object}  model.JsonOut{data=[]model.SampleInfo}
// @Failure  500 {object}  model.ServerError
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
// @Produce json
// @Description 新增样本数据
// @Param level formData int true "默认为 空" default(70)
// @Param md5 formData string true "默认为 空" default(2a57220fe8f64481b1311c892b788da5)
// @Param sha1 formData string true "默认为 空" default(ff4cd7d8ee07f35037e834cc0f356f5fa159c871)
// @Param operator formData string false "默认为 空" default(admin)
// @Param token formData string false "默认为 空"
// @Success  200 {object}  model.ExecSuccess
// @Failure  500 {object}  model.ServerError
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
// @Produce json
// @Description 更新样本数据
// @Param level formData int true "默认为 空" default(0)
// @Param id formData int true "默认为 空" default(0)
// @Param token formData int true "默认为 空"
// @Success  200 {object}  model.ExecSuccess
// @Failure  500 {object}  model.ServerError
// @Router /api/v1/information/sample/update [post]
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

// @Summary 删除样本列表
// @Tags 标签
// @Produce json
// @Description 删除样本数据
// @Param level body int true "默认为 空" default(0)
// @Param token formData int true "默认为 空"
// @Success  200 {object}  model.ExecSuccess
// @Failure  500 {object}  model.ServerError
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
