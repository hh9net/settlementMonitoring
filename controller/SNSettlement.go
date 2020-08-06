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
		c.JSON(http.StatusOK, dto.QueryResponse{Code: code, CodeMsg: types.StatusText(301), Data: *totaldata, Message: "查询省内结算数据表的总条数、总金额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQuerySNTotalSettlementDataError), Message: "查询省内结算数据表的总条数、总金额 失败"})
	}
}

//QuerySNSendTotalSettlementData
/*  接口1方法注释   */
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
		c.JSON(http.StatusOK, dto.QueryResponse{Code: code, CodeMsg: types.StatusText(302), Data: *totaldata, Message: "查询省内的已发送 总条数、总金额 成功"})
	}
	if code == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: code, Data: types.StatusText(types.StatusQuerySNSendTotalSettlemenError), Message: "查询省内的已发送 总条数、总金额 失败"})
	}
}
