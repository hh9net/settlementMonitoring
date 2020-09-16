package service

import (
	"log"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"testing"
)

//QuerTotalSettlementData()
func TestQuerTotalSettlementData(t *testing.T) {
	db.Newdb()
	QuerTotalSettlementData()
}

// QueryDataTurnMonitordata()
func TestQueryDataTurnMonitordata(t *testing.T) {
	db.Newdb()
	QueryDataTurnMonitordata()
}

// QuerySettlementTrend()
func TestQuerySettlementTrend(t *testing.T) {
	db.Newdb()
	QuerySettlementTrend()
}

//QuerySNRealTimeData
func TestQuerySNRealTimeData(t *testing.T) {
	db.Newdb()
	log.Println(QuerySNRealTimeData())
}

//QueryOverdueData
func TestQueryOverdueData(t *testing.T) {
	db.Newdb()
	log.Println(QueryOverdueData())
}

// Clarifydifference()
func TestClarifydifference(t *testing.T) {
	db.Newdb()
	log.Println(Clarifydifference())
}

//StatisticalClearlingcheck
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

//ExportExcel
func TestExportExcel(t *testing.T) {
	db.Newdb()
	req := dto.ReqClarifyExportExcel{BeginTime: "2020-08-10", EndTime: "2020-08-22", CheckState: 2, Orderstatus: 1}
	log.Println(ExportExcel(req))
}

//QueryHSDZData()
func TestQueryHSDZData(t *testing.T) {
	db.NewHSZDDB()
	log.Println(QueryHSDZData())
}
