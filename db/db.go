package db

import (
	"settlementMonitoring/utils"
)

var GormClient *utils.GormDB
var HmdGormClient *utils.HmdGormDB

//数据库的初始化
func DBInit(mstr string) {
	GormClient = utils.InitGormDB(&utils.DBConfig{
		DBAddr:       mstr,
		MaxIdleConns: 30,
		LogMode:      utils.Uint8ToBool(1),
	})
}

func HmdDBInit() {
	HmdGormClient = utils.HmdInitGormDB(&utils.HmdDBConfig{
		HmdDBAddr:    "root:Microvideo_1@tcp(122.51.24.189:3307)/blacklist?charset=utf8&parseTime=true&loc=Local",
		MaxIdleConns: 30,
		LogMode:      utils.Uint8ToBool(1),
	})
}
