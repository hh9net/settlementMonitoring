package db

import (
	"log"
	"settlementMonitoring/types"
	"testing"
)

func TestNewTables(t *testing.T) {
	NewTables()
}
func TestQueryTable(t *testing.T) {
	//数据库生成的表名是结构体名称的复数形式
	s1 := []string{"b_jsjk_jiesjkptyhb",
		"b_jsjk_zhuanjssjjk",
		"b_jsjk_yuqsjtj",
		"b_jsjk_yicsjtj",
		"b_jsjk_yicsjtcctj",
		"b_jsjk_tingccjssjtj",
		"b_jsjk_shujtbjk",
		"b_jsjk_shujbjk",
		"b_jsjk_shengwtccjsqs",
		"b_jsjk_shengwqftj",
		"b_jsjk_shengwjszysjtj",
		"b_jsjk_shengwjssjfl",
		"b_jsjk_shengwjsqs",
		"b_jsjk_shengnyfssjtj",
		"b_jsjk_shengntccjsqs",
		"b_jsjk_shengnsssjjk",
		"b_jsjk_shengnqktj",
		"b_jsjk_shengnjssjfl",
		"b_jsjk_shengnjsqs",
		"b_jsjk_heimdjk",
		"b_jsjk_shengnjfsjtj",
		"b_jsjk_qingfhd",
		"b_jsjk_jiestj",
		"JieSuanWssjs",
	}

	tablenames := make([]string, 0)
	tablenames = append(tablenames, s1...)
	log.Println(len(tablenames), tablenames)
	Newdb()
	for i, tablename := range tablenames {
		log.Println(1+i, tablename)
		QueryTable(tablename)
	}
}

//插入数据
func TestInsertTabledata(t *testing.T) {
	Newdb()
	InsertTabledata()
}

//
func TestQueryTabledata(t *testing.T) {
	Newdb()
	QueryTabledata()
}

func TestUpdateTabledata(t *testing.T) {
	Newdb()
	UpdateTabledata()
}

//QueryJieSuanTable()
func TestQueryJieSuanTable(t *testing.T) {
	Newdb()
	//查询结算表 总交易笔数、总金额
	c, je := QueryJieSuanTable()
	log.Println(c, je)
}

//测试查询各个卡网络号的总金额、总笔数
func TestQueryKawlhJieSuan(t *testing.T) {
	Newdb()
	for _, kawlh := range types.Gl_network {
		c, je := QueryKawlhJieSuan(kawlh)
		log.Println(c, je)
	}
}

//测试查询省内结算总金额、总条数
func TestQueryShengnJieSuan(t *testing.T) {
	Newdb()
	c, je := QueryShengnJieSuan()
	log.Println(c, je)
}

//QueryShengwClearingJieSuan()
func TestQueryShengwClearingJieSuan(t *testing.T) {
	Newdb()
	c, je := QueryShengwClearingJieSuan()
	log.Println(c, je)
}

//QueryDisputeJieSuanData()
func TestQueryDisputeJieSuanData(t *testing.T) {
	Newdb()
	c, je := QueryDisputeJieSuanData()
	log.Println(c, je)
}

//QueryAbnormalData
func TestQueryAbnormalData(t *testing.T) {
	Newdb()
	c, je := QueryAbnormalData()
	log.Println(c, je)
}
