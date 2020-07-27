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
