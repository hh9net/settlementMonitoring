package main

import (
	"settlementMonitoring/config"
	"settlementMonitoring/db"
	"settlementMonitoring/router"
	"settlementMonitoring/utils"

	"time"
)

func main() {

	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	db.DBInit()
	router.RouteInit()
}
