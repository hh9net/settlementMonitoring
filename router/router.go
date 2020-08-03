package router

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"settlementMonitoring/controller"
	_ "settlementMonitoring/docs"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteInit(IpAddress string) {
	router := gin.New()
	router.Use(Cors()) //跨域资源共享

	url := ginSwagger.URL("http://127.0.0.1:8088/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiV1 := router.Group("/settlementMonitoring/api/v1")
	APIV1Init(apiV1)

	http.Handle("/", router)
	gin.SetMode(gin.ReleaseMode)
	logrus.Print("服务端 IpAddress：", IpAddress)
	runerr := router.Run(IpAddress)
	if runerr != nil {
		logrus.Print("Run error")
		return
	}
}
func APIV1Init(route *gin.RouterGroup) {
	AuthAPIInit(route)
}

func AuthAPIInit(route *gin.RouterGroup) {
	//用户注册
	route.POST("/user/register", controller.Register)
	//用户登录
	route.POST("/user/login", controller.Login)
	//查询省外总交易额、总笔数 ok
	route.GET("/sw/totalsettlementdata", controller.QueryTotalSettlementData)
	//查询省外已清分（含有坏账）总交易额、总笔数 ok
	route.GET("/sw/totalclarify", controller.QueryTotalClarify)
	//查询省外坏账总交易额、总笔数 ok
	route.GET("/sw/totalBaddebts", controller.QueryTotalBaddebts)
	//查询省外存在争议总交易额、总笔数 ok
	route.GET("/sw/totaldisputedata", controller.QueryShengwDisputedata)
	//查询省外存在异常总交易额、总笔数 ok
	route.GET("/sw/totalAbnormaldata", controller.QueryAbnormaldata)
	//查询黑名单总数
	route.GET("/sw/totalblacklistdata", controller.Queryblacklistdata)
	//清分、争议包监控:查询清分包、争议包的接收时间、包号【每天统计一次】
	route.GET("/sw/clearlingAndDisputePackageSettlement", controller.QueryClearlingAndDisputePackage)
	//清分核对StatisticalClearlingcheck()
	route.GET("/sw/clearlingcheck", controller.Clearlingcheck)
	//省外数据分类
	route.GET("/sw/dataclassification", controller.Dataclassification)
	//省外转结算监控
	route.GET("/sw/dataturnmonitor", controller.QueryDataTurnMonitor)

}

//以下为cors实现
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域

			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段

			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析

			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
