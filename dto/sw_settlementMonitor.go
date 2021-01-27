package dto

import (
	"settlementMonitoring/types"
)

type QuerTotalSettlementDataResponse struct {
	Code    int `json:"code"  example:"200"`
	CodeMsg string
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}

type QuerResponse struct {
	Code    int `json:"code"  example:"200"`
	CodeMsg string
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}

//QueryClearlingAndDisputePackage
type QueryClearlingAndDisputeResponse struct {
	Code    int `json:"code"  example:"200"`
	CodeMsg string
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}
type TotalSettlementData struct {
	Amount string `json:"amount"  example:"1244547.00"`
	Count  int    `json:"count"  example:"42626"`
}

type TotalClarifyData struct {
	Amount   string `json:"amount"  example:"1244547.00"`
	Count    int    `json:"count"  example:"42626"`
	DateTime string `json:"datetime"  example:"2020-08-18"` //完成时间

}
type TotalBaddebtsData struct {
	Amount   string `json:"amount"  example:"1244547.00"`
	Count    int    `json:"count"  example:"42626"`
	DateTime string `json:"datetime"  example:"2020-08-18"` //完成时间

}
type TotalDisputeData struct {
	Amount   string `json:"amount"  example:"1244547.00"`
	Count    int    `json:"count"  example:"42626"`
	DateTime string `json:"datetime"  example:"2020-08-18"` //完成时间

}
type TotalAbnormalData struct {
	Amount   string `json:"amount"  example:"1244547.00"`
	Count    int    `json:"count"  example:"42626"`
	DateTime string `json:"datetime"  example:"2020-08-18"` //完成时间

}

type TotalBlacklistData struct {
	Blacklistcount int    `json:"blacklist_count"  example:"1244547"`
	ChangeCount    int    `json:"change_count"  example:"42626"`
	DateTime       string `json:"datetime"  example:"2020-08-18"` //完成时间
	Blacklist      []BlackList
}

type BlackList struct {
	Blacklistcount int    `json:"blacklist_count"  example:"1244547"`
	DateTime       string `json:"datetime"  example:"2020-08-18"` //完成时间

}

//ClearlingAndDisputeData
//type ClearlingAndDisputeData struct {
//	ClearPacgNo    string `json:"clearpacgno" example:"1244547"`
//	Cleardatetime  string `json:"cleardatetime"example:"1244547"`
//	DisputPacgeNo  string `json:"disputpacgeno"example:"1244547"`
//	Disputdatetime string `json:"disputdatetime"example:"1244547"`
//}
type ClearlAndDisputeData struct {
	ClearlingAndDisputedata []types.ClearlingAndDisputeData `json:"clearling_and_dispute_data" example:"1244547"`
}

//Clearlingcheck

type Clearlingcheckdata struct {
	Clearlingcheck []ClearlingcheckData
	ZongTS         int `json:"total_count" example:"1244547"` //总记录条数
	ZongYS         int `json:"page_num" example:"1244547"`    //总页数
}
type ClearlingcheckData struct {
	Clearlingpakgxh int64  `json:"clear_pakg_id" example:"1244547"`    //清分包序号
	Clearlingpakgje string `json:"clear_pakg_money" example:"1244547"` //清分包金额
	Clearlingpakgts int    `json:"clear_pakg_num" example:"1244547"`   //清分包条数
	Tongjqfje       string `json:"clear_money" example:"1244547"`      //统计清分金额
	Tongjqfts       int    `json:"clear_num" example:"1244547"`        //统计清分条数
	Hedjg           int    `json:"result" example:"1244547"`           //核对结果 1：核对一致
	Tongjrq         string `json:"datatime" example:"1244547"`         //统计日期
	Qingfbjssj      string `json:"receiv_time" example:"1244547"`      //清分包接收时间
	Tuifje          string `json:"refund_money" example:"1244547"`
	Tuifts          int    `json:"refund_count"`
}

type Dataclassification struct {
	Shengwzcount   int    `json:"total_count" example:"1244547"`        //省外结算总数据
	Yiqfcount      int    `json:"clear_count" example:"1244547"`        //已清分总条数（不含坏账）
	Jizcount       int    `json:"keepaccount_count" example:"1244547"`  //记账
	Zhengycount    int    `json:"dispute_count" example:"1244547"`      //争议
	Weidbcount     int    `json:"no_packaging_count" example:"1244547"` //未打包
	Yidbcount      int    `json:"packaging_count" example:"1244547"`    //已打包
	Yifscount      int    `json:"send_count" example:"1244547"`         //已发送
	Huaizcount     int    `json:"bad_debts_count" example:"1244547"`    //坏账
	DateTime       string `json:"datetime"  example:"2020-08-18"`       //完成时间
	FNbShengwjyzje string `json:"total_money" example:"1244547"`        //`F_NB_SHENGWJYZJE` bigint(20) DEFAULT NULL COMMENT '省外交易总金额（单位：元）',
	FNbShengwqfje  string `json:"clear_money" example:"1244547"`        //`F_NB_SHENGWQFJE` bigint(20) DEFAULT NULL COMMENT '省外清分金额（单位：元）',
	FNbShengwhzje  string `json:"bad_debts_money" example:"1244547"`    //`F_NB_SHENGWHZJE` bigint(20) DEFAULT NULL COMMENT '省外坏账金额元',

}

type TurnData struct {
	Jieszcount int    //结算表总数
	DDzcount   int    //单点出口总笔数
	ZDZcount   int    //总对总总笔数
	DateTime   string //统计时间
}

type TurnDataResponse struct {
	JieszCount  int    `json:"total_count" example:"1244547"` //结算表总数
	YuansCount  int    `json:"count" example:"1244547"`       //原始出口总笔数
	DifferCount int    `json:"differ_count" example:"1244547"`
	DateTime    string `json:"datetime"  example:"2020-08-18"` //完成时间

}

//省外结算趋势
type SettlementTrend struct {
	JiesuanAmount string `json:"amount" example:"1244547"`        //今日结算金额
	QingfAmount   string `json:"clear_amount" example:"1244547"`  //今日清分总笔数
	DifferAmount  string `json:"differ_amount" example:"1244547"` //差额
	QingfCount    int    `json:"clear_count" example:"1244547"`   //今日清分
	JiesuanCount  int    `json:"count" example:"1244547"`         //今日结算条数
	DateTime      string `json:"datetime"  example:"2020-08-18"`  //完成时间

}

type PacketMonitoringdata struct {
	Dabaosl       int    `json:"package_number" example:"1244547"`             //今日打包数量
	Dabaojine     string `json:"package_amount" example:"1244547"`             //打包金额
	Fasbsl        int    `json:"send_package_number" example:"1244547"`        //已发送原始交易消息包数量
	Fasbjine      string `json:"send_package_amount" example:"1244547"`        //已发送原始交易消息包金额
	Jizbsl        int    `json:"keepaccount_package_number" example:"1244547"` //记账包数量
	Jizbjine      string `json:"keepaccount_package_amount" example:"1244547"` //记账包金额
	Yuansyingdbsl int    `json:"reply_package_number" example:"1244547"`       //原始交易消息应答包数量
	DateTime      string `json:"datetime"  example:"2020-08-18"`               //完成时间

}

type DifferAmount struct {
	Differamount string `json:"differ_amount" example:"1244547"` //差额
	DateTime     string `json:"datetime"  example:"2020-08-18"`  //完成时间

}

type ReqQueryClarify struct {
	BeginTime     string `json:"begin_time" example:"2020-08-13 13:13:13"`
	EndTime       string `json:"end_time" example:"2020-08-13 13:13:13"`
	CheckState    int    `json:"check_state" example:"1"` //0:全部、1：校验成功 2：校验失败
	Currentpageid int    `json:"currentpageid" example:"1"`
	Prepage       int    `json:"count" example:"1"`
	Orderstatus   int    `json:"order_status" example:"1"`
}

type ReqQuery struct {
	BeginTime string `json:"begin_time" example:"2020-08-13"`
}

type ReqClarifyExportExcel struct {
	BeginTime   string `json:"begin_time" example:"2020-08-13 13:13:13"`
	EndTime     string `json:"end_time" example:"2020-08-13 13:13:13"`
	CheckState  int    `json:"check_state" example:"1"` //0:全部、1：校验成功 2：校验失败
	Orderstatus int    `json:"order_status" example:"1"`
}

//{
//"begintime":"2020-07-31",
//"endtime":"2020-08-14",
//"checkstate":"2"
//"currentpageid":"2",
//"prepage":"2",
//}
