package db

import (
	"settlementMonitoring/config"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"time"
)

var GormClient *utils.GormDB
var HmdGormClient *utils.HmdGormDB

//数据库的初始化
func DBInit() {
	GormClient = utils.InitGormDB(&utils.DBConfig{
		DBAddr:       config.Optional.MysqlStr,
		MaxIdleConns: 30,
		LogMode:      utils.Uint8ToBool(config.Optional.DBLog),
	})
}

func HmdDBInit() {
	HmdGormClient = utils.HmdInitGormDB(&utils.HmdDBConfig{
		HmdDBAddr:    config.Optional.MysqlHMDStr,
		MaxIdleConns: 30,
		LogMode:      utils.Uint8ToBool(config.Optional.HmdDBLog),
	})
}

//创建表【测试代码】
func NewTables() {
	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	DBInit() //初始化数据库
	if GormClient.Client.HasTable(&types.JieSuanWssj{}) {
		GormClient.Client.AutoMigrate(&types.JieSuanWssj{})
	} else {
		GormClient.Client.CreateTable(&types.JieSuanWssj{})
	}

	if GormClient.Client.HasTable(&types.JieSuanJiangssj{}) {
		GormClient.Client.AutoMigrate(&types.JieSuanJiangssj{})
	} else {
		GormClient.Client.CreateTable(&types.JieSuanJiangssj{})
	}
}

//
func CreateModel(value interface{}) error {
	if GormClient.Client.NewRecord(value) {
		if mydb := GormClient.Client.Create(value); mydb.Error != nil {
			return mydb.Error
		}
	}
	return nil
}
