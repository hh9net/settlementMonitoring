package db

import (
	"github.com/sirupsen/logrus"
	"settlementMonitoring/config"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"time"
)

//结算监控平台数据层：数据的增删改查
func Newdb() {
	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	DBInit() //初始化数据库
}

//1、查询表是否存在
func QueryTable(tablename string) {
	db := utils.GormClient.Client
	is := db.HasTable(tablename)

	if is == false {
		logrus.Println("不存在", tablename)
		return
	}
	logrus.Println("表存在：", tablename, is)
}

//2、Insert b_jsjk_jiestj
func InsertTabledata() error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkJiestj)
	//赋值
	Jiestj.FDtKaistjsj = utils.StrTimeToNowtime()      //开始统计时间
	Jiestj.FDtTongjwcsj = utils.StrTimeTodefaultdate() //统计完成时间
	Jiestj.FDtTongjrq = utils.StrTimeTodefaultdate()   //统计日期
	if err := db.Table("b_jsjk_jiestj").Create(&Jiestj).Error; err != nil {
		// 错误处理...
		logrus.Println("Insert b_jsjk_jiestj error", err)
		return err
	}
	logrus.Println("结算统计表插入成功！", Jiestj.FDtTongjrq)
	return nil
}

//3、 Query b_jsjk_jiestj
func QueryTabledata() error {
	db := utils.GormClient.Client
	Jiestj := make([]types.BJsjkJiestj, 0)
	//赋值
	db.Table("b_jsjk_jiestj").Where("F_DT_TONGJRQ=?", "2020-01-02 00:00:00").Find(&Jiestj)

	for _, sj := range Jiestj {
		logrus.Println("查询结果", sj.FDtTongjrq.Format("2006-01-02 15:04:05"), sj.FDtKaistjsj)
	}
	return nil
}

//4、update b_jsjk_jiestj
func UpdateTabledata() error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkJiestj)

	logrus.Println(Jiestj.FNbId)
	Jiestj.FNbZongje = 1212000
	Jiestj.FNbZongts = 12
	Jiestj.FNbKawlh = 3202
	if err := db.Table("b_jsjk_jiestj").Where("F_NB_ID=?", 16).Updates(&Jiestj).Error; err != nil {
		logrus.Error(err)
	}
	return nil
}

//业务处理数据层操作
//查询结算数据表 获得总数与总金额
func QueryJieSuanTable() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Count(&count)
	logrus.Println("查询结算表总交易笔数", count)

	var total_money []int64

	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj`
	db.Raw(sqlstr).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询结算表总交易笔数为%d， 查询总金额为：%d", count, total_money[0])
	return count, total_money[0]
}

//按卡网络号查询结算表数据
func QueryKawlhJieSuan(kawlh int) (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", kawlh).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, kawlh).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", kawlh, count, total_money)
	return count, total_money[0]
}

//按卡网络号查询结算表数据
func QueryShengnJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4.2.2	查询省内的已发送 总条数、总金额【不做】
func QueryShengnSendedJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201). /*Where().*/ Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4.2.3	查询省内已请款的数据总条数、总金额【不做】
func QueryShengnPleaseedJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201). /*Where().*/ Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4.2.4	查询坏账（拒付）数据 总条数、总金额【不做】
func QueryShengnBadDebtsJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201). /*Where().*/ Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4.2.5	省内实时数据

//4.1.2	查询数据库中省外已清分的交易 总条数、总金额【包含坏账的】
func QueryShengwClearingJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//已清分
	db.Table("b_js_jiessj").Where("F_NB_QINGFJG = ?", 1).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_NB_QINGFJG = ?`
	db.Raw(sqlstr, 1).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询结算表总清分的交易笔数%d，查询已清分总金额为：%d", count, total_money)
	return count, total_money[0]
}

// 查询 已清分的坏账 Bad debts
func QueryShengwBadDebtsJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//1:已清分 F_NB_ZHENGYCLJG 争议处理结果：坏账2
	db.Table("b_js_jiessj").Where("F_NB_QINGFJG = ?", 1).Where("F_NB_ZHENGYCLJG = ?", 2).Count(&count)

	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_NB_QINGFJG = ? and F_NB_ZHENGYCLJG = ?`
	db.Raw(sqlstr, 1, 2).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询结算表已清分的坏账交易笔数%d，查询已清分的坏账总金额为：%d", count, total_money)
	return count, total_money[0]
}

//4.1.3	查询省外结算数据中存在争议的数据总条数、总金额
func QueryDisputeJieSuanData() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//2：争议数据 0：争议数据未处理
	db.Table("b_js_jiessj").Where("F_NB_JIZJG  = ?", 2).Where("F_NB_ZHENGYCLJG = ?", 0).Count(&count)

	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_NB_JIZJG = ?  and F_NB_ZHENGYCLJG = ?`
	//2：争议数据 0：争议数据未处理
	db.Raw(sqlstr, 2, 0).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询结算表 待处理存在争议的数据总笔数:%d笔，查询待处理存在争议的数据总金额为：%d分", count, total_money)
	return count, total_money[0]
}

//4.1.4	查询待处理的异常数据 总条数、总金额【单点+总对总】
func QueryAbnormalData() (int, int64) {
	db := utils.GormClient.Client
	ddcount := 0
	//出口异常表
	db.Table("b_dd_chedckyssjlycb").Count(&ddcount)
	var ddtotal_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_dd_chedckyssjlycb `
	db.Raw(sqlstr).Pluck("SUM(F_NB_JINE) as total_money", &ddtotal_money)

	logrus.Printf("查询单点异常数据表 异常的交易笔数%d，查询异常的交易总金额为：%d", ddcount, ddtotal_money)
	zdzcount := 0
	//出口异常表
	db.Table("b_zdz_chedckyssjlycb").Count(&zdzcount)
	var zdztotal_money []int64
	zdzsqlstr := `select SUM(F_NB_JINE) as total_money from b_zdz_chedckyssjlycb `
	db.Raw(zdzsqlstr).Pluck("SUM(F_NB_JINE) as total_money", &zdztotal_money)

	logrus.Printf("查询总对总异常数据表 异常的交易笔数%d，查询异常的交易总金额为：%d", zdzcount, zdztotal_money)
	return ddcount + zdzcount, ddtotal_money[0] + zdztotal_money[0]
}

//4.1.5	数据包实时状态监控
func PacketMonitoring() {

}

//今日打包数量
//打包金额
//已发送原始交易消息包数量
//已发送原始交易消息包金额
//记账包数量
//记账包金额
//原始交易消息应答包数量
//func
