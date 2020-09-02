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

//省内结算监控模块

/*  接口1方法注释   */
//@Summary 省内结算数据表的总条数、总金额 api
//@Tags 省内结算数据表的总条数、总金额
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/totalsettlementdata [get]
func QuerySNTotalSettlementData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省外数据包监控
	code, err, totaldata := service.QuerSNTotalSettlementData()
	if err != nil {
		logrus.Errorf("QuerSNTotalSettlementData err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerSNTotalSettlementData err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 301 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省内结算数据表的总条数、总金额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNTotalSettlementDataError, Data: types.StatusText(types.StatusQuerySNTotalSettlementDataError), Message: "查询省内结算数据表的总条数、总金额 失败"})
	}
}

//QuerySNSendTotalSettlementData
/*  接口2方法注释   */
//@Summary 查询省内的已发送 总条数、总金额 api
//@Tags 查询省内的已发送 总条数、总金额
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/sendtotalsettlementdata [get]
func QuerySNSendTotalSettlementData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省内的已发送 总条数、总金额
	code, err, totaldata := service.QuerySNSendTotalSettlemen()
	if err != nil {
		logrus.Errorf("QuerySNSendTotalSettlemen err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySNSendTotalSettlemen err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 302 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省内的已发送 总条数、总金额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNTotalSettlementDataError, Data: types.StatusText(types.StatusQuerySNSendTotalSettlemenError), Message: "查询省内的已发送 总条数、总金额 失败"})
	}
}

//QuerySNAlreadyPleaseData
/*  接口3方法注释   */
//@Summary 查询省内已请款的数据总条数、总金额 api
//@Tags 查询省内已请款的数据总条数、总金额
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/alreadyplease [get]
func QuerySNAlreadyPleaseData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省内已请款的数据总条数、总金额
	code, err, totaldata := service.QuerySNAlreadyPleaseData()
	if err != nil {
		logrus.Errorf("QuerySNAlreadyPleaseData err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySNAlreadyPleaseData err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 303 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省内已请款的数据总条数、总金额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNAlreadyPleaseDataError, Data: types.StatusText(types.StatusQuerySNAlreadyPleaseDataError), Message: "查询省内已请款的数据总条数、总金额 失败"})
	}
}

//QuerySNRefusePayData
/*  接口4方法注释   */
//@Summary 查询坏账（拒付）数据 总条数、总金额 api
//@Tags 查询坏账（拒付）数据 总条数、总金额
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/refusepay [get]
func QuerySNRefusePayData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询坏账（拒付）数据 总条数、总金额
	code, err, totaldata := service.QuerySNRefusePayData()
	if err != nil {
		logrus.Errorf("QuerySNRefusePayData err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySNRefusePayData err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 304 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询坏账（拒付）数据 总条数、总金额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNRefusePayDataError, Data: types.StatusText(types.StatusQuerySNRefusePayDataError), Message: "查询坏账（拒付）数据 总条数、总金额 失败"})
	}
}

/*  接口5方法注释   */
//@Summary 省内结算实时数据监控 api
//@Tags 查询省内结算实时数据监控
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/realtimedata [get]
func QuerySNRealTimeData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询省内结算实时数据监控
	code, err, totaldata := service.QuerySNRealTimeData()
	if err != nil {
		logrus.Errorf("QuerySNRealTimeData err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySNRealTimeData err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 305 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省内结算实时数据监控 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNRealTimeDataError, Data: types.StatusText(types.StatusQuerySNRealTimeDataError), Message: "查询省内结算实时数据监控 失败"})
	}
}

//QuerySNSettlementTrend
/*  接口6方法注释   */
//@Summary 前30日省内结算趋势概览 api
//@Tags 查询前30日省内结算趋势概览
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/settlementtrend [get]
func QuerySNSettlementTrend(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询前30日省内结算趋势概览
	code, err, totaldata := service.QuerySNSettlementTrend()
	if err != nil {
		logrus.Errorf("QuerySNSettlementTrend err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySNSettlementTrend err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 306 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省内前30日省内结算趋势概览 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNSettlementTrendError, Data: types.StatusText(types.StatusQuerySNSettlementTrendError), Message: "查询省内前30日省内结算趋势概览 失败"})
	}
}

//QueryDataSync
/*  接口7方法注释   */
//@Summary 海岭数据同步监控 api
//@Tags 查询海岭数据同步监控
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/datasync [get]
func QueryDataSync(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	// 查询海岭数据同步监控
	code, err, totaldata := service.QueryDataSync()
	if err != nil {
		logrus.Errorf("QueryDataSyncerr: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryDataSync err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 307 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询海岭数据同步监控 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataSyncError, Data: types.StatusText(types.StatusQueryDataSyncError), Message: "查询海岭数据同步监控 失败"})
	}
}

//QuerySNDataClassification
/*  接口8方法注释   */
//@Summary 省内结算数据分类 api
//@Tags 查询省内结算数据分类
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/dataclassification [get]
func QuerySNDataClassification(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	//  查询省内结算数据分类
	code, err, totaldata := service.QuerySNDataClassification()
	if err != nil {
		logrus.Errorf("QuerySNDataClassification err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QuerySNDataClassification err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 308 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: " 查询省内结算数据分类 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySNDataClassificationError, Data: types.StatusText(types.StatusQuerySNDataClassificationError), Message: " 查询省内结算数据分类 失败"})
	}
}

//QueryAbnormalDataParking
//QuerySNDataClassification
/*  接口9方法注释   */
//@Summary  查询异常数据停车场top10 api
//@Tags  查询异常数据停车场top10
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/abnormaldataparking [get]
func QueryAbnormalDataParking(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	//  查询异常数据停车场top10
	logrus.Print("查询异常数据停车场top10")
	code, err, totaldata := service.QueryAbnormalDataParking()
	if err != nil {
		logrus.Errorf("QueryAbnormalDataParking err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryAbnormalDataParking err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 309 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "  查询异常数据停车场top10 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryAbnormalDataParkingError, Data: types.StatusText(types.StatusQueryAbnormalDataParkingError), Message: "  查询异常数据停车场top10 失败"})
	}
}

/*  接口10方法注释   */
//@Summary 逾期数据停车场top10 api
//@Tags 查询逾期数据停车场top10
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sn/overduedata [get]
func QueryOverdueData(c *gin.Context) {
	respFailure := dto.ResponseFailure{}
	//查询逾期数据停车场top10
	code, err, totaldata := service.QueryOverdueData()
	if err != nil {
		logrus.Errorf("QueryOverdueData err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryOverdueData err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	if code == 310 {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询逾期数据停车场top10 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryOverdueDataError, Data: types.StatusText(types.StatusQueryOverdueDataError), Message: "查询逾期数据停车场top10 失败"})
	}
}
