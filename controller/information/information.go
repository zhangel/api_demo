package information

import (
	"app_demo/common"
	"app_demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InformationController struct {
	Dao *dao.InformationDao
}


func (i InformationController) GetData(c *gin.Context) {
		dataList:=i.Dao.GetDataList()
		common.Json(http.StatusOK,"OK",dataList,c)
}

