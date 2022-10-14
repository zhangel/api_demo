package page

import (
	"app_demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Page struct {
}

func (i Page) Page404(c *gin.Context) {
	data := []string{}
	common.Json(http.StatusOK, "404 not found", data, c)
}
