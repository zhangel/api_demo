package index

import (
	"api_demo/common"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (i IndexController) Index(c *gin.Context) {
	data := make([]string, 0)
	common.Json(200, "OK", data, c)
}
