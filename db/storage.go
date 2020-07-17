package db

import (
	"github.com/sirupsen/logrus"
	"settlementMonitoring/config"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"time"
)

//结算监控平台数据层：数据的增删改查
func Newdb() {
	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	DBInit() //初始化数据库
}

//1、查询表是否存在
func QueryTable(tablename string) {
	db := utils.GormClient.Client
	is := db.HasTable(tablename)

	if is == false {
		logrus.Println("不存在", tablename)
		return
	}
	logrus.Println("表存在：", tablename, is)
}

//2、Update  b_jsjk_jiestj
func UpdateTabledata() {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkJiestj)
	Jiestj.FDtKaistjsj = "2020-02-20 12:12:12"
	db.Model("b_jsjk_jiestj").Update(Jiestj)

}
