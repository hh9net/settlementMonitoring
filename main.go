package main

import (
	"settlementMonitoring/config"
	"settlementMonitoring/db"
	"settlementMonitoring/router"
	"settlementMonitoring/utils"
	"time"
)

//项目介绍注释
// @title 结算数据监控平台
// @version 1.0
// @description Gin swagger 结算数据监控平台
// @host 127.0.0.1:8088
func main() {

	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	db.DBInit() //初始化数据库
	router.RouteInit()
}
