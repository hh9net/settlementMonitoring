package controller

import (
	"fmt"
	"net/http"
	"settlementMonitoring/dto"
	"settlementMonitoring/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*  接口方法注释   */
//@Summary 登录api
//@Tags 登录
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /login [post]
func Login(c *gin.Context) {
	req := dto.Reqlogin{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("Login json unmarshal err: %v", err.Error())
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	//登录处理
	code, err := service.Login(req)
	if err != nil {
		logrus.Errorf("Login  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Login err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Code: code, Message: "success"})
}

/*  接口方法注释   */
//@Summary 注册api
//@Tags 注册
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /register [post]
func Register(c *gin.Context) {
	req := dto.ReqRegister{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("ReqRegister json unmarshal err: %v", err.Error())
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	//注册处理
	code, err := service.Register(req)
	if err != nil {
		logrus.Errorf("Register  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Register err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Code: code, Message: "success"})
}

/*  接口方法注释   */
//@Summary 查询api
//@Tags 查询
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /yiqingfen [post]
func Queryyiqingfen(c *gin.Context) {
	req := dto.ReqRegister{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("ReqRegister json unmarshal err: %v", err.Error())
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	//注册处理
	code, err := service.Register(req)
	if err != nil {
		logrus.Errorf("Register  err: %v", err.Error())
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Register err: %v", err.Error())
		c.JSON(http.StatusOK, respFailure)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Code: code, Message: "success"})
}
