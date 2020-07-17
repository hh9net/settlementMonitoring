package service

import (
	"github.com/sirupsen/logrus"
	"settlementMonitoring/dto"
)

func Login(req dto.Reqlogin) (int, error) {
	logrus.Print("登录请求参数：", req)
	//获取请求数据

	//校验数据

	//返回数据
	return 200, nil
}
