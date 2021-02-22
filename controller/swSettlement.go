package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"

	"net/http"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/service"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"strconv"
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
		log.Println("QuerTotalSettlementData  err: %v", err)
		respFailure.Code = types.StatusQueryTotalSettlementDataError
		respFailure.Message = fmt.Sprintf("查询结算总金额、总笔数 失败: %v", err)
	}
	if code == types.StatusQuerySWTotalSettlementDataSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询结算总金额、总笔数 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryTotalSettlementDataError, Data: types.StatusText(types.StatusQueryTotalSettlementDataError), Message: fmt.Sprintf("查询结算总金额、总笔数 失败: %v", err)})
		return
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
		log.Println("QuerTotalClarify  err: %v", err)
		respFailure.Code = types.StatusQueryTotalSettlementDataError
		respFailure.Message = fmt.Sprintf("QuerTotalClarify err: %v", err)
		//c.JSON(types.StatusQueryTotalBaddebtsError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询已清分总金额、总笔数 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryTotalSettlementDataError, Data: types.StatusText(types.StatusQueryTotalSettlementDataError), Message: "查询已清分总金额、总笔数 失败"})
		return
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
		log.Println("QueryTotalBaddebts  err: %v", err)
		respFailure.Code = types.StatusQueryTotalBaddebtsError
		respFailure.Message = fmt.Sprintf("QueryTotalBaddebts err: %v", err)
		//c.JSON(types.StatusQueryTotalBaddebtsError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询坏账总金额、总笔数 成功"})
		return
	}
	//返回失败
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryTotalBaddebtsError, Data: types.StatusText(types.StatusQueryTotalBaddebtsError), Message: "查询坏账总金额、总笔数 失败"})
		return
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
		log.Println("QueryDisputedata  err: %v", err)
		respFailure.Code = types.StatusQueryShengwDisputedataError
		respFailure.Message = fmt.Sprintf("QueryDisputedata err: %v", err)
		//c.JSON(types.StatusQueryShengwDisputedataError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询存在争议总金额、总笔数 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryShengwDisputedataError, Data: types.StatusText(types.StatusQueryShengwDisputedataError), Message: "查询存在争议总金额、总笔数 失败"})
		return
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
		log.Println("QueryAbnormaldata  err: %v", err)
		respFailure.Code = types.StatusQueryAbnormaldataError
		respFailure.Message = fmt.Sprintf("QueryAbnormaldata err: %v", err)
		//c.JSON(types.StatusQueryAbnormaldataError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询异常数据总金额、总笔数 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryAbnormaldataError, Data: types.StatusText(types.StatusQueryAbnormaldataError), Message: "查询异常数据总金额、总笔数 失败"})
		return
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
		log.Println("Queryblacklistdata  err: %v", err)
		respFailure.Code = types.StatusQueryblacklistdataError
		respFailure.Message = fmt.Sprintf("Queryblacklistdata err: %v", err)
		//c.JSON(types.StatusQueryblacklistdataError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询黑名单总数、较2小时前变化值 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryblacklistdataError, Data: types.StatusText(types.StatusQueryblacklistdataError), Message: "查询黑名单总数、较2小时前变化值 失败"})
		return
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
		log.Println("QueryClearlingAndDisputePackagedata  err: %v", err)
		respFailure.Code = types.StatusQueryClearlingAndDisputePkgError
		respFailure.Message = fmt.Sprintf("QueryClearlingAndDisputePackagedata err: %v", err)
		//c.JSON(types.StatusQueryClearlingAndDisputePkgError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询清分包、争议包的接收时间、包号 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryClearlingAndDisputePkgError, Data: types.StatusText(types.StatusQueryClearlingAndDisputePkgError), Message: "查询清分包、争议包的接收时间、包号 失败"})
		return
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
		log.Println("StatisticalClearlingcheck  err: %v", err)
		respFailure.Code = types.StatusQueryClearlingcheckOneDataError
		respFailure.Message = fmt.Sprintf("StatisticalClearlingcheck err: %v", err)
		//c.JSON(types.StatusQueryClearlingcheckOneDataError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询清分核对结果 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryClearlingcheckOneDataError, Data: types.StatusText(types.StatusQueryClearlingcheckOneDataError), Message: "查询清分核对结果 失败"})
		return
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
		log.Println("QueryDataclassification  err: %v", err)
		respFailure.Code = types.StatusQueryDataclassificationError
		respFailure.Message = fmt.Sprintf("QueryDataclassification err: %v", err)
		//c.JSON(types.StatusQueryDataclassificationError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省外数据分类 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataclassificationError, Data: types.StatusText(types.StatusQueryDataclassificationError), Message: "查询省外数据分类 失败"})
		return
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
		log.Println("QueryDataTurnMonitordata err: %v", err)
		respFailure.Code = types.StatusQueryDataTurnMonitorError
		respFailure.Message = fmt.Sprintf("QueryDataTurnMonitordata err: %v", err)
		//c.JSON(types.StatusQueryDataTurnMonitorError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省外转结算 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataTurnMonitorError, Data: types.StatusText(types.StatusQueryDataTurnMonitorError), Message: "查询省外转结算 失败"})
		return
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
		log.Println("QuerySettlementTrend err: %v", err)
		respFailure.Code = types.StatusQuerySettlementTrendError
		respFailure.Message = fmt.Sprintf("QuerySettlementTrend err: %v", err)
		//c.JSON(types.StatusQuerySettlementTrendError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省外结算趋势 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQuerySettlementTrendError, Data: types.StatusText(types.StatusQuerySettlementTrendError), Message: "查询省外结算趋势 失败"})
		return
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
		log.Println("QueryPacketMonitoring err: %v", err)
		respFailure.Code = types.StatusQueryPacketMonitoringError
		respFailure.Message = fmt.Sprintf("QueryPacketMonitoring err: %v", err)
		//c.JSON(types.StatusQueryPacketMonitoringError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询省外数据包监控 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryPacketMonitoringError, Data: types.StatusText(types.StatusQueryPacketMonitoringError), Message: "查询省外数据包监控 失败"})
		return
	}
}

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
		log.Println("Clarifydifference err: %v", err)
		respFailure.Code = types.StatusQueryClarifydifferenceError
		respFailure.Message = fmt.Sprintf("QueryClarifydifference err: %v", err)
		//c.JSON(types.StatusQueryClarifydifferenceError, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询最近15天清分包数据差额 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryClarifydifferenceError, Data: types.StatusText(types.StatusQueryClarifydifferenceError), Message: "查询最近15天清分包数据差额 失败"})
		return
	}
}

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
		log.Println("ReqQueryClarify json unmarshal err: %v", err)
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err)
		//c.JSON(types.StatusQueryClarifyError, respFailure)
		//return
	}
	//查询清分核对
	code, err, totaldata := service.ClarifyQuery(req)
	if err != nil {
		log.Println("QueryClarify err: %v", err)
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryClarify err: %v", err)
		log.Println("+++++++++++++++service.ClarifyQuery(req) err: %v", err)
		//c.JSON(http.StatusOK, respFailure)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryClarifyError, Data: types.StatusText(types.StatusQueryClarifyError), Message: "查询清分核对 失败"})
		return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "查询清分核对 成功"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryClarifyError, Data: types.StatusText(types.StatusQueryClarifyError), Message: "查询清分核对 失败"})
		return
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
		log.Println("Clarifyconfirmerr: %v", err)
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Clarifyconfirm err: %v", err)
		//c.JSON(http.StatusOK, respFailure)
		//return
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QuerResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Data: *totaldata, Message: "清分确认成功 【待实现】"})
		return
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryPacketMonitoringError, Data: types.StatusText(types.StatusQueryPacketMonitoringError), Message: "清分确认 失败 【待实现】"})
		return
	}
}

/*  接口14方法注释  */
//@Summary 导出清分核对记录为excel api
//@Tags 导出清分核对记录为excel
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/exportexcel [post]
func ExportExcel(c *gin.Context) {
	req := dto.ReqClarifyExportExcel{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		log.Println("ReqQueryClarify json unmarshal err: %v", err)
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err)
		//c.JSON(http.StatusOK, respFailure)
		//return
	}
	//导出清分核对记录为excel
	code, err, totaldata, fileName := service.ExportExcel(req)
	if err != nil {
		log.Println("QueryClarify err: %v", err)
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("QueryClarify err: %v", err)
		c.JSON(types.StatusExportExcelError, respFailure)
		return
	}
	log.Println("fileName:", fileName)
	if code == types.StatusSuccessfully {
		//c.JSON(http.StatusOK, dto.QuerResponse{Code: 0, CodeMsg: types.StatusText(0), Data: totaldata, Message: "导出清分包核对记录表 成功"})
		c.Header("Content-Disposition", "attachment;filename="+fileName)
		n, _ := c.Writer.Write(totaldata)
		log.Println(n)
		utils.DelFile("./" + fileName)
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusExportExcelError, Data: types.StatusText(types.StatusExportExcelError), Message: "导出清分包核对记录表 失败"})
		return
	}
}

/*  接口14方法注释  */
//@Summary 导出清分核对记录为excel api
//@Tags 导出清分核对记录为excel
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /sw/exportexcel [post]
func SetRedis(c *gin.Context) {

	conn := utils.Pool.Get() //初始化redis
	//redis set新值
	s := strconv.Itoa(1) + "|" + strconv.Itoa(1) + "|" + strconv.Itoa(1) + "|" + strconv.Itoa(1) + "|" + strconv.Itoa(1) + "|" + strconv.Itoa(1)
	rseterr := utils.RedisSet(&conn, "snshishishuju", s)
	if rseterr != nil {
		log.Print("set redis snshishishuju 零值error", rseterr)
	}
	//redis set新值
	rhseterr := utils.RedisSet(&conn, "swjiesuantotal", strconv.Itoa(1)+"|"+strconv.Itoa(1))
	if rhseterr != nil {
		log.Print("set redis swjiesuantotal 零值error", rhseterr)
	}
	//redis set新值
	rsnseterr := utils.RedisSet(&conn, "snjiesuantotal", strconv.Itoa(1)+"|"+strconv.Itoa(1))
	if rsnseterr != nil {
		log.Print("set redis snjiesuantotal 零值error", rsnseterr)
	}

	m := make(map[string]string, 0)
	m["2020-09-00"] = "999990" + "|" + "2020-09-00 11:11:11"
	hmseterr := utils.RedisHMSet(&conn, "clear", m)
	if hmseterr != nil {
		log.Print("set redis clear 零值error", rsnseterr)
	}

	m["2020-09-00"] = "999990" + "|" + "2020-09-00 11:11:11"
	//2、把数据存储于redis  接收时间、包号
	chmseterr := utils.RedisHMSet(&conn, "disput", m)
	if chmseterr != nil {
		log.Print("set redis  disput 零值error", chmseterr)
	}
	defer func() {
		_ = conn.Close()
	}()
	log.Println("set redis 成功 ", s, strconv.Itoa(1)+"|"+strconv.Itoa(1), strconv.Itoa(1)+"|"+strconv.Itoa(1))
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "set redis ok", Message: "set redis   零值 ok "})
}

//清分核对校准
func Clearcalibration(c *gin.Context) {
	req := dto.ReqQuery{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		log.Println("ReqQueryClarify json unmarshal err: %v", err)
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err)
		c.JSON(types.StatusQueryClarifyError, respFailure)
		return
	}
	log.Println("清分核对校准的清分包接收日期：", req.BeginTime)

	//1、获取今天的清分包数据
	qerr, clears := db.QueryClearlingdata(req.BeginTime)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "清分核对校准,获取昨日的清分包数据error", Message: "清分核对校准,获取昨日的清分包数据error "})
		return
	}
	if clears == nil {
		log.Println("+++++++++++++++++++++++++++昨日没有清分包【1.5】++++++++++++++++++++++")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "清分核对校准,昨日没有清分包", Message: "清分核对校准,昨日没有清分包"})
		return
	}
	for _, clear := range *clears {
		qcrerr := db.QueryCheckResult(clear.FNbXiaoxxh)
		if qcrerr != nil {
			if fmt.Sprint(qcrerr) == "查询清分核对结果成功,不能重复插入" {
				log.Println("查询清分核对结果成功,不能重复插入")
				c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "查询清分核对结果成功,不能重复插入", Message: "查询清分核对结果成功,不能重复插入"})
				return
			}
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "清分核对校准error", Message: "清分核对校准error "})
			return
		}
		//2、统计昨日记账包总金额
		s1 := strings.Split(clear.FVcQingfmbr, "T")
		log.Println("++++统计记账包")
		keepAccount, keepAccountCount := db.StatisticalkeepAccount(s1[0])
		log.Println("++++统计记账包完成")

		//统计存在争议数据
		disputerr, Disput, zyfgsl := db.DisputedDataCanClearling(clear.FNbXiaoxxh)
		if disputerr != nil {
			//return disputerr
		}
		//统计退费数据
		SWRefund := db.StatisticalRefund(s1[0])
		f := strconv.FormatFloat(float64(SWRefund.Total), 'f', 2, 64)
		fs := strings.Split(f, ".")

		i, _ := strconv.Atoi(fs[0] + fs[1])

		var zhengyclje int64
		if Disput == nil {
			zhengyclje = 0
		} else {
			zhengyclje = Disput.FNbQuerxyjzdzje
		}
		log.Println("今日核对清分结果的总金额：", keepAccount+zhengyclje)
		log.Println("清分包清分总金额：", clear.FNbQingfzje)
		log.Println("清分包退费总金额：", int64(i))

		var is int
		if (clear.FNbQingfzje == keepAccount+zhengyclje-int64(i)) && (clear.FNbQingfsl == keepAccountCount+zyfgsl+SWRefund.Count) {
			is = 1
			log.Println("清分核对正确++++++++++++++")
		} else {
			is = 2
			log.Println("清分核对不正确++++++++++++++++++++++++++++++++")
		}

		//把清分核对结果存数据库
		data := new(types.BJsjkQingfhd)
		//赋值
		data.FNbQingfqrzt = 0                          // `F_NB_QINGFQRZT` int DEFAULT NULL COMMENT '清分确认状态',
		data.FNbQingfts = clear.FNbQingfsl             //`F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
		data.FNbTongjqfts = keepAccountCount + zyfgsl  //`F_NB_TONGJQFTS` int DEFAULT NULL COMMENT '统计清分条数',
		data.FDtQingfbjssj = clear.FDtJiessj           //`F_VC_QINGFBJSSJ` int DEFAULT NULL COMMENT '清分包接收时间',
		data.FNbQingfbxh = clear.FNbXiaoxxh            //   `F_NB_QINGFBXH` bigint DEFAULT NULL COMMENT '清分包序号',
		data.FNbQingfje = clear.FNbQingfzje            //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
		data.FNbTongjqfje = keepAccount + (zhengyclje) //   `F_NB_TONGJQFJE` bigint DEFAULT NULL COMMENT '统计清分金额',
		data.FNbHedjg = is                             //   `F_NB_HEDJG` int DEFAULT NULL COMMENT '核对结果 是否一致,1:一致，2:不一致',
		data.FNbTuifje = int64(i)                      //退费总金额  分
		data.FNbTuifts = SWRefund.Count                //退费总条数

		s := strings.Split(clear.FVcQingfmbr, "T")
		data.FVcTongjrq = s[0] //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',【清分包的清分目标日】

		cherr := db.CheckResultInsert(data)
		if cherr != nil {
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "清分核对校准 插入 error", Message: "清分核对校准 插入 error "})
			return
		}
		log.Println("清分金额核对完成++++++++++++++++++++【1.5】+++++++++++++")
	}

	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "清分核对校准 ok", Message: "清分核对校准 ok "})

}

//清分包、争议包校准
func ClearlingAndDisputePackagecalibration(c *gin.Context) {
	conn := utils.Pool.Get() //初始化redis

	defer func() {
		_ = conn.Close()
	}()
	req := dto.ReqQuery{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		log.Println("ReqQueryClarify json unmarshal err: %v", err)
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err)
		c.JSON(types.StatusQueryClarifyError, respFailure)
		return
	}

	//1、获取清分包、争议包数据
	Yesterday := req.BeginTime
	qcerr, clears := db.QueryClearlingdata(Yesterday)
	if qcerr != nil {
		log.Println("获取清分包数据 错误 qcerr:", qcerr)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "取清分包数据 错误 error", Message: "取清分包数据 错误 error "})
		return
	}

	if clears == nil {
		Clear := types.ClearlingAndDispute{
			DataType:  "clear",
			PackageNo: "",
			DateTime:  "",
		}
		m := make(map[string]string, 0)
		// key:日期    value:"包号"｜"时间"
		m[Yesterday] = Clear.PackageNo + "|" + Clear.DateTime
		//2、把数据存储于redis  接收时间、包号
		hmseterr := utils.RedisHMSet(&conn, Clear.DataType, m)
		if hmseterr != nil {
			log.Println("utils.RedisHMSet 错误")
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "utils.RedisHMSet 错误", Message: "utils.RedisHMSet 错误"})
			return
		}
		log.Println("获取清分包-【RedisHSet】 v:=clear 成功 【++++++++++++[1.4]++++++++++++++++++】")
	} else {
		for _, clear := range *clears {
			Clear := types.ClearlingAndDispute{
				DataType:  "clear",
				PackageNo: strconv.Itoa(int(clear.FNbXiaoxxh)),
				DateTime:  clear.FDtChulsj.Format("2006-01-02 15:04:05"),
			}
			m1 := make(map[string]string, 0)
			// key:日期    value:"包号"｜"时间"
			sj := strings.Split(clear.FVcQingfmbr, "T")
			m1[sj[0]] = Clear.PackageNo + "|" + Clear.DateTime
			//2、把数据存储于redis  接收时间、包号
			hmseterr := utils.RedisHMSet(&conn, Clear.DataType, m1)
			if hmseterr != nil {
				log.Println("utils.RedisHMSet 错误")
				c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "utils.RedisHMSet 错误 error", Message: "utils.RedisHMSet 错误 error "})
				return
			}
			log.Println("获取清分包-【RedisHSet】 v:=clear 成功 【++++++++++++[1.4]++++++++++++++++++】")

		}
	}

	//1查询争议处理数据
	qderr, dispute := db.QueryDisputedata(Yesterday)
	if qderr != nil {
		log.Println("获取争议包数据 错误")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "获取争议包数据 错误 error", Message: "获取争议包数据 错误 "})
		return
	}
	var Disput types.ClearlingAndDispute
	if dispute == nil {
		Disput = types.ClearlingAndDispute{
			DataType:  "disput",
			PackageNo: "",
			DateTime:  "",
		}
	} else {
		Disput = types.ClearlingAndDispute{
			DataType:  "disput",
			PackageNo: strconv.Itoa(int(dispute.FNbXiaoxxh)),
			DateTime:  utils.DateTimeFormat(dispute.FDtZhengyclsj),
		}
	}
	m2 := make(map[string]string, 0)
	//2、把数据存储于redis  接收时间、包号
	m2[Yesterday] = Disput.PackageNo + "|" + Disput.DateTime

	dishmseterr := utils.RedisHMSet(&conn, Disput.DataType, m2)
	if dishmseterr != nil {
		log.Println(" utils.RedisHMSet 错误")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "utils.RedisHMSet 错误 error", Message: "utils.RedisHMSet 错误 "})
		return
	}
	log.Println("获取争议包-【RedisHSet】 v:=disput 成功 【++++++++++++[1.4]++++++++++++++++++】")
	log.Println("  【++++++++++++【1.4 是 ok的】++++++++++++++++++】")
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "清分包、争议包校准 ok", Message: "清分包、争议包校准 ok "})

}

func Settlementtrendcalibration(c *gin.Context) {

	//省外结算趋势
	qserr := db.SettlementTrendbyDay()
	if qserr != nil {
		log.Println("+++++++++++++++++++++【1.6error】+++++++++++++++++=查询省外结算趋势定时任务 error:", qserr)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "省内、省外结算趋势校准 error", Message: "省内、省外结算趋势校准error "})
		return
	}

	//省内结算趋势
	qsnqserr := db.QueryShengNSettlementTrenddata()
	if qsnqserr != nil {
		log.Println("+++++++++++++++++++++【1.8error】+++++++++++++++++=查询省内结算分类 定时任务 error:", qsnqserr)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "省内、省外结算趋势校准 error", Message: "省内、省外结算趋势校准error "})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "省内、省外结算趋势校准 ok", Message: "省内、省外结算趋势校准ok "})
}

//
////关于前一天以及之前的未记账处理的原始消息包查询
//func GetUnkeepaccountMsg(c *gin.Context) {
//
//	qsnqserr := db.QueryUnkeepaccountMsg()
//	if qsnqserr != nil {
//		log.Println("+++++++++++++++++++++【1.8error】+++++++++++++++++=查询省内结算分类 定时任务 error:", qsnqserr)
//		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "省内、省外结算趋势校准 error", Message: "省内、省外结算趋势校准error "})
//		return
//	}
//
//	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "省内、省外结算趋势校准 ok", Message: "省内、省外结算趋势校准ok "})
//
//}

//省外结算趋势更新
func SWSettlementTrendUpdate(c *gin.Context) {
	//更新省外结算趋势数据
	code, err := service.SWSettlementTrendUpdate()
	if err != nil {
		log.Println("SWSettlementTrendUpdate err: %v", err)
		//respFailure.Code = code
		//respFailure.Message = fmt.Sprintf("SettlementTrendUpdate err: %v", err)
	}
	if code == types.StatusSuccessfully {
		c.JSON(http.StatusOK, dto.QueryResponse{Code: types.StatusSuccessfully, CodeMsg: types.StatusText(types.StatusSuccessfully), Message: "更新省外前30日省外结算趋势概览成功"})
	}
	if code == types.Statuszero {
		c.JSON(http.StatusOK, dto.Response{Message: "更新省外前30日省外结算趋势概览失败"})
	}
}
