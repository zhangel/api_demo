package server

import (
	"tip/common"
	"tip/router"
	"tip/utils/config"
)

func ServeRun() {
	httpAddress := config.String("server.http")
	server := router.NewRouter()
	common.RegistryDB()
	server.Run(httpAddress)

}
