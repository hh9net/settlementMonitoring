package db

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"

	"settlementMonitoring/types"
	"settlementMonitoring/utils"
)

//用户注册 查询用户信息是否以及存在
func QueryUsermsg(username string) (error, int) {
	db := utils.GormClient.Client
	users := make([]types.BJsjkJiesjkptyhb, 0)
	res := db.Table("b_jsjk_jiesjkptyhb").Where("F_NB_YONGHID = ?", username).First(&users)
	if res.Error != nil {
		logrus.Println(res.Error)
		return res.Error, 0
	}
	logrus.Println("users:", users)
	if res.RecordNotFound() {
		log.Print("Record not found")
		return errors.New("Record not found"), 0
	} else {
		logrus.Println("查询数据：", users)
		return nil, len(users)
	}
}
func QueryUserLoginmsg(username string) (error, *types.BSysYongh) {
	db := utils.GormClient.Client
	user := new(types.BSysYongh)
	if err := db.Table("b_sys_yongh").Where("F_VC_ZHANGH = ?", username).First(user).Error; err != nil {
		logrus.Println("查询用户登录信息失败！")
		return err, nil
	}
	logrus.Println("查询用户登录信息 ok:", user.FVcMingc, user.FVcZhangh)

	return nil, user

}

//用户注册 插入数据库
func UserInsert(data *types.BJsjkJiesjkptyhb) error {
	db := utils.GormClient.Client
	user := new(types.BJsjkJiesjkptyhb)
	//赋值
	user.FNbYonghid = data.FNbYonghid //   '用户id',//手机号 或者任意6位数字
	user.FVcYonghmm = data.FVcYonghmm //   '用户密码',
	user.FVcShoujh = data.FVcShoujh   //   '手机号',
	user.FVcYoux = data.FVcYoux       //   '邮箱',
	user.FVcYonghnc = data.FVcYonghnc // '用户昵称',

	if err := db.Table("b_jsjk_jiesjkptyhb").Create(&user).Error; err != nil {
		// 错误处理...
		logrus.Println("Insert b_jsjk_jiesjkptyhb error", err)
		return err
	}
	logrus.Println("用户表插入成功！")
	return nil
}
