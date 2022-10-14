package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Json(Code int, Message string, Data interface{}, c *gin.Context) {
	c.JSON(Code, gin.H{"code": Code, "message": Message, "data": Data})
}

func CheckArgs(value, key string, c *gin.Context) bool {
	data := make([]string, 0)
	if value == "" {
		Json(http.StatusOK, fmt.Sprintf("%s cannot be empty", key), data, c)
		return true
	}
	return false
}
