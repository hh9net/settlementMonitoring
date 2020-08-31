package types

const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297
	StatusSuccessfully       = 0

	StatusRegisteredSuccessfully                 = 200 //注册成功
	StatusLoginSuccessfully                      = 201 //登录成功
	StatusQuerySWTotalSettlementDataSuccessfully = 203 //查询结算总金额、总笔数 成功
	StatusQueryTotalClarifySuccess               = 204 //查询已清分总金额、总笔数 成功
	StatusQueryTotalBaddebtsSuccess              = 205 //查询坏账总金额、总笔数 成功
	StatusShengwDisputedataSuccess               = 206 //查询存在争议总金额、总笔数 成功
	StatusQueryAbnormaldataSuccess               = 207 //查询异常数据总金额、总笔数 成功
	StatusQueryblacklistdataSuccess              = 208 //查询黑名单总数、较2小时前变化值 成功
	StatusQueryClearlingAndDisputePkgSuccess     = 209 //查询清分包、争议包的接收时间、包号 成功
	StatusQueryClearlingcheckOneDataSuccess      = 210 //查询清分核对结果成功
	StatusQueryDataclassificationSuccess         = 211 //查询省外数据分类 成功
	StatusQueryDataTurnMonitorSuccess            = 212 //查询省外转结算 成功
	StatusQuerySettlementTrendSuccess            = 213 //查询省外结算趋势 成功
	StatusQueryPacketMonitoringSuccess           = 214 //查询省外数据包监控 成功
	StatusQueryClarifydifferenceSuccess          = 215 //查询最近15天清分包数据差额 成功
	StatusQueryClarifySuccess                    = 216 //查询清分核对 失败

	StatusQuerySNTotalSettlementDataSuccess = 301 //查询省内结算数据表的总条数、总金额 成功
	StatusQuerySNSendTotalSettlemenSuccess  = 302 //查询省内的已发送 总条数、总金额 成功
	StatusQuerySNAlreadyPleaseDataSuccess   = 303 //查询省内已请款的数据总条数、总金额 成功
	StatusQuerySNRefusePayDataSuccess       = 304 //查询坏账（拒付）数据 总条数、总金额 成功
	StatusQuerySNRealTimeDataSuccess        = 305 //查询省内结算实时数据监控 成功
	StatusQuerySNSettlementTrendSuccess     = 306 //查询省内前30日省内结算趋势概览 成功
	StatusQueryDataSyncSuccess              = 307 //查询海岭数据同步监控 成功
	StatusQuerySNDataClassificationSuccess  = 308 //查询省内结算数据分类 成功
	StatusQueryAbnormalDataParkingSuccess   = 309 //查询异常数据停车场top10 成功
	StatusQueryOverdueDataSuccess           = 310 //查询逾期数据停车场top10 失败

	StatusRepeatedRegistration = 4001 //注册重复
	StatusPleaseRegister       = 4002 //请先注册
	StatusPasswordError        = 4003 //密码错误,请重新输入

	StatusQueryTotalSettlementDataError    = 4004 //查询结算总金额、总笔数 失败
	StatusQueryTotalClarifyError           = 4005 //查询已清分总金额、总笔数 失败
	StatusQueryTotalBaddebtsError          = 4006 //查询坏账总金额、总笔数 失败
	StatusQueryShengwDisputedataError      = 4007 //查询存在争议总金额、总笔数 失败
	StatusQueryAbnormaldataError           = 4008 //查询异常数据总金额、总笔数 失败
	StatusQueryblacklistdataError          = 4009 //查询黑名单总数、较2小时前变化值 失败
	StatusQueryClearlingAndDisputePkgError = 4010 //查询清分包、争议包的接收时间、包号 失败
	StatusQueryClearlingcheckOneDataError  = 4011 //查询清分核对结果 失败
	StatusQueryDataclassificationError     = 4012 //查询省外数据分类 失败
	StatusQueryDataTurnMonitorError        = 4013 //查询省外转结算 失败
	StatusQuerySettlementTrendError        = 4014 //查询省外结算趋势 失败
	StatusQueryPacketMonitoringError       = 4015 //查询省外数据包监控 失败
	StatusQueryClarifydifferenceError      = 4016 //查询最近15天清分包数据差额 失败
	StatusQueryClarifyError                = 4017 //按条件查询清分核对 失败
	StatusExportExcelError                 = 4018 //导出清分包核对记录表 失败

	StatusQuerySNTotalSettlementDataError = 5001 //查询省内结算数据表的总条数、总金额 失败
	StatusQuerySNSendTotalSettlemenError  = 5002 //查询省内的已发送 总条数、总金额 失败
	StatusQuerySNAlreadyPleaseDataError   = 5003 //查询省内已请款的数据总条数、总金额 失败
	StatusQuerySNRefusePayDataError       = 5004 //查询坏账（拒付）数据 总条数、总金额 失败
	StatusQuerySNRealTimeDataError        = 5005 //查询省内结算实时数据监控 失败
	StatusQuerySNSettlementTrendError     = 5006 //查询省内前30日省内结算趋势概览 失败
	StatusQueryDataSyncError              = 5007 //查询海岭数据同步监控 失败
	StatusQuerySNDataClassificationError  = 5008 //查询省内结算数据分类 失败
	StatusQueryAbnormalDataParkingError   = 5009 //查询异常数据停车场top10 失败
	StatusQueryOverdueDataError           = 5010 //查询逾期数据停车场top10 失败

)

var statusText = map[int]string{
	StatusContinue:                               "Continue",
	StatusSwitchingProtocols:                     "Switching Protocols",
	StatusProcessing:                             "Processing",
	StatusEarlyHints:                             "Early Hints",
	StatusRegisteredSuccessfully:                 "Registered Successfully",
	StatusRepeatedRegistration:                   "Repeated Registration",
	StatusPleaseRegister:                         "Please Register",
	StatusLoginSuccessfully:                      "Login Successfully",
	StatusPasswordError:                          "Password Error",
	StatusQuerySWTotalSettlementDataSuccessfully: "Query Total Settlement Data Success",
	StatusQueryTotalSettlementDataError:          "Query Total Settlement Data Error",
	StatusQueryTotalClarifySuccess:               "Query Total Clarify Success",
	StatusQueryTotalClarifyError:                 "Query Total Clarify Error",
	StatusQueryTotalBaddebtsSuccess:              "Query Total Baddebts Success",
	StatusQueryTotalBaddebtsError:                "Query Total Baddebts Error",
	StatusShengwDisputedataSuccess:               "Query Shengw Dispute data Success",
	StatusQueryShengwDisputedataError:            "Query Shengw Dispute data Error",
	StatusQueryAbnormaldataSuccess:               "Query Abnormal data Success",
	StatusQueryAbnormaldataError:                 "Query Abnormal data Error",
	StatusQueryblacklistdataSuccess:              "Query blacklist data Success",
	StatusQueryblacklistdataError:                "Query blacklist data Error",
	StatusQueryClearlingAndDisputePkgSuccess:     "Query Clearling And Dispute Pkg Success",
	StatusQueryClearlingAndDisputePkgError:       "Query Clearling And Dispute Pkg Error",
	StatusQueryClearlingcheckOneDataSuccess:      "Query Clearling check OneData Success",
	StatusQueryClearlingcheckOneDataError:        "Query Clearling check OneData Error",
	StatusQueryDataclassificationSuccess:         "Query Data classification Success",
	StatusQueryDataclassificationError:           "Query Data classification Error",
	StatusQueryDataTurnMonitorSuccess:            "Query Data Turn Monitor Success",
	StatusQueryDataTurnMonitorError:              "Query Data Turn Monitor Error",
	StatusQuerySettlementTrendSuccess:            "Query Settlement Trend Success",
	StatusQuerySettlementTrendError:              "Query Settlement Trend Error",
	StatusQueryPacketMonitoringSuccess:           "Query Packet Monitoring Success",
	StatusQueryPacketMonitoringError:             "Query Packet Monitoring Error",
	StatusQuerySNTotalSettlementDataSuccess:      "Query SN TotalSettlement Data Success",
	StatusQuerySNTotalSettlementDataError:        "Query SN TotalSettlement Data Error",
	StatusQuerySNSendTotalSettlemenSuccess:       "Query SN Send Total Settlemen Success",
	StatusQuerySNSendTotalSettlemenError:         "Query SN Send Total Settlemen Error",
	StatusQuerySNAlreadyPleaseDataSuccess:        "Query SN Already Please Data Success",
	StatusQuerySNAlreadyPleaseDataError:          "Query SN Already Please Data Error",
	StatusQuerySNRefusePayDataSuccess:            "Query SN Refuse Pay Data Success",
	StatusQuerySNRefusePayDataError:              "Query SN Refuse Pay Data Error",
	StatusQuerySNRealTimeDataSuccess:             "Query SN Real Time Data Success",
	StatusQuerySNRealTimeDataError:               "Query SN Real Time Data Error",
	StatusQuerySNSettlementTrendSuccess:          "Query SN Settlement Trend Success",
	StatusQuerySNSettlementTrendError:            "Query SN Settlement Trend Error",
	StatusQueryDataSyncSuccess:                   "Query Data Sync Success",
	StatusQueryDataSyncError:                     "Query Data Sync Error",
	StatusQuerySNDataClassificationSuccess:       "Query SN Data Classification Success",
	StatusQuerySNDataClassificationError:         "Query SN Data Classification Error",
	StatusQueryAbnormalDataParkingSuccess:        "Query Abnormal Data Parking Success",
	StatusQueryAbnormalDataParkingError:          "Query Abnormal Data Parking Error",
	StatusQueryOverdueDataSuccess:                "Query Overdue Data Success",
	StatusQueryOverdueDataError:                  "Query Overdue Data Error",
	StatusQueryClarifydifferenceSuccess:          "Query Clarify difference Success",
	StatusQueryClarifydifferenceError:            "Query Clarify difference Error",
	StatusQueryClarifySuccess:                    "Query Clarify Success",
	StatusQueryClarifyError:                      "Query Clarify Error",
	StatusSuccessfully:                           "Query Success",
	StatusExportExcelError:                       "Export Excel Error",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
