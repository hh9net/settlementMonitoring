package main

import (
	"github.com/garyburd/redigo/redis"
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
	log.Println("结算监控配置文件信息：", *conf)
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)

	//黑名单数据库：	"root:Microvideo_1@tcp(122.51.24.189:3307)/blacklist?charset=utf8&parseTime=true&loc=Local"
	types.HmdDBAddrconf = conf.HMUserName + ":" + conf.HMPass + "@tcp(" + conf.HMHostname + ":" + conf.HMPort + ")/" + conf.HMdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	log.Println("HmdDBAddrconf:=", types.HmdDBAddrconf)
	db.HmdDBInit() //初始化黑名单数据库

	//结算监控数据库 "root:Microvideo_1@tcp(122.51.24.189:3307)/blacklist?charset=utf8&parseTime=true&loc=Local"
	mstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + conf.Mdatabasename + "?charset=utf8&parseTime=true&loc=Local"

	db.DBInit(mstr) //初始化数据库
	HSDZstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + "localsettle" + "?charset=utf8&parseTime=true&loc=Local"
	db.HSDZGormClientDB = utils.HSDZInitGormDB(HSDZstr)

	//快照频率
	types.Frequency = conf.Frequency
	log.Println("10分钟快照的频率 的条数：", types.Frequency)
	//kafkaip
	types.KafkaIpa = conf.KafkaIpa
	types.KafkaIpb = conf.KafkaIpb
	types.KafkaIpc = conf.KafkaIpc
	types.DdkafkaTopic = conf.DdkafkaTopic
	types.ZdzkafkaTopic = conf.ZdzkafkaTopic
	log.Println("KafkaIp:", types.KafkaIpa, types.KafkaIpb, types.KafkaIpc, "DdkafkaTopic:", types.DdkafkaTopic, "zdzkafkaTopic:", types.ZdzkafkaTopic)

	types.RedisAddr = conf.RedisAddr
	log.Println("RedisAddrConf:=", conf.RedisAddr)

	types.HlsyncAddr = conf.HlsyncAddr
	log.Println("HlsyncAddrConf:=", conf.HlsyncAddr)

	types.Parkids = conf.Parkids
	types.Tradestarttime = conf.Tradestarttime

	utils.Pool = &redis.Pool{
		MaxIdle:     16,  //最大空闲连接数
		MaxActive:   0,   //最大活跃连接数  0为没有限制
		IdleTimeout: 300, //空闲连接超时时间
		//连接方法
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.RedisAddr)
		},
	}

	defer func() {
		_ = utils.Pool.Close()
	}()
	IpAddress := conf.IpAddress

	//goroutine1
	go db.HandleDayTasks()
	//goroutine2
	go db.HandleHourTasks()
	//goroutine3
	go db.HandleMinutesTasks()
	//goroutine4
	go db.HandleSixHourTasks()
	////kafka处理
	go db.HandleKafka()
	//http处理
	router.RouteInit(IpAddress)
	tiker := time.NewTicker(time.Minute * 1)
	for {
		log.Println("执行主go程 ", <-tiker.C)
		log.Println("执行主go程 休息3分钟 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		time.Sleep(time.Minute * 3)
	}

}
