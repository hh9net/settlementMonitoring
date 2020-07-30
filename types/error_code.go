package types

const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusRegisteredSuccessfully                 = 200 //注册成功
	StatusLoginSuccessfully                      = 201 //登录成功
	StatusQuerySWTotalSettlementDataSuccessfully = 203 //查询结算总金额、总笔数 成功
	StatusQueryTotalClarifySuccess               = 204 //查询已清分总金额、总笔数 成功
	StatusQueryTotalBaddebtsSuccess              = 205 //查询坏账总金额、总笔数 成功
	StatusShengwDisputedataSuccess               = 206 //查询存在争议总金额、总笔数 成功
	StatusQueryAbnormaldataSuccess               = 207 //查询异常数据总金额、总笔数 成功
	StatusQueryblacklistdataSuccess              = 208 //查询黑名单总数、较2小时前变化值 成功
	StatusQueryClearlingAndDisputePkgSuccess     = 209 //查询清分包、争议包的接收时间、包号 成功

	StatusRepeatedRegistration             = 401 //注册重复
	StatusPleaseRegister                   = 402 //请先注册
	StatusPasswordError                    = 403 //密码错误,请重新输入
	StatusQueryTotalSettlementDataError    = 404 //查询结算总金额、总笔数 失败
	StatusQueryTotalClarifyError           = 405 //查询已清分总金额、总笔数 失败
	StatusQueryTotalBaddebtsError          = 406 //查询坏账总金额、总笔数 失败
	StatusQueryShengwDisputedataError      = 407 //查询存在争议总金额、总笔数 失败
	StatusQueryAbnormaldataError           = 408 //查询异常数据总金额、总笔数 失败
	StatusQueryblacklistdataError          = 409 //查询黑名单总数、较2小时前变化值 失败
	StatusQueryClearlingAndDisputePkgError = 410 //查询清分包、争议包的接收时间、包号 失败
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
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
