package service

import (
	"log"
	"settlementMonitoring/db"
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
