package dto

import "settlementMonitoring/types"

type QuerTotalSettlementDataResponse struct {
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
