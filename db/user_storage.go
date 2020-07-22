package db

import (
	"github.com/sirupsen/logrus"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
)

//用户注册 查询用户信息是否以及存在
func QueryUsermsg(username string) (error, int) {
	db := utils.GormClient.Client
	//user := new(types.BJsjkJiesjkptyhb)

	user := make([]types.BJsjkJiesjkptyhb, 0)

	if err := db.Table("b_jsjk_jiesjkptyhb").Where("F_NB_YONGHID = ?", username).Find(&user).Error; err != nil {
		// 错误处理...
		logrus.Println("Query b_jsjk_jiesjkptyhb error", err)
		return err, 0
	}
	logrus.Println(len(user))

	return nil, len(user)
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
