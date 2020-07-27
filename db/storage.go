package db

import (
	"github.com/pkg/errors"
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

//1、Insert b_jsjk_jiestj 新增结算统计
func InsertTabledata(lx int) error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkJiestj)
	//赋值
	Jiestj.FNbKawlh = lx //统计类型 10000 ：省外

	Jiestj.FDtKaistjsj = utils.StrTimeToNowtime()           //开始统计时间
	Jiestj.FDtTongjwcsj = utils.StrTimeTodefaultdate()      //统计完成时间
	Jiestj.FVcTongjrq = utils.StrTimeTodefaultdatetimestr() //统计日期
	if err := db.Table("b_jsjk_jiestj").Create(&Jiestj).Error; err != nil {
		// 错误处理...
		logrus.Println("Insert b_jsjk_jiestj error", err)
		return err
	}
	logrus.Println("结算统计表插入成功！", Jiestj.FDtKaistjsj)
	return nil
}

//2、 Query b_jsjk_jiestj
func QueryTabledata(lx int) (error, *types.BJsjkJiestj) {
	db := utils.GormClient.Client
	//Jiestjs := make([]types.BJsjkJiestj, 0)
	Jiestjs := new(types.BJsjkJiestj)
	//赋值
	if err := db.Table("b_jsjk_jiestj").Where("F_NB_KAWLH=?", lx).Last(&Jiestjs).Error; err != nil {
		logrus.Println("查询 结算监控统计表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	logrus.Println("查询结果:", Jiestjs)
	return nil, Jiestjs
}

//3、更新结算统计表 update b_jsjk_jiestj
func UpdateTabledata(data *types.BJsjkJiestj, lx int, id int) error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkJiestj)

	Jiestj.FNbZongje = data.FNbZongje
	Jiestj.FNbZongts = data.FNbZongts
	//Jiestj.FNbKawlh = lx //10000： 省外 3201 ：省内
	Jiestj.FDtTongjwcsj = data.FDtTongjwcsj //统计完成时间
	Jiestj.FVcTongjrq = data.FVcTongjrq
	if err := db.Table("b_jsjk_jiestj").Where("F_NB_ID=?", id).Where("F_NB_KAWLH=?", lx).Updates(&Jiestj).Error; err != nil {
		logrus.Println("更新结算统计表 error", err)
		return err
	}
	return nil
}

//业务处理数据层操作
//1查询结算数据表 获得总数与总金额
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

//2按卡网络号查询结算表数据
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

//3按卡网络号查询结算表省内数据
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

//4按卡网络号查询省外结算表数据
func QueryShengwJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Not("F_VC_KAWLH = ?", 3201).Count(&count)
	var total_money []int64 //
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where  NOT (F_VC_KAWLH =?)`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	logrus.Printf("查询省外结算交易，结算表总交易笔数%d，查询总金额为：%d", count, total_money)
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

//省外业务数据层逻辑

//4.1.2	查询数据库中省外已清分的交易 总条数、总金额
//1、查询数据库中省外已清分的交易 总条数、总金额【包含坏账的金额和条数】
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

//2、新增省外已清分统计开始记录
func ShengwClearingInsert() error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkShengwqftj)
	//赋值
	Jiestj.FDtKaistjsj = utils.StrTimeToNowtime()      //开始统计时间
	Jiestj.FDtTongjwcsj = utils.StrTimeTodefaultdate() //统计完成时间
	//Jiestj.FVcTongjrq = utils.StrTimeTodefaultdatetimestr() //统计日期
	if err := db.Table("b_jsjk_shengwqftj").Create(&Jiestj).Error; err != nil {
		// 错误处理...
		logrus.Println("Insert b_jsjk_shengwqftj error", err)
		return err
	}
	logrus.Println("结算统计表插入成功！", Jiestj.FDtKaistjsj)
	return nil
}

//3、查询最新包含坏账的已清分数据 b_jsjk_shengwqftj
func QueryShengwClearingdata() (error, *types.BJsjkShengwqftj) {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkShengwqftj)
	//赋值
	if err := db.Table("b_jsjk_shengwqftj").Last(&Jiestj).Error; err != nil {
		logrus.Println("查询 省外清分统计表最新数据时 QueryShengwClearingdata error :", err)
		return err, nil
	}
	logrus.Println("查询省外清分统计表最新数据结果:", Jiestj)
	return nil, Jiestj
}

//4、更新最新的清分统计结果 b_jsjk_shengwqftj
func UpdateShengwClearingdata(data *types.BJsjkShengwqftj, id int) error {
	db := utils.GormClient.Client
	qingftj := new(types.BJsjkShengwqftj)

	qingftj.FNbZongje = data.FNbZongje       //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	qingftj.FNbZongts = data.FNbZongts       //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	qingftj.FDtTongjwcsj = data.FDtTongjwcsj //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	qingftj.FVcTongjrq = data.FVcTongjrq     //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
	qingftj.FNbHuaizje = data.FNbHuaizje     //`F_NB_HUAIZJE` bigint DEFAULT NULL COMMENT '坏账金额',
	qingftj.FNbHuaizts = data.FNbHuaizts     //`F_NB_HUAIZTS` bigint DEFAULT NULL COMMENT '坏账条数',

	if err := db.Table("b_jsjk_shengwqftj").Where("F_NB_ID=?", id).Updates(&qingftj).Error; err != nil {
		logrus.Println("更新最新的清分统计结果 error", err)
		return err
	}
	return nil
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

//查询结算表坏账的数据【测试】
func QueryJieSuan() {
	db := utils.GormClient.Client
	js := make([]types.BJsJiessj, 0)
	//1:已清分 F_NB_ZHENGYCLJG 争议处理结果：坏账2
	db.Table("b_js_jiessj").Where("F_NB_ZHENGYCLJG = ?", 2).Find(&js)
	logrus.Println("查询结算表已清分的坏账交易笔数:", len(js), js)
}

//4.1.3	查询省外结算数据中存在争议的数据总条数、总金额
//1、查询省外结算数据中存在争议的数据总条数、总金额
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

//2、新增省外存在争议的数据的统计开始记录
func ShengwDisputeInsert() error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkShengwjszysjtj)
	//赋值
	Jiestj.FDtKaistjsj = utils.StrTimeToNowtime()      //开始统计时间
	Jiestj.FDtTongjwcsj = utils.StrTimeTodefaultdate() //统计完成时间
	//Jiestj.FVcTongjrq = utils.StrTimeTodefaultdatetimestr() //统计日期
	if err := db.Table("b_jsjk_shengwjszysjtj").Create(&Jiestj).Error; err != nil {
		// 错误处理...
		logrus.Println("Insert b_jsjk_shengwjszysjtj error", err)
		return err
	}
	logrus.Println("新增省外存在争议的数据的统计开始记录成功！", Jiestj.FDtKaistjsj)
	return nil
}

//3、查询最新的争议的数据记录
func QueryShengwDispute() (error, *types.BJsjkShengwjszysjtj) {
	db := utils.GormClient.Client
	Jiestjs := new(types.BJsjkShengwjszysjtj)
	//赋值
	if err := db.Table("b_jsjk_shengwjszysjtj").Last(&Jiestjs).Error; err != nil {
		logrus.Println("查询最新的争议的数据记录时，QueryShengwDispute error :", err)
		return err, nil
	}
	logrus.Println("查询最新的争议的数据记录结果:", Jiestjs)
	return nil, Jiestjs
}

//4、更新最新的争议的数据记录 update b_jsjk_shengwjszysjtj
func UpdateShengwDispute(data *types.BJsjkShengwjszysjtj, id int) error {
	db := utils.GormClient.Client
	zytj := new(types.BJsjkShengwjszysjtj)

	zytj.FNbZongje = data.FNbZongje
	zytj.FNbZongts = data.FNbZongts
	zytj.FDtTongjwcsj = data.FDtTongjwcsj //统计完成时间
	zytj.FVcTongjrq = data.FVcTongjrq
	if err := db.Table("b_jsjk_shengwjszysjtj").Where("F_NB_ID=?", id).Updates(&zytj).Error; err != nil {
		logrus.Println("最新的争议的数据记录 error", err)
		return err
	}
	return nil
}

//4.1.4	查询待处理的异常数据 总条数、总金额【单点+总对总】
func QueryAbnormalData(lx int) (int, int64, error) {
	db := utils.GormClient.Client
	zdzcount := 0
	ddcount := 0
	var zdztotal_money []int64
	var ddtotal_money []int64
	if lx == 1 {
		//出口异常表
		db.Table("b_zdz_chedckyssjlycb").Count(&zdzcount)
		zdzsqlstr := `select SUM(F_NB_JINE) as total_money from b_zdz_chedckyssjlycb `
		db.Raw(zdzsqlstr).Pluck("SUM(F_NB_JINE) as total_money", &zdztotal_money)
		logrus.Printf("查询总对总异常数据表 异常的交易笔数%d，查询异常的交易总金额为：%d", zdzcount, zdztotal_money)

		return zdzcount, zdztotal_money[0], nil
	}

	if lx == 2 {
		//出口异常表
		db.Table("b_dd_chedckyssjlycb").Count(&ddcount)
		sqlstr := `select SUM(F_NB_JINE) as total_money from b_dd_chedckyssjlycb `
		db.Raw(sqlstr).Pluck("SUM(F_NB_JINE) as total_money", &ddtotal_money)

		logrus.Printf("查询单点异常数据表 异常的交易笔数%d，查询异常的交易总金额为：%d", ddcount, ddtotal_money)

		return ddcount, ddtotal_money[0], nil
	}
	return 0, 0, errors.New("查询待处理的异常数据 error")
}

//
//2、新增异常数据的统计开始记录
func AbnormalDataInsert(lx int) error {
	db := utils.GormClient.Client
	yctj := new(types.BJsjkYicsjtj)
	//赋值
	yctj.FDtKaistjsj = utils.StrTimeToNowtime()      //开始统计时间
	yctj.FDtTongjwcsj = utils.StrTimeTodefaultdate() //统计完成时间
	yctj.FNbTongjlx = lx                             //1:zdz 2:dd
	if err := db.Table("b_jsjk_yicsjtj").Create(&yctj).Error; err != nil {
		// 错误处理...
		logrus.Println("Insert b_jsjk_yicsjtj error", err)
		return err
	}
	logrus.Println("新增异常数据的数据的统计开始记录成功！", yctj.FDtKaistjsj)
	return nil
}

//3、查询最新的异常数据统计记录
func QueryAbnormaltable(lx int) (error, *types.BJsjkYicsjtj) {
	db := utils.GormClient.Client
	yctjs := new(types.BJsjkYicsjtj)
	//赋值
	if err := db.Table("b_jsjk_yicsjtj").Where("F_NB_TONGJLX=?", lx).Last(&yctjs).Error; err != nil {
		logrus.Println("查询最新的异常数据的数据记录时，QueryShengwDispute error :", err)
		return err, nil
	}
	logrus.Println("查询最新的异常数据的数据记录结果:", yctjs)
	return nil, yctjs
}

//4、更新最新的异常数据统计记录
func UpdateAbnormalData(data *types.BJsjkYicsjtj, id int) error {
	db := utils.GormClient.Client
	zytj := new(types.BJsjkYicsjtj)

	zytj.FNbZongje = data.FNbZongje
	zytj.FNbZongts = data.FNbZongts
	zytj.FDtTongjwcsj = data.FDtTongjwcsj //统计完成时间
	zytj.FVcTongjrq = data.FVcTongjrq
	if err := db.Table("b_jsjk_yicsjtj").Where("F_NB_ID=?", id).Updates(&zytj).Error; err != nil {
		logrus.Println("最新的异常数据的数据记录 error", err)
		return err
	}
	return nil
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
