package router

import (
	"app_demo/controller/information"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router:=gin.Default()
	router.GET("/api/information/get",information.InformationController{}.GetData)
	return router
}
