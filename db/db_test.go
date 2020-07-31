package db

import (
	"github.com/sirupsen/logrus"
	"log"
	"settlementMonitoring/types"
	"testing"
	"time"
)

func TestNewTables(t *testing.T) {
	Newdb()
	//NewTables()
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
	InsertTabledata(10000)
}

//
func TestQueryTabledata(t *testing.T) {
	Newdb()
	QueryTabledata(10000)
}

func TestUpdateTabledata(t *testing.T) {
	Newdb()
	data := &types.BJsjkJiestj{FNbId: 19, FNbZongts: 9999, FNbZongje: 90000}
	UpdateTabledata(data, 10000, 19)
}

//QueryJieSuanTable()
func TestQueryJieSuanTable(t *testing.T) {
	Newdb()
	//查询结算表 总交易笔数、总金额
	c, je := QueryJieSuanTable()
	log.Println(c, je) //查询结算表总交易笔数为8058， 查询总金额为：11933100"
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
	log.Println(c, je) //结算表总交易笔数7011，查询总金额为：[10309200]
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

//查询待处理的异常数据 总条数、总金额【单点+总对总】
func TestQueryAbnormalData(t *testing.T) {
	Newdb()
	c, je, err := QueryAbnormalData(1)
	log.Println(c, je, err)
}

// 查询 已清分的坏账 Bad debts
func TestQueryShengwBadDebtsJieSuan(t *testing.T) {
	Newdb()
	c, je := QueryShengwBadDebtsJieSuan()
	log.Println(c, je)
}

//(省外总金额)
func TestQueryShengwJieSuan(t *testing.T) {
	Newdb()
	c, je := QueryShengwJieSuan()
	log.Println(c, je) //总交易笔数1047，查询总金额为：[1623900]
}

//ShengwClearingInsert()
func TestShengwClearingInsert(t *testing.T) {
	Newdb()
	c := ShengwClearingInsert()
	log.Println(c)
}

//QueryJieSuan()
func TestQueryJieSuan(t *testing.T) {
	Newdb()
	QueryJieSuan()
}

//QueryShengwClearingdata()
func TestQueryShengwClearingdata(t *testing.T) {
	Newdb()
	QueryShengwClearingdata()
}

//ShengwDisputeInsert()
func TestShengwDisputeInsert(t *testing.T) {
	Newdb()
	ShengwDisputeInsert()
}

//QueryShengwDispute
func TestQueryShengwDispute(t *testing.T) {
	Newdb()
	QueryShengwDispute()
}

//AbnormalDataInsert()
func TestAbnormalDataInsert(t *testing.T) {
	Newdb()
	AbnormalDataInsert(1)
}

//QueryAbnormaltable()
func TestQueryAbnormaltable(t *testing.T) {
	Newdb()
	QueryAbnormaltable(1)
}

//QueryblacklistTable
func TestQueryblacklistTable(t *testing.T) {
	Newdb()
	QueryblacklistTable() //64张表
}

//QueryBlacklistcount
func TestQueryBlacklistcount(t *testing.T) {
	//config.InitConfigs()  //初始化配置
	QueryBlacklistcount() //"表数64获取黑名单总数:3165万6389"
	//"表数64获取黑名单总数:3165万6389"
}

//QueryBlacklistTiaoshutable
func TestQueryBlacklistTiaoshutable(t *testing.T) {
	Newdb()
	QueryBlacklistTiaoshutable(7, 3)
}

func TestQueryTingccJieSuandata(t *testing.T) {
	Newdb()
	c := QueryTingccJieSuandata()
	log.Println(c)
}

//InsertTingjiesuan
func TestInsertTingjiesuan(t *testing.T) {
	Newdb()
	InsertTingjiesuan()
}

//QueryTingjiesuan
func TestQueryTingjiesuan(t *testing.T) {
	Newdb()
	QueryTingjiesuan()
}

//UpdateTingjiesuan
func TestUpdateTingjiesuan(t *testing.T) {
	Newdb()

	err := UpdateTingjiesuan(&types.BJsjkTingccjssjtj{
		FNbZongje:    1314,        //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
		FNbZongts:    1212,        //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
		FDtTongjwcsj: time.Now(),  //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
		FVcTongjrq:   "2020-7-28", //   `F_VC_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
		FVcTingccid:  "20201314",  //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	}, "20201314", 1)
	logrus.Print(err)
}

// QueryTingjiesuanById
func TestQueryTingjiesuanById(t *testing.T) {
	Newdb()
	QueryTingjiesuanById(2)
}

//QueryClearlingdata()
func TestQueryClearlingdata(t *testing.T) {
	Newdb()
	log.Println(QueryClearlingdata("2020-07-30"))
}

//StatisticalkeepAccount
func TestStatisticalkeepAccount(t *testing.T) {
	Newdb()
	StatisticalkeepAccount() //41400 【30日】
}

//QueryDisputedata
func TestQueryDisputedata(t *testing.T) {
	Newdb()
	log.Println(QueryDisputedata("2020-07-30"))

}

//DisputedDataCanClearling
func TestDisputedDataCanClearling(t *testing.T) {
	Newdb()
	log.Println(DisputedDataCanClearling(317671)) //9550

}

//StatisticalClearlingcheck()
func TestStatisticalClearlingcheck(t *testing.T) {
	Newdb()
	log.Println(StatisticalClearlingcheck())

}
