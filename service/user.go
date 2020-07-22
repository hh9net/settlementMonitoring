package service

import (
	"github.com/sirupsen/logrus"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/types"
)

//用户注册
func Register(req dto.ReqRegister) (int, error) {
	logrus.Print("用户注册请求参数：", req)
	//获取请求数据

	err, jg := db.QueryUsermsg(req.UserName)
	if err != nil {
		return 400, err
	}
	//校验数据
	if jg > 0 {
		logrus.Println("重复注册")
		return 401, nil
	}
	data := new(types.BJsjkJiesjkptyhb)

	data.FNbYonghid = req.UserName //   '用户id',//手机号 或者任意6位数字
	data.FVcYonghmm = req.Password //   '用户密码',
	data.FVcShoujh = req.Num       //   '手机号',
	data.FVcYoux = req.Email       //   '邮箱',
	data.FVcYonghnc = req.Name     // '用户昵称',
	inerr := db.UserInsert(data)
	if inerr != nil {
		logrus.Println("db.UserInsert error!")
		return 402, inerr
	}
	logrus.Println("注册成功")
	//返回数据
	return 200, nil
}

func Login(req dto.Reqlogin) (int, error) {
	logrus.Print("登录请求参数：", req)
	//获取请求数据

	//校验数据

	//返回数据
	return 200, nil
}
