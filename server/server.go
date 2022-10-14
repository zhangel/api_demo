package server

import (
	"app_demo/common"
	"app_demo/router"
	"app_demo/utils/config"
)





func ServeRun() {
	httpAddress:=config.String("server.http")
	server:=router.NewRouter()
	common.RegistryDB()
	server.Run(httpAddress)

}

