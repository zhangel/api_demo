package router

import (
	"api_demo/controller/index"
	"api_demo/controller/information"
	"api_demo/controller/page"
	"api_demo/controller/user"
	"api_demo/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFS("static", http.Dir("static"))

	router.NoRoute(func(context *gin.Context) {
		page.Page{}.Page404(context)
	})
	router.Use(gin.Recovery())
	store := cookie.NewStore([]byte("cookie_secret"))
	router.Use(sessions.Sessions("gin_session", store))
	//router.Use(middleware.Session{}.Check)
	router.Use(middleware.JWTAuth())
	router.GET("/", index.IndexController{}.Index)

	v1 := router.Group("/api/v1")
	{
		v1.POST("login", user.Login{}.Login)
		v1.POST("logout", user.Login{}.Logout)
		v1.GET("token", user.Login{}.GenerateToken)
		info := v1.Group("information")
		info.GET("sample/get", information.InformationController{}.Get)
		info.POST("sample/insert", information.InformationController{}.Insert)
		info.POST("sample/update", information.InformationController{}.Update)
		info.POST("sample/delete", information.InformationController{}.Delete)
	}
	return router
}
