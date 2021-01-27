package service

import (
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/config"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/utils"
	"testing"
	"time"
)

func TestQuerTotalSettlementData(t *testing.T) {
	db.Newdb()
	QuerTotalSettlementData()
}

func TestQueryDataTurnMonitordata(t *testing.T) {
	db.Newdb()
	QueryDataTurnMonitordata()
}

func TestQuerySettlementTrend(t *testing.T) {
	db.Newdb()
	QuerySettlementTrend()
}

func TestQuerySNRealTimeData(t *testing.T) {
	db.Newdb()
	log.Println(QuerySNRealTimeData())
}

func TestQueryOverdueData(t *testing.T) {
	db.Newdb()
	log.Println(QueryOverdueData())
}

func TestClarifydifference(t *testing.T) {
	db.Newdb()
	log.Println(Clarifydifference())
}

func TestStatisticalClearlingcheck(t *testing.T) {
	db.Newdb()
	log.Println(StatisticalClearlingcheck())
}

//
func TestClarifyQuery(t *testing.T) {
	db.Newdb()
	req := dto.ReqQueryClarify{BeginTime: "2020-08-15", EndTime: "2020-08-19", CheckState: 0, Currentpageid: 1, Prepage: 10}
	log.Println(ClarifyQuery(req))
}

func TestExportExcel(t *testing.T) {
	db.Newdb()
	req := dto.ReqClarifyExportExcel{BeginTime: "2020-08-10", EndTime: "2020-08-22", CheckState: 2, Orderstatus: 1}
	log.Println(ExportExcel(req))
}

func TestQueryHSDZData(t *testing.T) {
	db.NewHSZDDB()
	log.Println(QueryHSDZData())
}

func TestHandleSixHourTasks(t *testing.T) {
	conf := config.ConfigInit() //初始化配置
	log.Error("结算监控配置文件信息：", *conf)
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)
	db.Newdb()
	db.HandleSixHourTasks()
}
