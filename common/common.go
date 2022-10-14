package common

import "github.com/gin-gonic/gin"

func Json(Code int,Message string,Data interface{},c *gin.Context) {
	c.JSON(Code,gin.H{"code":Code,"message":Message,"data":Data})
}