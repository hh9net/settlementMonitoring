package db

import "testing"

//HandleDayTasks()
func TestHandleDayTasks(t *testing.T) {
	Newdb()
	HandleDayTasks()
}

//任务1
func TestQuerTotalSettlementData(t *testing.T) {
	Newdb()
	QuerTotalSettlementData()
}

//任务2
//()
func TestQuerTotalClarify(t *testing.T) {
	Newdb()
	QuerTotalClarify()
}

//QueryblacklistCount()
func TestQueryblacklistCount(t *testing.T) {
	Newdb()
	QueryblacklistCount() //表数64获取黑名单总数:3165万7282"
	//表数64获取黑名单总数:3165万7282
	//表数64获取黑名单总数:3165万9304  有增量包
}

//QueryTingccJieSuan()
func TestQueryTingccJieSuan(t *testing.T) {
	Newdb()
	QueryTingccJieSuan()
}

//QueryClearlingAndDisputePackage
func TestQueryClearlingAndDisputePackage(t *testing.T) {
	Newdb()
	QueryClearlingAndDisputePackage()
}

//DataClassification()
func TestDataClassification(t *testing.T) {
	Newdb()
	DataClassification()
}

//SettlementTrendbyDay()
func TestSettlementTrendbyDay(t *testing.T) {
	Newdb()
	SettlementTrendbyDay()
}

//PacketMonitoring()
func TestPacketMonitoring(t *testing.T) {
	Newdb()
	PacketMonitoring()
}

//ShengnJieSuanData()
func TestShengnJieSuanData(t *testing.T) {
	Newdb()
	ShengnJieSuanData()
}

//QueryShengnRefusePayData()
func TestQueryShengnRefusePayData(t *testing.T) {
	Newdb()
	QueryShengnRefusePayData()
}

//QueryShengnAlreadyPleaseData
func TestQueryShengnAlreadyPleaseData(t *testing.T) {
	Newdb()
	QueryShengnAlreadyPleaseData()
}

//QuerySNDataClassificationData
func TestQuerySNDataClassificationData(t *testing.T) {
	Newdb()
	QuerySNDataClassificationData()
}
