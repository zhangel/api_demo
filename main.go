package main

import "tip/server"

// @title 应用系统接口服务
// @version 1.0
// @description 应用系统接口基于Golang Gin框架开发，Gin是一个非常受欢迎的Golang Web框架，在GitHub上已经有47k的星星，后续开发可以参考官方手册,https://gin-gonic.com/
// @termsOfService http://127.0.0.1:8080/api/v1/token
// @contact.name Cameron
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://127.0.0.1:8080
// @BasePath /api/v1
func main() {
	server.ServeRun()
}
