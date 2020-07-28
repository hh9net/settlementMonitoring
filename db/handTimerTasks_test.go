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
