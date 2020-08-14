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
	Amount string `json:"amount"  example:"1244547.00"`
	Count  int    `json:"count"  example:"42626"`
}
type TotalBaddebtsData struct {
	Amount string `json:"amount"  example:"1244547.00"`
	Count  int    `json:"count"  example:"42626"`
}
type TotalDisputeData struct {
	Amount string `json:"amount"  example:"1244547.00"`
	Count  int    `json:"count"  example:"42626"`
}
type TotalAbnormalData struct {
	Amount string `json:"amount"  example:"1244547.00"`
	Count  int    `json:"count"  example:"42626"`
}

type TotalBlacklistData struct {
	Blacklistcount int `json:"blacklistcount"  example:"1244547"`
	ChangeCount    int `json:"changecount"  example:"42626"`
}

//ClearlingAndDisputeData
//type ClearlingAndDisputeData struct {
//	ClearPacgNo    string `json:"clearpacgno" example:"1244547"`
//	Cleardatetime  string `json:"cleardatetime"example:"1244547"`
//	DisputPacgeNo  string `json:"disputpacgeno"example:"1244547"`
//	Disputdatetime string `json:"disputdatetime"example:"1244547"`
//}
type ClearlAndDisputeData struct {
	ClearlingAndDisputedata []types.ClearlingAndDisputeData `json:"clearlinganddisputedata" example:"1244547"`
}

//Clearlingcheck
type ClearlingcheckData struct {
	Clearlingpakgxh int64  `json:"clearlingpakgxh" example:"1244547"`
	Clearlingpakgje int64  `json:"clearlingpakgje" example:"1244547"`
	Clearlingpakgts int    `json:"clearlingpakgts" example:"1244547"`
	Tongjqfje       int64  `json:"tongjqfje" example:"1244547"`
	Tongjqfts       int    `json:"tongjqfts" example:"1244547"`
	Hedjg           int    `json:"hedjg" example:"1244547"`
	Tongjrq         string `json:"tongjrq" example:"1244547"`
	Qingfbjssj      string `json:"qingfbjssj" example:"1244547"`
}

type Dataclassification struct {
	Shengwzcount int `json:"shengwzcount" example:"1244547"` //省外结算总数据
	Yiqfcount    int `json:"yiqfcount" example:"1244547"`    //已清分总条数（不含坏账）
	Jizcount     int `json:"jizcount " example:"1244547"`    //记账
	Zhengycount  int `json:"zhengycount" example:"1244547"`  //争议
	Weidbcount   int `json:"weidbcount" example:"1244547"`   //未打包
	Yidbcount    int `json:"yidbcount" example:"1244547"`    //已打包
	Yifscount    int `json:"yifscount " example:"1244547"`   //已发送
	Huaizcount   int `json:"huaizcount" example:"1244547"`   //坏账
}

type TurnData struct {
	Jieszcount int //结算表总数
	DDzcount   int //单点出口总笔数
	ZDZcount   int //总对总总笔数
}

type TurnDataResponse struct {
	JieszCount  int `json:"jieszcount " example:"1244547"` //结算表总数
	YuansCount  int `json:"ddcount" example:"1244547"`     //原始出口总笔数
	DifferCount int `json:"differcount" example:"1244547"`
}

//省外结算趋势
type SettlementTrend struct {
	JiesuanAmount int64 `json:"jiesuanamount " example:"1244547"` //今日结算金额
	QingfAmount   int64 `json:"qingfamount" example:"1244547"`    //今日清分总笔数
	DifferAmount  int64 `json:"differamount" example:"1244547"`   //差额
	QingfCount    int   `json:"qingfcount" example:"1244547"`     //今日清分
	JiesuanCount  int   `json:"jiesuancount" example:"1244547"`   //今日结算条数
}

type PacketMonitoringdata struct {
	Dabaosl   int   `json:"dabaosl" example:"1244547"`   //今日打包数量
	Dabaojine int64 `json:"dabaojine" example:"1244547"` //打包金额
	Fasbsl    int   `json:"fasbsl" example:"1244547"`    //已发送原始交易消息包数量
	Fasbjine  int64 `json:"fasbjine" example:"1244547"`  //已发送原始交易消息包金额
	Jizbsl    int   `json:"jizbsl" example:"1244547"`    //记账包数量
	Jizbjine  int64 `json:"jizbjine" example:"1244547"`  //记账包金额
	Yuansbsl  int   `json:"yuansbsl" example:"1244547"`  //原始交易消息应答包数量
}

type DifferAmount struct {
	Differamount int64 `json:"differamount" example:"1244547"` //差额
}

type ReqQueryClarify struct {
	BeginTime  string `json:"BeginTime" example:"2020-08-13 13:13:13"`
	EndTime    string `json:"EndTime" example:"2020-08-13 13:13:13"`
	CheckState string `json:"CheckState" example:"1"` //0:全部、1：校验成功 2：校验失败
}
