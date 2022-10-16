package page

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tip/common"
)

type Page struct {
}

func (i Page) Page404(c *gin.Context) {
	data := []string{}
	common.Json(http.StatusOK, "404 not found", data, c)
}
