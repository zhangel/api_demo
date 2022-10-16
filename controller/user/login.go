package user

import (
	"api_demo/common"
	"api_demo/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

type Login struct {
}

var (
	RealUserName = "admin"
	RealPassword = "admin"
)

func (l Login) Login(c *gin.Context) {
	username := c.PostForm("username")
	if username != RealUserName {
		common.Json(http.StatusOK, "user name does not exist", false, c)
		return
	}
	password := c.PostForm("password")
	if password != RealPassword {
		common.Json(http.StatusOK, "Password error", false, c)
		return
	}
	//write session information
	session := sessions.Default(c)
	session.Set("username", RealUserName)
	session.Set("is_login", true)
	session.Save()
	loginUser := session.Get("username")
	logger.Debug("username=%v", loginUser)
	common.Json(http.StatusOK, "OK", true, c)
}

func (l Login) Logout(c *gin.Context) {
	session := sessions.Default(c)
	loginUser := session.Get("username")
	logger.Debug("username=%v", loginUser)
	if loginUser != "" {
		session.Delete("username")
		session.Delete("is_login")
		session.Save()
	}
	common.Json(http.StatusOK, "OK", true, c)
}

func (l Login) GenerateToken(c *gin.Context) {
	logger.Debug("GENERATE TOKEN")
	token, err := middleware.GenerateToken("11")
	if err != nil {
		logger.Error("generate token fail,error=%+v", err)
		common.Json(http.StatusOK, "获取token失败", false, c)
		return
	}
	data := map[string]string{"token": token}
	common.Json(http.StatusOK, "OK", data, c)
}
