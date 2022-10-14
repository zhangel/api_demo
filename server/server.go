package server

import (
	"api_demo/common"
	"api_demo/router"
	"api_demo/utils/config"
)

func ServeRun() {
	httpAddress := config.String("server.http")
	server := router.NewRouter()
	common.RegistryDB()
	server.Run(httpAddress)

}
