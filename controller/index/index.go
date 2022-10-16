package index

import (
	"github.com/gin-gonic/gin"
	"tip/common"
)

type IndexController struct {
}

func (i IndexController) Index(c *gin.Context) {
	data := make([]string, 0)
	common.Json(200, "OK", data, c)
}
