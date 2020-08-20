package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"settlementMonitoring/dto"
	"settlementMonitoring/service"
	"settlementMonitoring/types"
)

/*  接口1方法注释   */
//@Summary 查询省外结算总金额、总笔数api
//@Tags 查询结算总金额、总笔数
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/totalsettlementdata [get]
func QueryTotalSettlementData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询结算总金额、总笔数 处理
	code, err, totaldata := service.QuerTotalSettlementData()
	if err != nil {
		logrus.Errorf("QuerTotalSettlementData  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerTotalSettlementData err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 203 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询结算总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4004, Data: types.StatusText(types.StatusQueryTotalSettlementDataError), Message: "查询结算总金额、总笔数 失败"})
	}
}

/*  接口2方法注释   */
//@Summary 查询省外已清分总金额、总笔数api
//@Tags 查询省外已清分总金额、总笔数
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/totalclarify [get]
func QueryTotalClarify(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省外已清分总金额、总笔数 处理
	code, err, totaldata := service.QuerTotalClarify()
	if err != nil {
		logrus.Errorf("QuerTotalClarify  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerTotalClarify err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 204 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询已清分总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4004, Data: types.StatusText(types.StatusQueryTotalSettlementDataError), Message: "查询已清分总金额、总笔数 失败"})
	}
}

/*  接口3方法注释   */
//@Summary 查询省外坏账总金额、总笔数api
//@Tags 查询省外坏账总金额、总笔数
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/totalBaddebts [get]
func QueryTotalBaddebts(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询坏账 处理
	code, err, totaldata := service.QuerTotalBaddebts()
	if err != nil {
		logrus.Errorf("QueryTotalBaddebts  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryTotalBaddebts err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 205 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询坏账总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4006, Data: types.StatusText(types.StatusQueryTotalBaddebtsError), Message: "查询坏账总金额、总笔数 失败"})
	}
}

/*  接口4方法注释   */
//@Summary 查询省外存在争议总金额、总笔数api
//@Tags 查询省外存在争议总金额、总笔数
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/totaldisputedata [get]
func QueryShengwDisputedata(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询存在争议 处理
	code, err, totaldata := service.QueryDisputedata()
	if err != nil {
		logrus.Errorf("QueryDisputedata  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryDisputedata err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 206 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询存在争议总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4007, Data: types.StatusText(types.StatusQueryShengwDisputedataError), Message: "查询存在争议总金额、总笔数 失败"})
	}
}

/*  接口5方法注释   */
//@Summary 查询异常数据总金额、总笔数api
//@Tags 查询异常数据总金额、总笔数
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/totalAbnormaldata [get]
func QueryAbnormaldata(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询异常数据总金额、总笔数 处理
	code, err, totaldata := service.QueryAbnormaldata()
	if err != nil {
		logrus.Errorf("QueryAbnormaldata  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryAbnormaldata err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 207 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询异常数据总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4008, Data: types.StatusText(types.StatusQueryAbnormaldataError), Message: "查询异常数据总金额、总笔数 失败"})
	}
}

/*  接口6方法注释   */
//@Summary 查询黑名单总数、较2小时前变化值 api
//@Tags 查询黑名单总数、较2小时前变化值
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/totalblacklistdata [get]
func Queryblacklistdata(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询黑名单总数、较2小时前变化值
	code, err, totaldata := service.Queryblacklistdata()
	if err != nil {
		logrus.Errorf("Queryblacklistdata  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Queryblacklistdata err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 208 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询黑名单总数、较2小时前变化值 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4009, Data: types.StatusText(types.StatusQueryblacklistdataError), Message: "查询黑名单总数、较2小时前变化值 失败"})
	}
}

/*  接口7方法注释   */
//@Summary 查询清分包、争议包的接收时间、包号 api
//@Tags 查询清分包、争议包的接收时间、包号
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/clearlingAndDisputePackageSettlement [get]
func QueryClearlingAndDisputePackage(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询清分包、争议包的接收时间、包号
	code, err, totaldata := service.QueryClearlingAndDisputePackagedata()
	if err != nil {
		logrus.Errorf("QueryClearlingAndDisputePackagedata  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryClearlingAndDisputePackagedata err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 209 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询清分包、争议包的接收时间、包号 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4010, Data: types.StatusText(types.StatusQueryClearlingAndDisputePkgError), Message: "查询清分包、争议包的接收时间、包号 失败"})
	}
}

/*  接口8方法注释   */
//@Summary 查询全部清分核对 api
//@Tags 查询全部清分核对
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/clearlingcheck [get]
func Clearlingcheck(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询全部清分核对
	code, err, totaldata := service.StatisticalClearlingcheck()
	if err != nil {
		logrus.Errorf("StatisticalClearlingcheck  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("StatisticalClearlingcheck err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 210 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询清分核对结果 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4011, Data: types.StatusText(types.StatusQueryClearlingcheckOneDataError), Message: "查询清分核对结果 失败"})
	}
}

/*  接口9方法注释   */
//@Summary 省外数据分类 api
//@Tags 省外数据分类
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/dataclassification [get]
func Dataclassification(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省外数据分类
	code, err, totaldata := service.Dataclassification()
	if err != nil {
		logrus.Errorf("QueryDataclassification  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryDataclassification err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 211 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询省外数据分类 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4012, Data: types.StatusText(types.StatusQueryDataclassificationError), Message: "查询省外数据分类 失败"})
	}
}

/*  接口10方法注释   */
//@Summary 省外转结算 api
//@Tags 省外转结算
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/dataturnmonitor [get]
func QueryDataTurnMonitor(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省外转结算
	code, err, totaldata := service.QueryDataTurnMonitordata()
	if err != nil {
		logrus.Errorf("QueryDataTurnMonitordata err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryDataTurnMonitordata err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 212 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询省外转结算 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4013, Data: types.StatusText(types.StatusQueryDataTurnMonitorError), Message: "查询省外转结算 失败"})
	}
}

/*  接口11方法注释   */
//@Summary 省外结算趋势 api
//@Tags 省外结算趋势
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/settlementtrend [get]
func QuerySettlementTrendbyDay(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省外结算趋势
	code, err, totaldata := service.QuerySettlementTrend()
	if err != nil {
		logrus.Errorf("QuerySettlementTrend err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySettlementTrend err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 213 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询省外结算趋势 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4014, Data: types.StatusText(types.StatusQuerySettlementTrendError), Message: "查询省外结算趋势 失败"})
	}
}

/*  接口12方法注释   */
//@Summary 省外数据包监控 api
//@Tags 省外数据包监控
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/packetmonitoring [get]
func PacketMonitoring(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省外数据包监控
	code, err, totaldata := service.QueryPacketMonitoring()
	if err != nil {
		logrus.Errorf("QueryPacketMonitoring err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryPacketMonitoring err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 214 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询省外数据包监控 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4015, Data: types.StatusText(types.StatusQueryPacketMonitoringError), Message: "查询省外数据包监控 失败"})
	}
}

//Clarifydifference

/*  接口13方法注释  */
//@Summary 查询最近15天清分包数据差额 api
//@Tags 查询最近15天清分包数据差额
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/clarifydifference [get]
func Clarifydifference(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	//查询最近15天清分包数据差额
	code, err, totaldata := service.Clarifydifference()
	if err != nil {
		logrus.Errorf("Clarifydifference err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryClarifydifference err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 215 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询最近15天清分包数据差额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4016, Data: types.StatusText(types.StatusQueryClarifydifferenceError), Message: "查询最近15天清分包数据差额 失败"})
	}
}

//ClarifyQuery
/*  接口14方法注释  */
//@Summary 查询清分核对 api
//@Tags 查询清分核对
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/clarifyquery [post]
func ClarifyQuery(c *gin.Context) {
	req := dto.ReqQueryClarify{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("ReqQueryClarify json unmarshal err: %v", err.Error())
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	//查询清分核对
	code, err, totaldata := service.ClarifyQuery(req)
	if err != nil {
		logrus.Errorf("QueryClarify err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryClarify err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 216 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "查询清分核对 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4017, Data: types.StatusText(types.StatusQueryClarifyError), Message: "查询清分核对 失败"})
	}
}

/*  接口15方法注释  */
//@Summary 省外清分核对确认【待处理】 api
//@Tags 省外清分核对确认【待处理】
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/clarifyconfirm [get]
func Clarifyconfirm(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 省外清分核对确认
	code, err, totaldata := service.Clarifyconfirm()
	if err != nil {
		logrus.Errorf("Clarifyconfirmerr: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Clarifyconfirm err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 217 {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: *totaldata, Message: "清分确认成功 【待实现】"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: 4015, Data: types.StatusText(types.StatusQueryPacketMonitoringError), Message: "清分确认 失败 【待实现】"})
	}
}
