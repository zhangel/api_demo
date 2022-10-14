package router

import (
	"api_demo/controller/index"
	"api_demo/controller/information"
	"api_demo/controller/page"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFS("static", http.Dir("static"))

	router.NoRoute(func(context *gin.Context) {
		page.Page{}.Page404(context)
	})
	router.GET("/", index.IndexController{}.Index)
	v1 := router.Group("/api/v1")
	{
		info := v1.Group("information")
		info.GET("sample/get", information.InformationController{}.Get)
		info.POST("sample/insert", information.InformationController{}.Insert)
		info.POST("sample/update", information.InformationController{}.Update)
		info.POST("sample/delete", information.InformationController{}.Delete)
	}
	return router
}
