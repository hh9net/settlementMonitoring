package controller

import (
	"fmt"
	"net/http"
	"settlementMonitoring/dto"
	"settlementMonitoring/service"
	"settlementMonitoring/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*  接口方法注释   */
//@Summary 注册api
//@Tags 注册
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /user/register [post]
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
	if code == 401 {
		c.JSON(http.StatusOK, dto.Response{Code: 4001, Data: types.StatusText(4001), Message: "重复注册"})
	}
	if code == 200 {
		c.JSON(http.StatusOK, dto.Response{Code: 0, Message: "注册成功"})
	}
}

/*  接口方法注释   */
//@Summary 登录api
//@Tags 登录
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /user/login [post]
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

	if code == 402 {
		logrus.Println("用户未注册，请先注册")
		c.JSON(http.StatusOK, dto.Response{Code: 4002, Data: types.StatusText(4002), Message: "用户未注册，请先注册"})
		return
	}

	if code == 403 {
		logrus.Println("密码错误,请重新输入")
		c.JSON(http.StatusOK, dto.Response{Code: 4003, Data: types.StatusText(4003), Message: "密码错误,请重新输入"})
		return
	}

	if code == 201 {
		logrus.Println("用户登录成功")
		c.JSON(http.StatusOK, dto.Response{Code: 0, Data: types.StatusText(0), Message: "用户登录成功"})
		return
	}
}
