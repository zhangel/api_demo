package middleware

import (
	"api_demo/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"strings"
)

type Session struct {
}

func (s Session) Check(c *gin.Context) {

	if strings.Contains(c.Request.RequestURI, "login") || strings.Contains(c.Request.RequestURI, "logout") {
		return
	}

	session := sessions.Default(c)
	username := session.Get("username")
	logger.Debug("username=%v", username)
	if username == nil {
		common.Json(http.StatusOK, "Please login to the system", false, c)
		c.Abort()
	}

}
