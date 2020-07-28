package service

import (
	"settlementMonitoring/db"
	"testing"
)

//QuerTotalSettlementData()
func TestQuerTotalSettlementData(t *testing.T) {
	db.Newdb()
	QuerTotalSettlementData()
}
