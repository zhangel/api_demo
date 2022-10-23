package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tip/controller/index"
	"tip/controller/information"
	"tip/controller/page"
	"tip/controller/user"
	"tip/docs"
	"tip/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	//router.StaticFS("static", http.Dir("static"))

	router.NoRoute(func(context *gin.Context) {
		page.Page{}.Page404(context)
	})
	router.Use(gin.Recovery())
	store := cookie.NewStore([]byte("cookie_secret"))
	router.Use(sessions.Sessions("gin_session", store))
	//router.Use(middleware.Session{}.Check)
	router.Use(middleware.JWTAuth())
	router.GET("/", index.IndexController{}.Index)
	docs.SwaggerInfo.BasePath = ""
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		//docs.SwaggerInfo.BasePath = "/api/v1/docs/doc.json"
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
