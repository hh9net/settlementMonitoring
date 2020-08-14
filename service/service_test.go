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
	req := dto.ReqQueryClarify{BeginTime: "2020-07-31", EndTime: "2020-08-14", CheckState: "2"}
	log.Println(ClarifyQuery(req))
}
