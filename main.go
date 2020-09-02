package main

import (
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/config"
	"settlementMonitoring/db"
	"settlementMonitoring/router"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"time"
)

// 项目介绍注释
// @title 结算数据监控平台
// @version 1.0
// @description Gin swagger 结算数据监控平台
// @host 127.0.0.1:8088
func main() {
	conf := config.ConfigInit() //初始化配置
	log.Println("配置文件信息：", *conf)
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)
	types.RedisAddr = conf.RedisAddr
	log.Println("RedisAddrConf:=", conf.RedisAddr)
	utils.RedisInit() //初始化redis

	//黑名单数据库：	"root:Microvideo_1@tcp(122.51.24.189:3307)/blacklist?charset=utf8&parseTime=true&loc=Local"
	types.HmdDBAddrconf = conf.HMUserName + ":" + conf.HMPass + "@tcp(" + conf.HMHostname + ":" + conf.HMPort + ")/" + conf.HMdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	log.Println("HmdDBAddrconf:=", types.HmdDBAddrconf)

	//结算监控数据库 "root:Microvideo_1@tcp(122.51.24.189:3307)/blacklist?charset=utf8&parseTime=true&loc=Local"
	mstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + conf.Mdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	IpAddress := conf.IpAddress
	db.DBInit(mstr) //初始化数据库
	//快照频率
	types.Frequency = conf.Frequency
	log.Println("10分钟快照的频率 的条数：", types.Frequency)
	//kafkaip
	types.KafkaIp = conf.KafkaIp
	log.Println("KafkaIp:", types.KafkaIp)

	//goroutine1
	go db.HandleDayTasks()
	//goroutine2
	go db.HandleHourTasks()
	//goroutine3
	go db.HandleMinutesTasks()
	//http处理
	router.RouteInit(IpAddress)
	for {
		tiker := time.NewTicker(time.Second * 10)
		for {
			log.Println("执行主go程 处理kafka数据+++++++++++++++++++++++++++++++++++++++++++++++++++++++++处理kafka数据", utils.DateTimeFormat(<-tiker.C))
			//处理kafka数据
			utils.ConsumerGroup()
		}
	}
}
