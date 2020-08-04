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
		c.JSON(http.StatusOK, dto.QuerTotalSettlementDataResponse{Code: code, CodeMsg: types.StatusText(203), Data: *totaldata, Message: "查询结算总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryTotalSettlementDataError), Message: "查询结算总金额、总笔数 失败"})
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
		c.JSON(http.StatusOK, dto.QuerTotalSettlementDataResponse{Code: code, CodeMsg: types.StatusText(204), Data: *totaldata, Message: "查询已清分总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryTotalSettlementDataError), Message: "查询已清分总金额、总笔数 失败"})
	}
}

//totalBaddebts 坏账
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
		c.JSON(http.StatusOK, dto.QuerTotalSettlementDataResponse{Code: code, CodeMsg: types.StatusText(205), Data: *totaldata, Message: "查询坏账总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryTotalBaddebtsError), Message: "查询坏账总金额、总笔数 失败"})
	}
}

//totaldisputedata 存在争议
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
		c.JSON(http.StatusOK, dto.QuerTotalSettlementDataResponse{Code: code, CodeMsg: types.StatusText(206), Data: *totaldata, Message: "查询存在争议总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryShengwDisputedataError), Message: "查询存在争议总金额、总笔数 失败"})
	}
}

//QueryAbnormaldata 异常数据统计
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
		c.JSON(http.StatusOK, dto.QuerTotalSettlementDataResponse{Code: code, CodeMsg: types.StatusText(207), Data: *totaldata, Message: "查询异常数据总金额、总笔数 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryAbnormaldataError), Message: "查询异常数据总金额、总笔数 失败"})
	}
}

//Queryblacklistdata
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
		c.JSON(http.StatusOK, dto.QuerTotalSettlementDataResponse{Code: code, CodeMsg: types.StatusText(208), Data: *totaldata, Message: "查询黑名单总数、较2小时前变化值 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryblacklistdataError), Message: "查询黑名单总数、较2小时前变化值 失败"})
	}
}

//clearlingAndDisputePackageSettlement
/*  接口6方法注释   */
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
		c.JSON(http.StatusOK, dto.QueryClearlingAndDisputeResponse{Code: code, CodeMsg: types.StatusText(209), Data: *totaldata, Message: "查询清分包、争议包的接收时间、包号 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryClearlingAndDisputePkgError), Message: "查询清分包、争议包的接收时间、包号 失败"})
	}
}

//StatisticalClearlingcheck
/*  接口6方法注释   */
//@Summary 清分核对 api
//@Tags 清分核对
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/clearlingcheck [get]
func Clearlingcheck(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询清分包、争议包的接收时间、包号
	code, err, totaldata := service.StatisticalClearlingcheck()
	if err != nil {
		logrus.Errorf("StatisticalClearlingcheck  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("StatisticalClearlingcheck err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 210 {
		c.JSON(http.StatusOK, dto.QueryClearlingAndDisputeResponse{Code: code, CodeMsg: types.StatusText(210), Data: *totaldata, Message: "查询清分核对结果 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryClearlingcheckOneDataError), Message: "查询清分核对结果 失败"})
	}
}

//Dataclassification
/*  接口6方法注释   */
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
		c.JSON(http.StatusOK, dto.QueryClearlingAndDisputeResponse{Code: code, CodeMsg: types.StatusText(211), Data: *totaldata, Message: "查询省外数据分类 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryDataclassificationError), Message: "查询省外数据分类 失败"})
	}
}

//QueryDataTurnMonitor
/*  接口6方法注释   */
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
		c.JSON(http.StatusOK, dto.QueryClearlingAndDisputeResponse{Code: code, CodeMsg: types.StatusText(212), Data: *totaldata, Message: "查询省外转结算 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryDataTurnMonitorError), Message: "查询省外转结算 失败"})
	}
}

//QuerySettlementTrendbyDay
/*  接口6方法注释   */
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
		c.JSON(http.StatusOK, dto.QueryClearlingAndDisputeResponse{Code: code, CodeMsg: types.StatusText(213), Data: *totaldata, Message: "查询省外结算趋势 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQueryDataTurnMonitorError), Message: "查询省外结算趋势 失败"})
	}
}
