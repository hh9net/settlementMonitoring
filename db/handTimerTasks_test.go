package db

import "testing"

func TestHandleDayTasks(t *testing.T) {
	Newdb()
	HandleDayTasks()
}

func TestQuerTotalSettlementData(t *testing.T) {
	Newdb()
	QuerTotalSettlementData()
}

func TestQuerTotalClarify(t *testing.T) {
	Newdb()
	QuerTotalClarify()
}

func TestQueryblacklistCount(t *testing.T) {
	Newdb()
	QueryblacklistCount() //表数64获取黑名单总数:3165万7282"
	//表数64获取黑名单总数:3165万7282
	//表数64获取黑名单总数:3165万9304  有增量包
}

func TestQueryTingccJieSuan(t *testing.T) {
	Newdb()
	QueryTingccJieSuan()
}

func TestQueryClearlingAndDisputePackage(t *testing.T) {
	Newdb()
	QueryClearlingAndDisputePackage()
}

func TestDataClassification(t *testing.T) {
	Newdb()
	DataClassification()
}

func TestSettlementTrendbyDay(t *testing.T) {
	Newdb()
	SettlementTrendbyDay()
}

func TestPacketMonitoring(t *testing.T) {
	Newdb()
	PacketMonitoring()
}

func TestShengnJieSuanData(t *testing.T) {
	Newdb()
	ShengnJieSuanData()
}

func TestQueryShengnRefusePayData(t *testing.T) {
	Newdb()
	QueryShengnRefusePayData()
}

func TestQueryShengnAlreadyPleaseData(t *testing.T) {
	Newdb()
	QueryShengnAlreadyPleaseData()
}

func TestQuerySNDataClassificationData(t *testing.T) {
	Newdb()
	QuerySNDataClassificationData()
}

func TestShengNRealTimeSettlementData(t *testing.T) {
	Newdb()
	ShengNRealTimeSettlementData()
}
func TestQueryShengNSettlementTrenddata(t *testing.T) {
	Newdb()
	QueryShengNSettlementTrenddata()
}

func TestQueryDataSyncdata(t *testing.T) {
	Newdb()
	//QueryDataSyncdata()
}

func TestQueryAbnormalDataOfParkingdata(t *testing.T) {
	Newdb()
	QueryAbnormalDataOfParkingdata()
}

func TestOverduedata(t *testing.T) {
	Newdb()
	Overduedata()
}

func TestSWSettlementTrendbyDay(t *testing.T) {
	Newdb()
	SWSettlementTrendbyDay()
}
