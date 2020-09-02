package db

import (
	"settlementMonitoring/types"
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
		HmdDBAddr:    types.HmdDBAddrconf,
		MaxIdleConns: 30,
		LogMode:      utils.Uint8ToBool(1),
	})
}
