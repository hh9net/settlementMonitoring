package main

import (
	log "github.com/sirupsen/logrus"

	"settlementMonitoring/config"
	"settlementMonitoring/db"
	"settlementMonitoring/router"
	"settlementMonitoring/utils"
	"time"
)

// 项目介绍注释
// @title 结算数据监控平台
// @version 1.0
// @description Gin swagger 结算数据监控平台
// @host 127.0.0.1:8088
func main() {
	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	utils.RedisInit() //初始化redis
	db.DBInit()       //初始化数据库
	//goroutine1
	//go db.HandleDayTasks()
	//goroutine2
	//go db.HandleHourTasks()
	//goroutine3
	//go db.HandleMinutesTasks()
	//http处理
	router.RouteInit()
	for {
		tiker := time.NewTicker(time.Second * 10)
		for {
			log.Println("执行主go程 ", utils.DateTimeFormat(<-tiker.C))
		}
	}
}
