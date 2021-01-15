package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
	"settlementMonitoring/config"
	"settlementMonitoring/dto"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"strconv"
	"strings"
	"time"
)

//结算监控平台数据层：数据的增删改查
func Newdb() {
	conf := config.ConfigInit() //初始化配置
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)
	mstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + conf.Mdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	DBInit(mstr) //初始化数据库
}

//1、查询表是否存在
func QueryTable(tablename string) {
	db := utils.GormClient.Client
	is := db.HasTable(tablename)

	if is == false {
		log.Println("不存在", tablename)
		return
	}
	log.Println("表存在：", tablename, is)
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
		log.Println("Insert b_jsjk_jiestj error", err)
		return err
	}
	log.Println("省外-结算统计表插入成功！", "开始统计时间:=", Jiestj.FDtKaistjsj)
	return nil
}

//2、 Query b_jsjk_jiestj
func QueryTabledata(lx int) (error, *types.BJsjkJiestj) {
	db := utils.GormClient.Client
	//Jiestjs := make([]types.BJsjkJiestj, 0)
	Jiestjs := new(types.BJsjkJiestj)
	//赋值
	if err := db.Table("b_jsjk_jiestj").Where("F_NB_KAWLH=?", lx).Last(&Jiestjs).Error; err != nil {
		log.Println("查询 结算监控统计表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	log.Println("查询结算监控统计表最新数据结果:", Jiestjs)
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
		log.Println("更新结算统计表 error", err)
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
	log.Println("查询结算表总交易笔数", count)

	var total_money []int64

	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj`
	db.Raw(sqlstr).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询结算表总交易笔数为%d， 查询总金额为：%d", count, total_money[0])
	return count, total_money[0]
}

//1.1 按照停车场id 查询总金额、总条数
func QueryTingccJieSuandata() *[]types.Result {
	db := utils.GormClient.Client
	var result []types.Result

	sqlstr4 := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_NB_DABZT  not in (4)  GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr4).Scan(&result)
	log.Println("按照停车场id 查询总金额、总条数 result:", result)
	return &result
}

//1.2 新增停车场id 查询总金额、总条数记录
func InsertTingjiesuan() error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkTingccjssjtj)
	//赋值
	Jiestj.FDtKaistjsj = utils.StrTimeToNowtime()           //开始统计时间
	Jiestj.FDtTongjwcsj = utils.StrTimeTodefaultdate()      //统计完成时间
	Jiestj.FVcTongjrq = utils.StrTimeTodefaultdatetimestr() //统计日期
	if err := db.Table("b_jsjk_tingccjssjtj").Create(&Jiestj).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_jsjk_tingccjssjtj error", err)
		return err
	}
	log.Println("停车场结算数据统计表插入成功！")
	return nil
}

//1.3 查询 停车场结算数据统计表最新数据
func QueryTingjiesuan() (error, *types.BJsjkTingccjssjtj) {
	db := utils.GormClient.Client
	Jiestjs := new(types.BJsjkTingccjssjtj)
	//赋值
	if err := db.Table("b_jsjk_tingccjssjtj").Last(&Jiestjs).Error; err != nil {
		log.Println("查询 停车场结算数据统计表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	log.Println("查询停车场结算数据统计表结果:", Jiestjs)
	return nil, Jiestjs
}

//1.4 更新停车场结算数据统计表最新数据
func UpdateTingjiesuan(data *types.BJsjkTingccjssjtj, parkingid string, id int) error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkTingccjssjtj)

	Jiestj.FNbZongje = data.FNbZongje       //
	Jiestj.FNbZongts = data.FNbZongts       //
	Jiestj.FVcTingccid = parkingid          //
	Jiestj.FDtTongjwcsj = data.FDtTongjwcsj //统计完成时间
	Jiestj.FVcTongjrq = data.FVcTongjrq
	if err := db.Table("b_jsjk_tingccjssjtj").Where("F_NB_ID=?", id).Updates(&Jiestj).Error; err != nil {
		log.Println("更新结算统计表 error", err)
		return err
	}
	return nil
}

//1.5  用id 查询 停车场结算数据统计表最新数据
func QueryTingjiesuanById(id int) (error, *types.BJsjkTingccjssjtj) {
	db := utils.GormClient.Client
	Jiestjs := new(types.BJsjkTingccjssjtj)
	//赋值
	if err := db.Table("b_jsjk_tingccjssjtj").Where("F_NB_ID=?", id).Last(&Jiestjs).Error; err != nil {
		log.Println("查询 停车场结算数据统计表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	log.Println("查询停车场结算数据统计表结果:", Jiestjs)
	return nil, Jiestjs
}

//2按卡网络号查询结算表数据
func QueryKawlhJieSuan(kawlh int) (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", kawlh).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, kawlh).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", kawlh, count, total_money)
	return count, total_money[0]
}

//3 按卡网络号查询结算表省内数据
func QueryShengnJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//不统计历史数据
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Not("F_NB_DABZT = ?", 4).Not("F_NB_DABZT = ?", 5).Count(&count)
	var total_money []int64

	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ? and not F_NB_DABZT =? and not F_NB_DABZT =?`
	db.Raw(sqlstr, 3201, 4, 5).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4按卡网络号查询省外结算表数据
func QueryShengwJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0

	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Not("F_VC_KAWLH = ?", 3201).Count(&count)
	var total_money []int64 //
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where  NOT (F_VC_KAWLH =?) and not F_NB_DABZT =?`
	db.Raw(sqlstr, 3201, 4).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询省外结算交易，结算表总交易笔数%d，查询总金额为：%d", count, total_money)
	return count, total_money[0]
}

//4.2.2	查询省内的已发送 总条数、总金额
func QueryShengnSendedJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201). /*Where().*/ Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4.2.3	查询省内已请款的数据总条数、总金额
func QueryShengnPleaseedJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201). /*Where().*/ Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//4.2.4	查询坏账（拒付）数据 总条数、总金额
func QueryShengnBadDebtsJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201). /*Where().*/ Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, count, total_money)
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
	db.Table("b_js_jiessj").Where("F_NB_QINGFJG = ?", 1).Not("F_VC_KAWLH = ?", 3201).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_NB_QINGFJG = ? and not F_VC_KAWLH = ?`
	db.Raw(sqlstr, 1, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询结算表含坏账总清分的交易笔数%d，查询已清分总金额为：%d", count, total_money)
	c, m := QueryShengwBadDebtsJieSuan()
	log.Printf("查询结算表不含坏账总清分的交易笔数%d，查询已清分总金额为：%d", count-c, total_money[0]-m)
	return count - c, total_money[0] - m
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
		log.Println("Insert b_jsjk_shengwqftj error", err)
		return err
	}
	log.Println("结算统计表插入成功！", Jiestj.FDtKaistjsj)
	return nil
}

//3、查询最新包含坏账的已清分数据 b_jsjk_shengwqftj
func QueryShengwClearingdata() (error, *types.BJsjkShengwqftj) {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkShengwqftj)
	//赋值
	if err := db.Table("b_jsjk_shengwqftj").Last(&Jiestj).Error; err != nil {
		log.Println("查询 省外清分统计表最新数据时 QueryShengwClearingdata error :", err)
		return err, nil
	}

	log.Println("查询省外清分统计表最新数据结果:", Jiestj)
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
		log.Println("更新最新的清分统计结果 error", err)
		return err
	}
	return nil
}

func QueryShengwClearingdataById(id int) (error, *types.BJsjkShengwqftj) {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsjkShengwqftj)
	//赋值
	if err := db.Table("b_jsjk_shengwqftj").Where("F_NB_ID=?", id).Last(&Jiestj).Error; err != nil {
		log.Println("查询 省外清分统计表最新数据时 QueryShengwClearingdata error :", err)
		return err, nil
	}

	log.Println("查询省外清分统计表最新数据结果:", Jiestj)
	return nil, Jiestj
}

// 查询 已清分的坏账 Bad debts
func QueryShengwBadDebtsJieSuan() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//1:已清分 F_NB_ZHENGYCLJG 争议处理结果：坏账2
	db.Table("b_js_jiessj").Where("F_NB_QINGFJG = ?", 1).Where("F_NB_ZHENGYCLJG = ?", 2).Not("F_VC_KAWLH = ?", 3201).Count(&count)

	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_NB_QINGFJG = ? and F_NB_ZHENGYCLJG = ? and not F_VC_KAWLH = ?`
	db.Raw(sqlstr, 1, 2, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询结算表已清分的坏账交易笔数%d，查询已清分的坏账总金额为：%d", count, total_money)
	return count, total_money[0]
}

//查询结算表坏账的数据【测试】
func QueryJieSuan() {
	db := utils.GormClient.Client
	js := make([]types.BJsJiessj, 0)
	//1:已清分 F_NB_ZHENGYCLJG 争议处理结果：坏账2
	db.Table("b_js_jiessj").Where("F_NB_ZHENGYCLJG = ?", 2).Find(&js)
	log.Println("查询结算表已清分的坏账交易笔数:", len(js), js)
}

//4.1.3	查询省外结算数据中存在争议的数据总条数、总金额
//1、查询省外结算数据中存在争议的数据总条数、总金额
func QueryDisputeJieSuanData() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//2：争议数据 0：争议数据未处理
	db.Table("b_js_jiessj").Where("F_NB_JIZJG  = ?", 2).Where("F_NB_ZHENGYCLJG = ?", 0).Not("F_VC_KAWLH = ?", 3201).Count(&count)

	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_NB_JIZJG = ?  and F_NB_ZHENGYCLJG = ? and not F_VC_KAWLH = ?`
	//2：争议数据 0：争议数据未处理
	db.Raw(sqlstr, 2, 0, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询结算表 待处理存在争议的数据总笔数:%d笔，查询待处理存在争议的数据总金额为：%d分", count, total_money)
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
		log.Println("Insert b_jsjk_shengwjszysjtj error", err)
		return err
	}
	log.Println("新增省外存在争议的数据的统计开始记录成功！", Jiestj.FDtKaistjsj)
	return nil
}

//3、查询最新的争议的数据记录
func QueryShengwDispute() (error, *types.BJsjkShengwjszysjtj) {
	db := utils.GormClient.Client
	Jiestjs := new(types.BJsjkShengwjszysjtj)
	//赋值
	if err := db.Table("b_jsjk_shengwjszysjtj").Last(&Jiestjs).Error; err != nil {
		log.Println("查询最新的争议的数据记录时，QueryShengwDispute error :", err)
		return err, nil
	}
	log.Println("查询最新的争议的数据记录结果:", Jiestjs)
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
		log.Println("最新的争议的数据记录 error", err)
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
		db.Table("b_zdz_chedckyssjlycb").Where("F_VC_SHANCBJ = ?", 0).Count(&zdzcount)

		zdzsqlstr := `select SUM(F_NB_JINE) as zdztotal_money from b_zdz_chedckyssjlycb where  F_VC_SHANCBJ = ?`

		db.Raw(zdzsqlstr, 0).Pluck("SUM(F_NB_JINE) as zdztotal_money", &zdztotal_money)

		log.Printf("查询总对总异常数据表 异常的交易笔数%d，查询异常的交易总金额为：%d", zdzcount, zdztotal_money)

		return zdzcount, zdztotal_money[0], nil
	}

	if lx == 2 {
		//出口异常表
		db.Table("b_dd_chedckyssjlycb").Where("F_VC_SHANCBJ = ?", 0).Count(&ddcount)
		sqlstr := `select SUM(F_NB_JINE) as ddtotal_money from b_dd_chedckyssjlycb  where F_VC_SHANCBJ = ?`
		db.Raw(sqlstr, 0).Pluck("SUM(F_NB_JINE) as ddtotal_money", &ddtotal_money)

		log.Printf("查询单点异常数据表 异常的交易笔数%d，查询异常的交易总金额为：%d", ddcount, ddtotal_money)

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
		log.Println("Insert b_jsjk_yicsjtj error", err)
		return err
	}
	log.Println("新增异常数据的数据的统计开始记录成功！", yctj.FDtKaistjsj)
	return nil
}

//3、查询最新的异常数据统计记录
func QueryAbnormaltable(lx int) (error, *types.BJsjkYicsjtj) {
	db := utils.GormClient.Client
	yctjs := new(types.BJsjkYicsjtj)
	//赋值
	if err := db.Table("b_jsjk_yicsjtj").Where("F_NB_TONGJLX=?", lx).Last(&yctjs).Error; err != nil {
		log.Println("查询最新的异常数据的数据记录时，QueryShengwDispute error :", err)
		return err, nil
	}
	log.Println("查询最新的异常数据的数据记录结果:", yctjs)
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
		log.Println("最新的异常数据的数据记录 error", err)
		return err
	}
	return nil
}

//4.1.10	清分、争议包更新状态监控
//1、查询清分包数据
func QueryClearlingdata(yesterday string) (error, *[]types.BJsQingftjxx) {
	db := utils.GormClient.Client
	qingftjsjs := make([]types.BJsQingftjxx, 0)
	//赋值
	if err := db.Table("b_js_qingftjxx").Where("F_DT_JIESSJ>=?", yesterday+" 00:00:00").Where("F_DT_JIESSJ<=?", yesterday+" 23:59:59").Find(&qingftjsjs).Error; err != nil {
		if fmt.Sprint(err) == "record not found" {
			log.Println("QueryClearlingdata err== `record not found`:", err)
			return nil, nil
		}
		log.Println("查询清分包数据 最新数据时 QueryClearlingdata error :", err)
		return err, nil
	}
	log.Println("查询清分包数据表结果:", qingftjsjs)
	return nil, &qingftjsjs
}

//2、查询争议处理包数据
func QueryDisputedata(yesterday string) (error, *types.BJsZhengyjyclxx) {
	db := utils.GormClient.Client
	zytjsj := new(types.BJsZhengyjyclxx)

	if err := db.Table("b_js_zhengyjyclxx").Where("F_DT_ZHENGYCLSJ>=?", yesterday+" 00:00:00").Where("F_DT_ZHENGYCLSJ<=?", yesterday+" 23:59:59").Last(&zytjsj).Error; err != nil {

		if fmt.Sprint(err) == "record not found" {
			log.Println("QueryDisputedata err == `record not found`:", err)
			return nil, nil
		}
		log.Println("查询争议处理包数据表最新数据时 QueryDisputedata error :", err)
		return err, nil
	}
	log.Println("查询争议处理包数据表结果:", zytjsj)
	return nil, zytjsj
}

//4.1.11	清分核对
//1、统计清分数据
func StatisticalClearlingcheck() error {
	//1、获取昨日的清分包数据
	today := time.Now().Format("2006-01-02")
	qerr, clears := QueryClearlingdata(today)
	if qerr != nil {
		return qerr
	}
	if clears == nil {
		log.Println("今日这时还没有收到清分包")
		return errors.New("今日这时还没有收到清分包，需要检查清分包是否接收")
	}
	for _, clear := range *clears {
		qcrerr := QueryCheckResult(clear.FNbXiaoxxh)
		if qcrerr != nil {
			if fmt.Sprint(qcrerr) == "查询清分核对结果成功,不能重复插入" {
				log.Println("查询清分核对结果成功,不能重复插入")
				return nil
			}
			return qcrerr
		}
		//2、统计昨日记账包总金额
		s1 := strings.Split(clear.FVcQingfmbr, "T")
		keepAccount, keepAccountCount := StatisticalkeepAccount(s1[0])
		//统计存在争议数据
		disputerr, Disput, zyfgsl := DisputedDataCanClearling(clear.FNbXiaoxxh)
		if disputerr != nil {
			return disputerr
		}
		//统计退费数据
		SWRefund := StatisticalRefund(today)
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
			log.Println("清分核对正确+++++")
		} else {
			is = 2
			log.Println("清分核对不正确+++++")
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

		data.FNbTuifje = int64(i)       //退费总金额  分
		data.FNbTuifts = SWRefund.Count //退费总条数
		s := strings.Split(clear.FVcQingfmbr, "T")
		data.FVcTongjrq = s[0] //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',【清分包的清分目标日】

		cherr := CheckResultInsert(data)
		if cherr != nil {
			return cherr
		}
		log.Println("清分金额核对完成+++++")
	}

	return nil
}

//记帐处理结果仅返回有争议（可疑）的交易记录明细。未包含在争议交易记录明细中的交易，均默认为发行方已确认可以付款。
//Amount：确认记帐总金额
//2、统计记账包总金额、
func StatisticalkeepAccount(Yesterdaydate string) (int64, int) {
	db := utils.GormClient.Client
	var total_money []int64
	//时间范围
	begin := Yesterdaydate + " 00:00:00"
	end := Yesterdaydate + " 23:59:59"
	sqlstr := `select SUM(F_NB_ZONGJE) as total_money from b_js_jizclxx  where F_DT_CHULSJ>=? and F_DT_CHULSJ<=?`
	db.Raw(sqlstr, begin, end).Pluck("SUM(F_NB_ZONGJE) as total_money", &total_money)
	log.Printf("统计记账包总金额为：%d", total_money)
	var totalcount []int
	sqlstr1 := `select SUM(F_NB_JILSL) as totalcount from b_js_jizclxx  where F_DT_CHULSJ>=? and F_DT_CHULSJ<=?`
	db.Raw(sqlstr1, begin, end).Pluck("SUM(F_NB_JILSL) as totalcount", &totalcount)
	log.Printf("统计记账包总条数为：%d", totalcount)

	var czzycount []int
	sqlstr2 := `select SUM(F_NB_ZHENGYSL) as czzycount from b_js_jizclxx  where F_DT_CHULSJ>=? and F_DT_CHULSJ<=?`
	db.Raw(sqlstr2, begin, end).Pluck("SUM(F_NB_ZHENGYSL) as czzycount", &czzycount)
	log.Printf("统计记账包存在争议的数量为：%d", czzycount)

	//if total_money == nil {
	//	log.Printf("total_money == nil" )
	//	return 0, 0
	//}
	return total_money[0], totalcount[0] - czzycount[0]
}

//3、统计清分包中可以清分的争议数据的金额
func DisputedDataCanClearling(qingfxiaoxiid int64) (error, *(types.BJsZhengyjyclxx), int) {
	db := utils.GormClient.Client
	zytjsj := new(types.BJsZhengyjyclxx)
	qingfmxsj := new(types.BJsQingftjmx)
	//yesterday:=	utils.Yesterdaydate()  Where("F_DT_JIESSJ>=?", yesterday+" 00:00:00").Where("F_DT_JIESSJ<=?", yesterday+" 23:59:59")
	if err := db.Table("b_js_qingftjmx").Where("F_NB_QINGFTJXXXH=?", qingfxiaoxiid).Last(&qingfmxsj).Error; err != nil {
		if fmt.Sprint(err) == "record not found" {
			log.Println("QueryClearlingdata err== `record not found`:", err)
			return nil, nil, 0
		} else {
			log.Println("查询 b_js_qingftjmx 表 最新数据时  error :", err)
			return err, nil, 0
		}
	}
	log.Println("查询清分包数据表结果:", qingfmxsj)

	if qingfmxsj.FNbZhengycljgwjid == 0 {
		log.Println("查询清分包 没有争议处理数据，全部可以记账清分++++++++++++++++++++++++++++++++++++++++ ")
		return nil, nil, 0
	} else {
		//争议处理消息
		if err := db.Table("b_js_zhengyjyclxx").Where("F_VC_ZHENGYJGWJID=?", qingfmxsj.FNbZhengycljgwjid).Last(&zytjsj).Error; err != nil {
			if fmt.Sprint(err) == "record not found" {
				log.Println("QueryDisputedata err == `record not found`:", err)
				return err, nil, 0
			} else {
				log.Println("查询争议处理包数据表最新数据时 QueryDisputedata error :", err)
				return err, nil, 0
			}
		}
	}

	log.Println("查询争议处理包数据表结果:", zytjsj)
	//	查询争议放过的数量
	zyfgerr, zyfgsl := QueryDisputedData(zytjsj)
	if zyfgerr != nil {
		return zyfgerr, nil, 0
	}

	return nil, zytjsj, zyfgsl
}

//4、统计退费总金额、总条数
func StatisticalRefund(daydate string) *types.SWRefund {
	db := utils.GormClient.Client
	var result types.SWRefund
	//时间范围
	//	begin := Yesterdaydate + " 00:00:00"
	//	end := Yesterdaydate + " 23:59:59"
	//	sqlstr := `select SUM(F_VC_TUIFZJE) as total,SUM(F_NB_TUIFJLZSL) as  count  from b_js_tuifxx  where F_VC_TUIFTJCLSJ>=? and F_VC_TUIFTJCLSJ<=?`
	sqlstr := `select SUM(F_VC_TUIFZJE) as total,SUM(F_NB_TUIFJLZSL) as  count  from b_js_tuifxx  where F_VC_TUIFQFMBR=?`

	db.Raw(sqlstr, daydate).Scan(&result)
	log.Printf("统计退费总金额为：%f", result.Total)
	log.Printf("统计退费总条数为：%d", result.Count)
	return &result
}

//争议处理消息 mx 数量
func QueryDisputedData(zyxx *types.BJsZhengyjyclxx) (error, int) {
	db := utils.GormClient.Client
	zytjsjmx := make([]types.BJsZhengyjyclmx, 0)
	//查询争议放过的数量
	if err := db.Table("b_js_zhengyjyclmx").Where("F_NB_ZHENGYJYCLXXXH=?", zyxx.FNbXiaoxxh).Where("F_NB_CHULJG=?", 0).Find(&zytjsjmx).Error; err != nil {
		if fmt.Sprint(err) == "record not found" {
			log.Println("QueryDisputedata err == `record not found`:", err)
			return err, 0
		}
		log.Println("查询争议处理包mx数量 error :", err)
		return err, 0
	}
	log.Println("争议处理消息 mx 数量 :", len(zytjsjmx))

	return nil, len(zytjsjmx)
}

//4、把核对结果插入数据库
func CheckResultInsert(clear *types.BJsjkQingfhd) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_qingfhd").Create(&clear).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_jsjk_qingfhd error", err)
		return err
	}
	log.Println("新增清分核对结果成功！++++", clear.FNbQingfbxh)
	return nil
}

func QueryCheckResult(Qingfbxh int64) error {
	db := utils.GormClient.Client
	Clear := new(types.BJsjkQingfhd)
	if err := db.Table("b_jsjk_qingfhd").Where("F_NB_QINGFBXH = ?", Qingfbxh).First(Clear).Error; err != nil {
		if fmt.Sprint(err) == "record not found" {
			log.Println("Query b_jsjk_qingfhd err == `record not found`:", err)
			return nil
		}
		log.Println("Query b_jsjk_qingfhd error", err)
		return err
	}
	log.Println("查询清分核对结果成功！++++", Clear.FNbQingfbxh)
	return errors.New("查询清分核对结果成功,不能重复插入")
}

//查询最新一条清分核对结果
func QueryCheckResultbyTs(ts int) (error, *[]types.BJsjkQingfhd) {
	db := utils.GormClient.Client
	qingfhddata := make([]types.BJsjkQingfhd, 0)
	if err := db.Table("b_jsjk_qingfhd").Order("F_NB_ID desc").Limit(ts).Find(&qingfhddata).Error; err != nil {
		log.Println("查询清分核对数据 最新数据时 QueryClearlingdata error :", err)
		return err, nil
	}
	log.Println("查询清分核对数据表结果:", qingfhddata)
	return nil, &qingfhddata
}

//4.1.8	省外结算数据分类
func QuerySWDataClassification() *types.DataClassification {
	db := utils.GormClient.Client
	//省外总数据

	swzcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Not("F_VC_KAWLH = ?", 3201).Count(&swzcount)
	log.Printf("查询省外结算交易，结算表总交易笔数:%d", swzcount)
	//坏账       1:已清分 F_NB_ZHENGYCLJG 争议处理结果：坏账2
	huaizcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_QINGFJG = ?", 1).Where("F_NB_ZHENGYCLJG = ?", 2).Not("F_VC_KAWLH = ?", 3201).Count(&huaizcount)
	log.Printf("查询结算表已清分的坏账交易笔数:%d ", huaizcount)

	//已清分
	yiqfcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_QINGFJG = ?", 1).Not("F_VC_KAWLH = ?", 3201).Count(&yiqfcount)
	log.Printf("查询结算表含坏账总清分的交易笔数:%d ", yiqfcount)
	log.Printf("查询结算表不含坏账总清分的交易笔数:%d ", yiqfcount-huaizcount)

	//结算表 已记账
	jzcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_JIZJG = ?", 1).Not("F_VC_KAWLH = ?", 3201).Count(&jzcount)
	log.Printf("查询结算表 已记账的交易笔数:%d ", jzcount)

	//存在争议
	zycount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_JIZJG = ?", 2).Where("F_NB_ZHENGYCLJG = ?", 0).Not("F_VC_KAWLH = ?", 3201).Count(&zycount)
	log.Printf("查询结算表 存在争议的交易笔数:%d ", zycount)

	//未打包数据
	weidbcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_DABZT = ?", 0).Not("F_VC_KAWLH = ?", 3201).Count(&weidbcount)
	log.Printf("查询结算表 未打包数据的交易笔数:%d ", weidbcount)

	//打包中数据
	dabzcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_DABZT = ?", 1).Not("F_VC_KAWLH = ?", 3201).Count(&dabzcount)
	log.Printf("查询结算表 打包中数据的交易笔数:%d ", dabzcount)

	//已打包数据
	yidbcount := 0
	db.Table("b_js_jiessj").Not("F_NB_DABZT = ?", 4).Where("F_NB_DABZT = ?", 2).Not("F_VC_KAWLH = ?", 3201).Count(&yidbcount)
	log.Printf("查询结算表 已打包数据的交易笔数:%d ", yidbcount)

	//已fs数据
	var yifscount []int
	zdzsqlstr := `select SUM(F_NB_JILSL) as yifscount from b_js_yuansjyxx where F_NB_FASZT = ? `
	db.Raw(zdzsqlstr, 2).Pluck("SUM(F_NB_JILSL) as yifscount", &yifscount)
	log.Printf("查询结算表 已发送的交易笔数:%d ", yifscount)
	var dataClassification types.DataClassification
	dataClassification.Yiqfcount = yiqfcount - huaizcount //已清分总条数（不含坏账）

	dataClassification.Shengwzcount = swzcount  //省外结算总数据
	dataClassification.Jizcount = jzcount       //记账
	dataClassification.Zhengycount = zycount    //争议
	dataClassification.Weidbcount = weidbcount  //未打包
	dataClassification.Yidbcount = yidbcount    //已打包
	dataClassification.Yifscount = yifscount[0] //已发送
	dataClassification.Huaizcount = huaizcount  //坏账

	return &dataClassification
}

//新增数据分类表记录
func InsertSWDataClassification() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengwjssjfl)
	data.FDtKaistjsj = utils.StrTimeToNowtime()
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate()
	if err := db.Table("b_jsjk_shengwjssjfl").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增数据分类表记录 Insert b_jsjk_shengwjssjfl error", err)
		return err
	}
	log.Println("新增数据分类表记录成功！")
	return nil
}

//查询最新一条
func QuerySWDataClassificationTable() (error, *types.BJsjkShengwjssjfl) {
	db := utils.GormClient.Client
	shujufl := new(types.BJsjkShengwjssjfl)
	if err := db.Table("b_jsjk_shengwjssjfl").Last(&shujufl).Error; err != nil {
		log.Println(" QuerySWDataClassificationTable error :", err)
		return err, nil
	}
	log.Println("查询省外结算数据分类表结果:", shujufl)
	return nil, shujufl
}

//更新记录
func UpdateSWDataClassificationTable(data *types.BJsjkShengwjssjfl, id int) error {
	db := utils.GormClient.Client
	swfltj := new(types.BJsjkShengwjssjfl)

	swfltj.FNbJiaoyzts = data.FNbJiaoyzts     //   `F_NB_JIAOYZTS` int DEFAULT NULL COMMENT '交易总条数',
	swfltj.FNbQingfsjts = data.FNbQingfsjts   //   `F_NB_QINGFSJTS` int DEFAULT NULL COMMENT '清分数据条数',
	swfltj.FNbJizsjts = data.FNbJizsjts       //   `F_NB_JIZSJTS` int DEFAULT NULL COMMENT '记账数据条数',
	swfltj.FNbZhengysjts = data.FNbZhengysjts //   `F_NB_ZHENGYSJTS` int DEFAULT NULL COMMENT '争议数据条数 待处理',
	swfltj.FNbWeidbsjts = data.FNbWeidbsjts   //   `F_NB_WEIDBSJTS` int DEFAULT NULL COMMENT '未打包数据条数',
	swfltj.FNbYidbsjts = data.FNbYidbsjts     //   `F_NB_YIDBSJTS` int DEFAULT NULL COMMENT '已打包数据条数',
	swfltj.FNbYifssjts = data.FNbYifssjts     //   `F_NB_YIFSSJTS` int DEFAULT NULL COMMENT '已发送数据条数',
	swfltj.FNbHuaizsjts = data.FNbHuaizsjts   //   `F_NB_HUAIZSJTS` int DEFAULT NULL COMMENT '坏账数据条数',
	swfltj.FDtTongjwcsj = data.FDtTongjwcsj   //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	swfltj.FVcTongjrq = data.FVcTongjrq       //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	if err := db.Table("b_jsjk_shengwjssjfl").Where("F_NB_ID=?", id).Updates(&swfltj).Error; err != nil {
		log.Println("更新 最新的省外结算数据分类 记录 时 error", err)
		return err
	}
	log.Println("更新 最新的省外结算数据分类 记录 成功++++++++++++++++++++++++++++++++++++++++++++++")
	return nil
}

//根据 id 查询表获取数据
func QuerySWDataClassificationTableByID(id int) (error, *types.BJsjkShengwjssjfl) {
	db := utils.GormClient.Client
	shujufl := new(types.BJsjkShengwjssjfl)
	if err := db.Table("b_jsjk_shengwjssjfl").Where("F_NB_ID = ?", id).Last(&shujufl).Error; err != nil {
		log.Println(" QuerySWDataClassificationTable error :", err)
		return err, nil
	}
	log.Println("查询 省外结算数据分类表结果:", shujufl)
	return nil, shujufl
}

//4.1.9	全天24小时转结算数监控
func QueryDataTurnMonitor() *types.TurnData {
	db := utils.GormClient.Client
	//1、查处出 b_dd_chedckyssj，b_zdz_chedckyssj 数据量  F_NB_JIAOYZT=1：转结算ok
	ddckzcount := 0
	db.Table("b_dd_chedckyssj").Where("F_NB_JIAOYZT = ?", 1).Where("F_DT_JIAOYSJ >= ?", "2020-09-01 :00:00:00").Count(&ddckzcount)
	log.Printf("查询单点出口表总交易笔数ddckzcount:%d", ddckzcount)

	zdzckzcount := 0
	db.Table("b_zdz_chedckyssj").Where("F_NB_JIAOYZT = ?", 1).Where("F_DT_JIAOYSJ >= ?", "2020-09-01 :00:00:00").Count(&zdzckzcount)
	log.Printf("查询总对总出口表总交易笔数zdzckzcount:%d", zdzckzcount)

	//2、查处b_js_jiessj  数据量
	jszcount := 0
	db.Table("b_js_jiessj").Where("F_DT_JIAOYSJ >= ?", "2020-09-01 :00:00:00").Count(&jszcount)
	log.Printf("查询结算表总交易笔数jszcount:%d", jszcount)
	turndata := new(types.TurnData)
	turndata.DDzcount = ddckzcount
	turndata.ZDZcount = zdzckzcount
	turndata.Jieszcount = jszcount
	return turndata
}

//新增转结算记录
func InsertDataTurnMonitor(lx int) error {
	db := utils.GormClient.Client
	data := new(types.BJsjkZhuanjssjjk)
	data.FNbTongjlx = lx
	data.FDtKaistjsj = utils.StrTimeToNowtime()
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate()
	if err := db.Table("b_jsjk_zhuanjssjjk").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增转结算记录Insert b_jsjk_zhuanjssjjk  error", err)
		return err
	}
	log.Println("3333333333333333-------新增转结算记录成功！")
	return nil
}

//查询转结算表最新记录
func QueryDataTurnMonitorTable(lx int) (error, *types.BJsjkZhuanjssjjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkZhuanjssjjk)
	if err := db.Table("b_jsjk_zhuanjssjjk").Where("F_NB_TONGJLX = ?", lx).Last(&shuju).Error; err != nil {
		log.Println(" QueryDataTurnMonitorTable error :", err)
		return err, nil
	}
	log.Println("查询转结算表最新记录结果:", shuju)
	return nil, shuju
}

func QueryDataTurnMonitorTableByID(id int) (error, *types.BJsjkZhuanjssjjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkZhuanjssjjk)
	if err := db.Table("b_jsjk_zhuanjssjjk").Where("F_NB_ID = ?", id).Last(&shuju).Error; err != nil {
		log.Println(" QueryDataTurnMonitorTable error :", err)
		return err, nil
	}
	log.Println("查询转结算表最新记录结果:", shuju)
	return nil, shuju
}

//更新转结算表数据
func UpdateDataTurnMonitorTable(data *types.BJsjkZhuanjssjjk, id int) error {
	db := utils.GormClient.Client
	zhuanjsjl := new(types.BJsjkZhuanjssjjk)
	zhuanjsjl.FNbChedyssjts = data.FNbChedyssjts //  `F_NB_CHEDYSSJTS` int DEFAULT NULL COMMENT '车道原始数据条数',
	zhuanjsjl.FNbJiesbsjts = data.FNbJiesbsjts   //  `F_NB_JIESBSJTS` int DEFAULT NULL COMMENT '结算表数据条数',
	zhuanjsjl.FNbTongjlx = data.FNbTongjlx       //  `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:单点、2:总对总',
	zhuanjsjl.FDtTongjwcsj = data.FDtTongjwcsj   //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	zhuanjsjl.FVcKuaizsj = data.FVcKuaizsj       //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
	if err := db.Table("b_jsjk_zhuanjssjjk").Where("F_NB_ID=?", id).Updates(&zhuanjsjl).Error; err != nil {
		log.Println("更新转结算表数据 记录 时 error", err)
		return err
	}
	log.Println("更新转结算表数据 记录 完成")
	return nil
}

//5、查询最新的ts条转结算表数据
func QueryDataTurnMonitortable(ts, lx int) (error, *[]types.BJsjkZhuanjssjjk) {
	db := utils.GormClient.Client
	hmdtjs := make([]types.BJsjkZhuanjssjjk, 0)
	//赋值Order("created_at desc")
	if err := db.Table("b_jsjk_zhuanjssjjk").Where("F_NB_TONGJLX = ?", lx).Order("F_NB_ID desc").Limit(ts).Find(&hmdtjs).Error; err != nil {
		log.Println("查询最新的ts条转结算表数据时，QueryBlacklisttable error :", err)
		return err, nil
	}
	log.Println("查询最新的ts条转结算表数据结果:", hmdtjs)
	return nil, &hmdtjs
}

//4.1.6	前30天省外结算趋势 每天记录一次，统计30天的数据
//查询昨日交易金额、清分金额；
func QuerySettlementTrend(datetime string) *types.ClearandJiesuan {
	db := utils.GormClient.Client
	//时间范围
	begin := datetime + " 00:00:00"
	end := datetime + " 23:59:59"
	jszcount := 0
	//昨日的交易条数
	db.Table("b_js_jiessj").Where("F_DT_JIAOYSJ >= ?", datetime+" 00:00:00").Where("F_DT_JIAOYSJ <= ?", datetime+" 23:59:59").Not("F_VC_KAWLH = ?", 3201).Count(&jszcount)
	log.Println("昨日的交易条数jszcount :", jszcount)
	//昨日的交易金额
	total_money := make([]int64, 1)

	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj  where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and not F_VC_KAWLH =? `
	db.Raw(sqlstr, begin, end, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)
	log.Println("昨日的交易金额total_money :", total_money[0])

	//昨日清分条数  F_NB_QINGFJG=1  F_NB_ZHENGYCLJG ！=2坏账
	qfzcount := 0
	//昨日的交易条数
	db.Table("b_js_jiessj").Where("F_NB_QINGFJG = ?", 1).Where("F_DT_JIAOYSJ >= ?", begin).Where("F_DT_JIAOYSJ <= ?", end).Not("F_VC_KAWLH = ?", 3201).Not("F_NB_ZHENGYCLJG = ?", 2).Count(&qfzcount)
	log.Println("昨日的清分条数qfzcount :", qfzcount)

	//昨日清分金额  F_NB_QINGFJG=1  F_NB_ZHENGYCLJG ！=2坏账
	qftotal_money := make([]int64, 1)

	qfsqlstr := `select SUM(F_NB_JINE) as qftotal_money from b_js_jiessj  where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and F_NB_QINGFJG = ? and not F_NB_ZHENGYCLJG  =? and not F_VC_KAWLH =?`
	db.Raw(qfsqlstr, begin, end, 1, 2, 3201).Pluck("SUM(F_NB_JINE) as qftotal_money", &qftotal_money)
	log.Println("昨日清分金额qftotal_money :", qftotal_money[0])
	log.Println("查询日期 datetime:", datetime)
	return &types.ClearandJiesuan{
		ClearlingCount: qfzcount,
		ClearlingMoney: qftotal_money[0],
		JiesuanCount:   jszcount,
		JiesuanMoney:   total_money[0],
		Datetime:       datetime,
	}

}

//获取30天的交易金额、条数、清分金额、条数   从小到大
func QuerySettlementTrendbyDay() *[]types.ClearandJiesuan {
	//获取时间 之前30天
	datetimes := utils.OldData(30)
	Data := make([]types.ClearandJiesuan, 0)
	//获取数据
	for _, d := range datetimes {
		data := QuerySettlementTrend(d)
		Data = append(Data, *data)
	}
	log.Println("查询30天的数据Data:", Data)
	//返回数据
	return &Data
}

//新增省外趋势表
func InsertSettlementTrendbyDayTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengwjsqs)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengwjsqs").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省外趋势表记录Insert b_jsjk_shengwjsqs  error", err)
		return err
	}
	log.Println("新增省外趋势表记录成功！")
	return nil
}

//查询最新记录
func QuerySettlementTrendbyDayTable() (error, *types.BJsjkShengwjsqs) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengwjsqs)
	if err := db.Table("b_jsjk_shengwjsqs").Last(&shuju).Error; err != nil {
		log.Println(" QuerySettlementTrendbyDayTable error :", err)
		return err, nil
	}
	log.Println("查询转结算表最新记录结果:", shuju)
	return nil, shuju
}

//更新数据
func UpdateSettlementTrendbyDayTable(data *types.BJsjkShengwjsqs, id int) error {
	db := utils.GormClient.Client
	qushijl := new(types.BJsjkShengwjsqs)

	qushijl.FNbJiaoye = data.FNbJiaoye       //   `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
	qushijl.FNbQingdje = data.FNbQingdje     //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	qushijl.FNbChae = data.FNbChae           //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	qushijl.FNbJiaoyts = data.FNbJiaoyts     //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	qushijl.FNbQingfts = data.FNbQingfts     //   `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
	qushijl.FDtTongjwcsj = data.FDtTongjwcsj //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	qushijl.FVcTongjrq = data.FVcTongjrq     //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	if err := db.Table("b_jsjk_shengwjsqs").Where("F_NB_ID=?", id).Updates(&qushijl).Error; err != nil {
		log.Println("更新省外结算趋势表数据 记录 时 error", err)
		return err
	}
	log.Println("更新省外结算趋势表数据 记录 成功+++++++++++++++++")
	return nil
}

//查询省外结算趋势表记录
func QuerySettlementTrendbyday(ts int) (error, *[]types.BJsjkShengwjsqs) {
	db := utils.GormClient.Client
	hmdtjs := make([]types.BJsjkShengwjsqs, 0)
	//赋值Order("created_at desc")
	if err := db.Table("b_jsjk_shengwjsqs").Order("F_NB_ID desc").Limit(ts).Find(&hmdtjs).Error; err != nil {
		log.Println("查询最新的ts条省外结算趋势表数据时，QuerySettlementTrendbyday error :", err)
		return err, nil
	}
	log.Println("查询最新的ts条转结算表数据结果:", hmdtjs)
	return nil, &hmdtjs
}

//4.1.5	数据包实时状态监控  10分钟查一次
func QueryPacketMonitoring() *types.PacketMonitoringdata {
	db := utils.GormClient.Client
	datetime := utils.DateNowFormat()
	//时间范围
	begin := datetime + " 00:00:00"
	end := datetime + " 23:59:59"

	//今日打包数量
	dbcount := 0
	db.Table("b_js_yuansjyxx").Where("F_DT_DABSJ >= ?", begin).Where("F_DT_DABSJ <= ?", end).Count(&dbcount)
	log.Printf("查询今日打包数量:%d", dbcount)

	//打包金额 F_NB_ZONGJE
	dbjine := make([]int64, 1)
	dbsqlstr := `select SUM(F_NB_ZONGJE) as dbjjine from b_js_yuansjyxx where F_DT_DABSJ >= ? and F_DT_DABSJ <= ? `
	db.Raw(dbsqlstr, begin, end).Pluck("SUM(F_NB_ZONGJE) as dbjjine", &dbjine)
	log.Printf("查询今日打包金额:%d", dbjine[0])

	//已发送原始交易消息包数量
	fscount := 0
	db.Table("b_js_yuansjyxx").Where("F_DT_FASSJ >= ?", begin).Where("F_DT_FASSJ <= ?", end).Count(&fscount)
	log.Printf("查询今日发送原始交易包数量:%d", fscount)

	//已发送原始交易消息包金额
	fsjine := make([]int64, 1)
	fssqlstr := `select SUM(F_NB_ZONGJE) as fsjjine from b_js_yuansjyxx where F_DT_FASSJ >= ? and F_DT_FASSJ <= ? `
	db.Raw(fssqlstr, begin, end).Pluck("SUM(F_NB_ZONGJE) as fsjjine", &fsjine)
	log.Printf("查询今日发送原始交易包金额:%d", fsjine[0])

	//记账包数量
	jzbcount := 0
	db.Table("b_js_jizclxx").Where("F_DT_JIESSJ >= ?", begin).Where("F_DT_JIESSJ <= ?", end).Count(&jzbcount)
	log.Printf("查询今日接收记账包数量:%d", jzbcount)
	//记账包金额 F_DT_JIESSJ
	jzbjine := make([]int64, 1)
	jzbsqlstr := `select SUM(F_NB_ZONGJE) as jzbjjine from b_js_jizclxx where F_DT_JIESSJ >= ? and F_DT_JIESSJ <= ? `
	db.Raw(jzbsqlstr, begin, end).Pluck("SUM(F_NB_ZONGJE) as jzbjjine", &jzbjine)
	log.Printf("查询今日接收记账包金额:%d", jzbjine[0])

	//原始交易消息应答包数量
	ysydbcount := 0
	db.Table("b_js_yuansjyydxx").Where("F_DT_XIAOXJSSJ >= ?", begin).Where("F_DT_XIAOXJSSJ <= ?", end).Count(&ysydbcount)
	log.Printf("查询今日原始交易消息应答包数量:%d", ysydbcount)

	return &types.PacketMonitoringdata{Dabaojine: dbjine[0], Dabaosl: dbcount,
		Fasbsl: fscount, Fasbjine: fsjine[0],
		Jizbjine: jzbjine[0], Jizbsl: jzbcount,
		Yuansbsl: ysydbcount,
	}
}

//新增数据包监控表记录
func InsertPacketMonitoringTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShujbjk)
	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shujbjk").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增数据包监控表记录 error", err)
		return err
	}
	log.Println("新增数据包监控表记录成功！")
	return nil
}

//查询最新一条
func QueryPacketMonitoringTable() (error, *types.BJsjkShujbjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShujbjk)
	if err := db.Table("b_jsjk_shujbjk").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条数据包监控表记录 error :", err)
		return err, nil
	}
	log.Println("查询数据包监控表最新记录结果:", shuju)
	return nil, shuju
}

//更新数据包监控记录；
func UpdatePacketMonitoringTable(data *types.BJsjkShujbjk, id int) error {
	db := utils.GormClient.Client
	shujub := new(types.BJsjkShujbjk)
	shujub.FNbDabsl = data.FNbDabsl               //   `F_NB_DABSL` int DEFAULT NULL COMMENT '打包数量',
	shujub.FNbDabje = data.FNbDabje               //   `F_NB_DABJE` bigint DEFAULT NULL COMMENT '打包金额',
	shujub.FNbFasysjybsl = data.FNbFasysjybsl     //   `F_NB_FASYSJYBSL` int DEFAULT NULL COMMENT '已发送原始交易消息包数量',
	shujub.FNbFasysjybje = data.FNbFasysjybje     //   `F_NB_FASYSJYBJE` bigint DEFAULT NULL COMMENT '已发送原始交易消息包金额',
	shujub.FNbJizbsl = data.FNbJizbsl             //   `F_NB_JIZBSL` int DEFAULT NULL COMMENT '记账包数量',
	shujub.FNbJizbje = data.FNbJizbje             //   `F_NB_JIZBJE` bigint DEFAULT NULL COMMENT '记账包金额',
	shujub.FNbYuansjyydbsl = data.FNbYuansjyydbsl //   `F_NB_YUANSJYYDBSL` int DEFAULT NULL COMMENT '原始交易消息应答包数量',
	shujub.FDtTongjwcsj = data.FDtTongjwcsj       //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	shujub.FVcKuaizsj = data.FVcKuaizsj           //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
	if err := db.Table("b_jsjk_shujbjk").Where("F_NB_ID=?", id).Updates(&shujub).Error; err != nil {
		log.Println("更新省外结算数据包表数据 记录 时 error", err)
		return err
	}
	log.Println("更新省外结算数据包表数据 记录 完成++++++++++++++++++++++")
	return nil
}

//查询最新的ts条省外数据包监控
func QueryPacketMonitoringtable(ts int) (error, *[]types.BJsjkShujbjk) {
	db := utils.GormClient.Client
	jg := make([]types.BJsjkShujbjk, 0)
	//赋值Order("created_at desc")
	if err := db.Table("b_jsjk_shujbjk").Order("F_NB_ID desc").Limit(ts).Find(&jg).Error; err != nil {
		log.Println("查询最新的ts条省外数据包监控表数据时，QueryPacketMonitoringtable error :", err)
		return err, nil
	}
	log.Println("查询最新的ts条省外数据包监控结果:", jg)
	return nil, &jg
}

//新增结算监控表记录
func InsertShengnJieSuanTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkJiestj)
	data.FNbKawlh = 3201
	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_jiestj").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增结算监控表记录 error", err)
		return err
	}
	log.Println("新增结算监控表记录成功！")
	return nil
}

//查询最新一条省内记录
func QueryShengnJieSuanTable() (error, *types.BJsjkJiestj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkJiestj)
	if err := db.Table("b_jsjk_jiestj").Where("F_NB_KAWLH =?", 3201).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条数据包监控表记录 error :", err)
		return err, nil
	}
	log.Println("查询数据包监控表最新记录结果:", shuju)
	return nil, shuju
}

//更新省内结算记录
func UpdateShengnJieSuanTable(data *types.BJsjkJiestj, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_jiestj").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新结算表数据 记录 时 error", err)
		return err
	}
	log.Println("省内结算总金额、总条数 更新结算表数据 记录 完成+++++++++++++")
	return nil
}

//4.2.2	查询省内的已发送 总条数、总金额
//省内发送数据
func QueryShengnSendjiessj() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//"F_NB_DABZT = ?",2 已发送用打包状态表示
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_DABZT = ?", 2).Not("F_NB_DABZT = ?", 4).Not("F_NB_DABZT = ?", 5).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ? and F_NB_DABZT = ? and not F_NB_DABZT = ? and not F_NB_DABZT = ?`
	db.Raw(sqlstr, 3201, 2, 4, 5).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，结算表已发送笔数%d，查询总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//新增省内已发送记录
func InsertShengnSendTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengnyfssjtj)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengnyfssjtj").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省内已发送记录 error", err)
		return err
	}
	log.Println("新增省内已发送记录成功！")
	return nil
}

//查询省内已发送记录
func QueryShengnSendTable() (error, *types.BJsjkShengnyfssjtj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnyfssjtj)
	if err := db.Table("b_jsjk_shengnyfssjtj").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内已发送记录 error :", err)
		return err, nil
	}
	log.Println("查询省内已发送记录最新结果:", shuju)
	return nil, shuju
}

//更新省内已发送记录
func UpdateShengnSendTable(data *types.BJsjkShengnyfssjtj, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengnyfssjtj").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省内已发送 记录 时 error", err)
		return err
	}
	log.Println("更新省内已发送记录 完成+++++++++++++++")
	return nil
}

//查询省内已发送记录ByID
func QueryShengnSendTableByID(id int) (error, *types.BJsjkShengnyfssjtj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnyfssjtj)
	if err := db.Table("b_jsjk_shengnyfssjtj").Where("F_NB_ID=?", id).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内已发送记录 error :", err)
		return err, nil
	}
	log.Println("查询省内已发送记录最新结果:", shuju)
	return nil, shuju
}

//4.2.4	查询坏账（拒付）数据 总条数、总金额
//1、省内拒付金额、条数
func QueryShengnRefusePay() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//"F_NB_ZHENGYCLJG =?",2 坏账状态表示
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_ZHENGYCLJG =?", 2).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ? and F_NB_ZHENGYCLJG =?`
	db.Raw(sqlstr, 3201, 2).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，省内拒付笔数%d，查询省内拒付总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//2、新增省内拒付记录
func InsertShengnRefusePayTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengnjfsjtj)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengnjfsjtj").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省内拒付记录 error", err)
		return err
	}
	log.Println("新增省内拒付记录 成功！")
	return nil
}

//3、查询最新一条数据
func QueryShengnRefusePayTable() (error, *types.BJsjkShengnjfsjtj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnjfsjtj)
	if err := db.Table("b_jsjk_shengnjfsjtj").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内拒付记录 error :", err)
		return err, nil
	}
	log.Println("查询省内拒付记录最新结果:", shuju)
	return nil, shuju
}

func QueryShengnRefusePayTableByID(id int) (error, *types.BJsjkShengnjfsjtj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnjfsjtj)
	if err := db.Table("b_jsjk_shengnjfsjtj").Where("F_NB_ID=?", id).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内拒付记录 error :", err)
		return err, nil
	}
	log.Println("查询省内拒付记录最新结果:", shuju)
	return nil, shuju
}

//4、更新最新一条
func UpdateShengnRefusePayTable(data *types.BJsjkShengnjfsjtj, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengnjfsjtj").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省内已发送 记录 时 error", err)
		return err
	}
	log.Println("更新省内已发送记录 完成++++++++++++++++++++++++")
	return nil
}

//查询省内已请款数据金额、条数
func QueryAlreadyPlease() (int, int64) {
	db := utils.GormClient.Client
	count := 0
	//"F_NB_ZHENGYCLJG =?",
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_QINGFJG =?", 1).Not("F_NB_ZHENGYCLJG = ?", 2).Not("F_NB_DABZT = ?", 4).Not("F_NB_DABZT = ?", 5).Count(&count)
	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ? and F_NB_QINGFJG = ?  and not F_NB_ZHENGYCLJG =? and not F_NB_DABZT =? and not F_NB_DABZT =?`
	db.Raw(sqlstr, 3201, 1, 2, 4, 5).Pluck("SUM(F_NB_JINE) as total_money", &total_money)

	log.Printf("查询卡网络号为%d，省内已请款数据笔数%d，查询省内已请款数据总金额为：%d", 3201, count, total_money)
	return count, total_money[0]
}

//新增省内已请款数据
func InsertShengnAlreadyPleaseTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengnqktj)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengnqktj").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增已请款数据 error", err)
		return err
	}
	log.Println("新增已请款数据 成功！")
	return nil
}

//3、查询省内已请款数据最新一条数据
func QueryShengnAlreadyPleaseTable() (error, *types.BJsjkShengnqktj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnqktj)
	if err := db.Table("b_jsjk_shengnqktj").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条已请款数据 error :", err)
		return err, nil
	}
	log.Println("查询省内已请款数据:", shuju)
	return nil, shuju
}

func QueryShengnAlreadyPleaseTableByID(id int) (error, *types.BJsjkShengnqktj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnqktj)
	if err := db.Table("b_jsjk_shengnqktj").Where("F_NB_ID=?", id).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条已请款数据 error :", err)
		return err, nil
	}
	log.Println("查询省内已请款数据:", shuju)
	return nil, shuju
}

//4、更新省内已请款数据最新一条
func UpdateShengnAlreadyPleaseTable(data *types.BJsjkShengnqktj, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengnqktj").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省已请款数据时 error", err)
		return err
	}
	log.Println("更新已请款数据 完成++++++++++++++++++")
	return nil
}

//4.2.8	省内结算数据分类
//查询省内结算分类
func QuerySNDataClassification() *types.ShengNDataClassification {
	db := utils.GormClient.Client
	//省内结算总数据

	snzcount := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Not("F_NB_DABZT = ?", 4).Not("F_NB_DABZT = ?", 5).Count(&snzcount)
	log.Printf("查询省内结算交易，结算表总交易笔数:%d", snzcount)

	//已请款数据
	qkcount := 0
	//"F_NB_ZHENGYCLJG =?",
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_QINGFJG =?", 1).Not("F_NB_DABZT = ?", 4).Not("F_NB_DABZT = ?", 5).Not("F_NB_ZHENGYCLJG = ?", 2).Count(&qkcount)
	log.Printf("查询省内结算表 已请款的交易笔数:%d ", qkcount)

	//未发送数据 "F_NB_DABZT = ?", 0
	wfscount := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_DABZT = ?", 0).Count(&wfscount)
	log.Printf("查询省内结算表 未发送的交易笔数:%d ", wfscount)

	//已发送数据
	fscount := 0
	//"F_NB_DABZT = ?",2 已发送用打包状态表示
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_DABZT = ?", 2).Count(&fscount)
	log.Printf("查询省内结算表 已发送的交易笔数:%d ", fscount)

	//拒付数据
	jfcount := 0
	//"F_NB_ZHENGYCLJG =?",2 坏账状态表示
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_ZHENGYCLJG =?", 2).Count(&jfcount)
	log.Printf("查询省内结算表 拒付的交易笔数:%d ", jfcount)

	return &types.ShengNDataClassification{Shengnzcount: snzcount, //结算总数据
		Yiqkcount:  qkcount,  //已清分总条数（不含坏账）
		Weifscount: wfscount, //未fas
		Yifscount:  fscount,  //已发送
		Jufuzcount: jfcount,  //坏账
	}
}

//新增省内结算数据分类
func InsertSNDataClassificationTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengnjssjfl)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengnjssjfl").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省内结算数据分类 error", err)
		return err
	}
	log.Println("新增省内结算数据分类 成功！")
	return nil
}

//3、查询省内已请款数据最新一条数据
func QuerySNDataClassificationTable() (error, *types.BJsjkShengnjssjfl) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnjssjfl)
	if err := db.Table("b_jsjk_shengnjssjfl").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内结算数据分类 error :", err)
		return err, nil
	}
	log.Println("查询省内结算数据分类:", shuju)
	return nil, shuju
}

func QuerySNDataClassificationTableByID(id int) (error, *types.BJsjkShengnjssjfl) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnjssjfl)
	if err := db.Table("b_jsjk_shengnjssjfl").Where("F_NB_ID=?", id).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内结算数据分类 error :", err)
		return err, nil
	}
	log.Println("查询省内结算数据分类:", shuju)
	return nil, shuju
}

//4、更新省内已请款数据最新一条
func UpdateSNDataClassificationTable(data *types.BJsjkShengnjssjfl, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengnjssjfl").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省内结算数据分类时 error", err)
		return err
	}
	log.Println("更新省内结算数据分类 完成++++++++++++++++++")
	return nil
}

//4.2.5	省内今日实时数据
func QueryRealTimeSettlementData() *types.RealTimeSettlementData {
	db := utils.GormClient.Client

	//省内产生的结算数据金额、条数
	snzcount := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Count(&snzcount)
	log.Printf("查询省内结算交易，结算表总交易笔数:%d", snzcount)

	var total_money []int64
	sqlstr := `select SUM(F_NB_JINE) as total_money from b_js_jiessj where F_VC_KAWLH = ?`
	db.Raw(sqlstr, 3201).Pluck("SUM(F_NB_JINE) as total_money", &total_money)
	log.Printf("查询卡网络号为%d，结算表今日总交易笔数%d，查询总金额为：%d", 3201, snzcount, total_money)

	//省内发送的数据金额、条数 2 已发送用打包状态表示
	fscount := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_DABZT = ?", 2).Count(&fscount)
	var fstotal_money []int64
	fssqlstr := `select SUM(F_NB_JINE) as fstotal_money from b_js_jiessj where F_VC_KAWLH = ? and F_NB_DABZT =?`
	db.Raw(fssqlstr, 3201, 2).Pluck("SUM(F_NB_JINE) as fstotal_money", &fstotal_money)

	log.Printf("查询卡网络号为%d，结算表已发送笔数%d，查询总金额为：%d", 3201, fscount, fstotal_money)

	//省内记账的数据金额、条数 F_NB_JIZJG=1
	jzcount := 0
	db.Table("b_js_jiessj").Where("F_VC_KAWLH = ?", 3201).Where("F_NB_JIZJG =?", 1).Count(&jzcount)
	var jztotal_money []int64
	jzsqlstr := `select SUM(F_NB_JINE) as jztotal_money from b_js_jiessj where F_VC_KAWLH = ? and F_NB_JIZJG =?`
	db.Raw(jzsqlstr, 3201, 1).Pluck("SUM(F_NB_JINE) as jztotal_money", &jztotal_money)

	log.Printf("查询卡网络号为%d，结算表已记账笔数%d，查询总金额为：%d", 3201, jzcount, jztotal_money)
	return &types.RealTimeSettlementData{
		Shengnjsjine: total_money[0],
		Shengnjssl:   snzcount,
		Fasjine:      fstotal_money[0],
		Fassl:        fscount,
		Jizjine:      jztotal_money[0],
		Jizsl:        jzcount,
	}
}

//新增省内今日实时数据
func InsertSNRealTimeSettlementDataTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengnsssjjk)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengnsssjjk").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省内今日实时数据 error", err)
		return err
	}
	log.Println("新增省内今日实时数据 成功！")
	return nil
}

//3、查询省内今日实时数据最新一条数据
func QuerySNRealTimeSettlementDataTable() (error, *types.BJsjkShengnsssjjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnsssjjk)
	if err := db.Table("b_jsjk_shengnsssjjk").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内今日实时数据 error :", err)
		return err, nil
	}
	log.Println("查询省内今日实时数据:", shuju)
	return nil, shuju
}

func QuerySNRealTimeSettlementData(ts int) (error, *[]types.BJsjkShengnsssjjk) {
	db := utils.GormClient.Client
	shuju := make([]types.BJsjkShengnsssjjk, ts)
	if err := db.Table("b_jsjk_shengnsssjjk").Order("F_NB_ID desc").Limit(ts).Find(&shuju).Error; err != nil {
		log.Println("查询省内今日实时数据 error :", err)
		return err, nil
	}
	log.Println("查询省内今日实时数据结果:", shuju)
	return nil, &shuju
}

//更具id查询记录
func QuerySNRealTimeSettlementDataTableByID(id int) (error, *types.BJsjkShengnsssjjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnsssjjk)
	if err := db.Table("b_jsjk_shengnsssjjk").Where("F_NB_ID=?", id).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内今日实时数据 error :", err)
		return err, nil
	}
	log.Println("查询省内今日实时数据:", shuju)
	return nil, shuju
}

//4、更新省内今日实时数据最新一条
func UpdateSNRealTimeSettlementDataTable(data *types.BJsjkShengnsssjjk, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengnsssjjk").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省内今日实时数据时 error", err)
		return err
	}
	log.Println("更新省内今日实时数据 完成+++++++++++++++++++")
	return nil
}

//4.2.6	前30日省内结算概览 一天一次 统计30条
//查询前30日的省内结算趋势
func QueryShengNSettlementTrend() *[]types.SNClearandJiesuan {
	ds := utils.OldData(30)
	datas := make([]types.SNClearandJiesuan, 0)
	for _, d := range ds {
		data := QueryShengNSettlementTrendData(d)
		datas = append(datas, *data)
	}
	return &datas
}

//省内结算趋势
func QueryShengNSettlementTrendData(datetime string) *types.SNClearandJiesuan {
	db := utils.GormClient.Client
	//时间范围
	begin := datetime + " 00:00:00"
	end := datetime + " 23:59:59"
	//查询省内交易金额 省内交易条数

	var result types.SNResult
	sqlstr := `select SUM(F_NB_JINE) as total ,count(F_NB_JINE) as count from b_js_jiessj where F_VC_KAWLH = ? and (F_DT_JIAOYSJ >= ?) and (F_DT_JIAOYSJ <= ?) AND  not substring( F_VC_KAH, 5, 4 ) = '2300' `
	db.Raw(sqlstr, 3201, begin, end).Scan(&result)
	log.Printf("查询卡网络号为%d，结算表总交易笔数%d，查询总金额为：%d", 3201, result.Count, result.Total)

	//省内请款金额 省内请款条数
	var qkresult types.SNResult
	qksqlstr := `select SUM(F_NB_JINE) as total ,count(F_NB_JINE) as count from b_js_jiessj where F_VC_KAWLH = ? and F_NB_QINGFJG = ?  and not F_NB_ZHENGYCLJG =? and (F_DT_JIAOYSJ >= ?) and (F_DT_JIAOYSJ <= ?) AND not  substring( F_VC_KAH, 5, 4 ) = '2300'`
	db.Raw(qksqlstr, 3201, 1, 2, begin, end).Scan(&qkresult)

	log.Printf("查询卡网络号为%d，省内已请款数据笔数%d，查询省内已请款数据总金额为：%d", 3201, qkresult.Count, qkresult.Total)
	return &types.SNClearandJiesuan{
		ClearlingMoney: qkresult.Total,
		ClearlingCount: qkresult.Count,
		JiesuanCount:   result.Count,
		JiesuanMoney:   result.Total,
		Datetime:       datetime,
	}
}

//新增省内结算趋势
func InsertShengNSettlementTrendTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengnjsqs)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengnjsqs").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省内结算趋势 error", err)
		return err
	}
	log.Println("新增省内结算趋势 成功！")
	return nil
}

//3、查询省内结算趋势最新一条数据
func QueryShengNSettlementTrendTable() (error, *types.BJsjkShengnjsqs) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengnjsqs)
	if err := db.Table("b_jsjk_shengnjsqs").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内结算趋势 error :", err)
		return err, nil
	}
	log.Println("查询省内结算趋势:", shuju)
	return nil, shuju
}

//4、更新省内结算趋势最新一条
func UpdateShengNSettlementTrendTable(data *types.BJsjkShengnjsqs, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengnjsqs").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省内结算趋势时 error", err)
		return err
	}
	log.Println("更新省内结算趋势 完成")
	return nil
}

func QueryShengNSettlementTrendtable(ts int) (error, *[]types.BJsjkShengnjsqs) {
	db := utils.GormClient.Client
	shujus := make([]types.BJsjkShengnjsqs, ts)
	if err := db.Table("b_jsjk_shengnjsqs").Order("F_NB_ID desc").Limit(ts).Find(&shujus).Error; err != nil {
		log.Println("查询最新一条省内结算趋势 error :", err)
		return err, nil
	}
	log.Println("查询省内结算趋势:", shujus)
	return nil, &shujus
}

//4.2.7	海岭数据同步监控 这个接口会因为
func postWithJson(tradestarttime string) *dto.SyncResponse {
	//post请求提交json数据 tradestarttime
	sync := dto.SyncRequest{tradestarttime, 2}
	sy, _ := json.Marshal(sync)
	//localhost:8092
	addr := types.HlsyncAddr
	resp, err := http.Post("http://"+addr+"/hl/syncStat/query", "application/json", bytes.NewBuffer([]byte(sy)))
	if err != nil {
		log.Println("查询海岭数据同步的时候error:", err)
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	Resp := new(dto.SyncResponse)

	unmerr := json.Unmarshal(body, Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error")
		return nil
	}

	log.Println("Post request with json result:", string(body), Resp)
	return Resp
}

func QueryDataSync() (int, int) {

	//查询海玲oracle数据库 B_TXF_CHEDXFYSSJ
	num := 0
	//num = oracledb.OrclQuerydata()
	//log.Println("oracle num:", num)
	//tradestarttime

	sr := postWithJson(types.Tradestarttime)
	if sr == nil {
		num = 0
	} else {
		num = sr.SyncData.Count
	}

	db := utils.GormClient.Client
	//查询结算数据 停车场id
	var result types.Result
	parkids := types.Parkids
	sqlstr := `select  count(F_NB_JINE) as count from b_js_jiessj where F_VC_TINGCCBH in (` + parkids + `) and F_DT_JIAOYSJ >='` + types.Tradestarttime + `'  and   F_NB_DABZT <> 4 and   F_NB_DABZT <> 5`

	log.Println("types.Tradestarttime", types.Tradestarttime, "parkids:", parkids, "sqlstr:", sqlstr)
	db.Raw(sqlstr).Scan(&result)
	log.Printf("查询海玲数据库数据量:%d，结算表数据同步数据量:=%v", num, result.Count)
	return num, result.Count
}

//新增数据同步监控表
func InsertDataSyncTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShujtbjk)

	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shujtbjk").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增数据同步监控表 error", err)
		return err
	}
	log.Println("新增数据同步监控表 成功！")
	return nil
}

//3、查询数据同步监控表最新一条数据
func QueryDataSyncTable() (error, *types.BJsjkShujtbjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShujtbjk)
	if err := db.Table("b_jsjk_shujtbjk").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条数据同步监控表 error :", err)
		return err, nil
	}
	log.Println("查询数据同步监控表:", shuju)
	return nil, shuju
}

func QueryDataSyncTableByID(id int) (error, *types.BJsjkShujtbjk) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShujtbjk)
	if err := db.Table("b_jsjk_shujtbjk").Where("F_NB_ID=?", id).Last(&shuju).Error; err != nil {
		log.Println("查询最新一条数据同步监控表 error :", err)
		return err, nil
	}
	log.Println("查询数据同步监控表:", shuju)
	return nil, shuju
}

func QueryDataSynctable(ts int) (error, *[]types.BJsjkShujtbjk) {
	db := utils.GormClient.Client
	shujus := make([]types.BJsjkShujtbjk, ts)
	if err := db.Table("b_jsjk_shujtbjk").Order("F_NB_ID desc").Limit(ts).Find(&shujus).Error; err != nil {
		log.Println("查询数据同步监控表 error :", err)
		return err, nil
	}
	log.Println("查询数据同步监控表:", shujus)
	return nil, &shujus
}

//4、更新数据同步监控表最新一条
func UpdateDataSyncTable(data *types.BJsjkShujtbjk, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shujtbjk").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新数据同步监控表时 error", err)
		return err
	}
	log.Println("更新数据同步监控表 完成")
	return nil
}

//4.2.9	异常数据停车场top10
func QueryAbnormalDataOfParking() (*[]types.Result, *[]types.Result) {
	db := utils.GormClient.Client
	var ddresult []types.Result
	//b_dd_chedckyssjlycb 、 b_zdz_chedckyssjlycb
	sqlstr := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_dd_chedckyssjlycb where F_VC_SHANCBJ = ? GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr, 0).Scan(&ddresult)
	log.Println("ddresult:", ddresult)

	var zdzresult []types.Result
	//b_dd_chedckyssjlycb 、 b_zdz_chedckyssjlycb
	sqlstr1 := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_zdz_chedckyssjlycb where F_VC_SHANCBJ = ? GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr1, 0).Scan(&zdzresult)
	log.Println("zdzresult:", zdzresult)
	return &ddresult, &zdzresult
}

//新增异常数据停车场
func InsertAbnormalDataOfParkingTable(lx int) error {
	db := utils.GormClient.Client
	data := new(types.BJsjkYicsjtcctj)

	data.FNbTongjlx = lx
	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_yicsjtcctj").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增异常数据停车场表 error", err)
		return err
	}
	log.Println("新增异常数据停车场表 成功！")
	return nil
}

//3、查询异常数据停车场最新一条数据
func QueryAbnormalDataOfParkingTable() (error, *types.BJsjkYicsjtcctj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkYicsjtcctj)
	if err := db.Table("b_jsjk_yicsjtcctj").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条异常数据停车场表 error :", err)
		return err, nil
	}
	log.Println("查询异常数据停车场表数据:", shuju)
	return nil, shuju
}

//异常数据
func QueryAbnormalDataOfParkingtable(data string, ts int) (error, *[]types.BJsjkYicsjtcctj) {
	db := utils.GormClient.Client
	shuju := make([]types.BJsjkYicsjtcctj, ts)
	if err := db.Table("b_jsjk_yicsjtcctj").Where("F_VC_KUAIZSJ=?", data).Order("F_NB_ZONGTS desc").Limit(ts).Find(&shuju).Error; err != nil {
		log.Println("查询异常数据停车场表 error :", err)
		return err, nil
	}
	log.Println("查询异常数据停车场表:", shuju)

	shujus := make([]types.BJsjkYicsjtcctj, ts)
	for i, yqsj := range shuju {
		s := "没有找到该停车场名称"
		if GetTingcc(yqsj.FVcTingccid) != "" {
			s = GetTingcc(yqsj.FVcTingccid)
		}
		shujus[i].FVcTingccid = s
		shujus[i].FVcKuaizsj = yqsj.FVcKuaizsj
		shujus[i].FDtKaistjsj = yqsj.FDtKaistjsj
		shujus[i].FNbZongts = yqsj.FNbZongts
		shujus[i].FNbId = yqsj.FNbId
		shujus[i].FDtTongjwcsj = yqsj.FDtTongjwcsj
		shujus[i].FDtKaistjsj = yqsj.FDtKaistjsj
		shujus[i].FNbZongje = yqsj.FNbZongje
	}
	return nil, &shujus
}

//4、更新异常数据停车场最新一条
func UpdateAbnormalDataOfParkingTable(data *types.BJsjkYicsjtcctj, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_yicsjtcctj").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新异常数据停车场表时 error", err)
		return err
	}
	log.Println("更新异常数据停车场表 完成")
	return nil
}

//4.2.10	逾期数据停车场top10  b_zdz_chedckyssjlycb】表中逾期数据【F_VC_YICLX =21
func QueryOverdueData() *[]types.Result {
	db := utils.GormClient.Client
	datetimes := utils.OldData(30)
	log.Println(datetimes)
	begin := datetimes[0] + " 00:00:00"
	end := datetimes[29] + " 23:59:59"
	var yuqiresult []types.Result
	sqlstr := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_zdz_chedckyssjlycb where F_VC_SHANCBJ = ? and F_VC_YICLX = ? and F_DT_CAIJSJ >= ? and F_DT_CAIJSJ <= ?  GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr, 0, 21, begin, end).Scan(&yuqiresult)
	log.Println("++++++++++++++++++++yuqiresult:", yuqiresult)
	return &yuqiresult
}

//新增逾期数据停车场
func InsertOverdueDataTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkYuqsjtj)
	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_yuqsjtj").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增逾期数据停车场表 error", err)
		return err
	}
	log.Println("新增逾期数据停车场表 成功！")
	return nil
}

//3、查询逾期数据停车场最新一条数据
func QueryOverdueDataTable() (error, *types.BJsjkYuqsjtj) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkYuqsjtj)
	if err := db.Table("b_jsjk_yuqsjtj").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条逾期数据停车场表 error :", err)
		return err, nil
	}
	log.Println("查询逾期数据停车场表:", shuju)
	return nil, shuju
}

//4、更新逾期数据停车场最新一条
func UpdateOverdueDataTable(data *types.BJsjkYuqsjtj, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_yuqsjtj").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新逾期数据停车场表时 error", err)
		return err
	}
	log.Println("更新逾期数据停车场表 完成")
	return nil
}

func QueryOverdueDatatable(data string, ts int) (error, *[]types.BJsjkYuqsjtj) {
	db := utils.GormClient.Client
	shuju := make([]types.BJsjkYuqsjtj, ts)
	if err := db.Table("b_jsjk_yuqsjtj").Where("F_VC_TONGJRQ=?", data).Order("F_NB_YUQZTS desc").Limit(ts).Find(&shuju).Error; err != nil {
		log.Println("查询逾期数据停车场表 error :", err)
		return err, nil
	}
	log.Println("查询逾期数据停车场表:", shuju)
	shujus := make([]types.BJsjkYuqsjtj, ts)
	for i, yqsj := range shuju {
		s := "没有找到该停车场名称"
		if GetTingcc(yqsj.FVcTingccid) != "" {
			s = GetTingcc(yqsj.FVcTingccid)
		}
		shujus[i].FVcTingccid = s
		shujus[i].FNbYuqzts = yqsj.FNbYuqzts
		shujus[i].FNbYuqzje = yqsj.FNbYuqzje
		shujus[i].FVcTongjrq = yqsj.FVcTongjrq
		shujus[i].FNbId = yqsj.FNbId
		shujus[i].FDtTongjwcsj = yqsj.FDtTongjwcsj
		shujus[i].FDtKaistjsj = yqsj.FDtKaistjsj
	}
	return nil, &shujus
}

//根据停车场编号 查询 停车场名称
func GetTingcc(tingccbh string) string {
	db := utils.GormClient.Client
	//停车场信息
	shuju := new(types.BTccTingcc)
	if err := db.Table("b_tcc_tingcc").Where("F_VC_TINGCCBH=?", tingccbh).Last(&shuju).Error; err != nil {
		log.Println("查询 逾期数据停车场名称 error :", err)
		return ""
	}
	log.Println("查询逾期数据停车场名称:", shuju.FVcMingc)
	return shuju.FVcMingc
}

//省外停车场结算趋势表
//b_jsjk_shengwtccjsqs
//获取30天的交易金额、条数、清分金额、条数   从小到大
func QuerySWSettlementTrendbyDay() *[][]types.ClearandJiesuanParkingdata {
	//获取时间 之前30天
	datetimes := utils.OldData(30)
	Data := make([][]types.ClearandJiesuanParkingdata, 0)
	//获取数据
	for _, d := range datetimes {
		//1天的数据
		data := QuerySWSettlementTrendOneday(d)
		Data = append(Data, *data)
	}
	log.Println("查询30天的数据Data:", Data)
	//返回数据
	return &Data
}
func QuerySWSettlementTrendbyOneDay() *[][]types.ClearandJiesuanParkingdata {
	//获取时间 之前30天
	datetimes := utils.OldData(1)
	Data := make([][]types.ClearandJiesuanParkingdata, 0)
	//获取数据
	for _, d := range datetimes {
		//1天的数据
		data := QuerySWSettlementTrendOneday(d)
		Data = append(Data, *data)
	}
	log.Println("查询30天的数据Data:", Data)
	//返回数据
	return &Data
}

func QuerySWSettlementTrendOneday(datetime string) *[]types.ClearandJiesuanParkingdata {
	db := utils.GormClient.Client
	//时间范围
	begin := datetime + " 00:00:00"
	end := datetime + " 23:59:59"
	var result []types.Result
	sqlstr4 := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and not F_VC_KAWLH =? GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr4, begin, end, 3201).Scan(&result)
	log.Println("省外停车场结算 产生交易 result:", result)

	var qfresult []types.Result // where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and F_NB_QINGFJG = ? and not F_NB_ZHENGYCLJG  =? and not F_VC_KAWLH =?
	qfsqlstr := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and F_NB_QINGFJG = ? and not F_NB_ZHENGYCLJG  =? and not F_VC_KAWLH =? GROUP BY F_VC_TINGCCBH `
	db.Raw(qfsqlstr, begin, end, 1, 2, 3201).Scan(&qfresult)
	log.Println("省外停车场结算 已清分 result:", qfresult)
	log.Println("查询日期 datetime:", datetime)

	datas := make([]types.ClearandJiesuanParkingdata, 0)
	var d types.ClearandJiesuanParkingdata
	for _, r := range result {
		d.Parkingid = r.Parkingid
		d.JiesuanMoney = r.Total
		d.JiesuanCount = r.Count
		datas = append(datas, d)
	}
	for i, qfr := range qfresult {
		if qfr.Parkingid == datas[i].Parkingid {
			datas[i].ClearlingMoney = qfr.Total
			datas[i].ClearlingCount = qfr.Count
		}
	}
	return &datas
}

func QuerySWSettlementTrendOne() *[]types.ClearandJiesuanParkingdata {
	db := utils.GormClient.Client
	//时间范围
	datetime := utils.Yesterdaydate()
	begin := datetime + " 00:00:00"
	end := datetime + " 23:59:59"
	var result []types.Result
	sqlstr4 := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and not F_VC_KAWLH =? GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr4, begin, end, 3201).Scan(&result)
	log.Println("省外停车场结算 产生交易 result:", result)

	var qfresult []types.Result // where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and F_NB_QINGFJG = ? and not F_NB_ZHENGYCLJG  =? and not F_VC_KAWLH =?
	qfsqlstr := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and F_NB_QINGFJG = ? and not F_NB_ZHENGYCLJG  =? and not F_VC_KAWLH =? GROUP BY F_VC_TINGCCBH `
	db.Raw(qfsqlstr, begin, end, 1, 2, 3201).Scan(&qfresult)
	log.Println("省外停车场结算 已清分 result:", qfresult)
	log.Println("查询日期 datetime:", datetime)

	datas := make([]types.ClearandJiesuanParkingdata, 0)
	var d types.ClearandJiesuanParkingdata
	for _, r := range result {
		d.Datetime = datetime
		d.Parkingid = r.Parkingid
		d.JiesuanMoney = r.Total
		d.JiesuanCount = r.Count
		datas = append(datas, d)
	}
	for i, qfr := range qfresult {
		if qfr.Parkingid == datas[i].Parkingid {
			datas[i].ClearlingMoney = qfr.Total
			datas[i].ClearlingCount = qfr.Count
		}
	}
	return &datas
}

//新增省外停车场结算
func InsertSWSettlementTrendTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengwtccjsqs)
	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengwtccjsqs").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省外停车场结算表 error", err)
		return err
	}
	log.Println("新增省外停车场结算表 成功！")
	return nil
}

//3、查询省外停车场结算最新一条数据
func QuerySWSettlementTrendTable() (error, *types.BJsjkShengwtccjsqs) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengwtccjsqs)
	if err := db.Table("b_jsjk_shengwtccjsqs").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省外停车场结算表 error :", err)
		return err, nil
	}
	log.Println("查询省外停车场结算表:", shuju)
	return nil, shuju
}

//4、更新省外停车场结算最新一条
func UpdateSWSettlementTrendTable(data *types.BJsjkShengwtccjsqs, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengwtccjsqs").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省外停车场结算表时 error", err)
		return err
	}
	log.Println("更新省外停车场结算表 完成+++++++++++++++++++++++")
	return nil
}

//省内停车场昨日结算趋势表
//b_jsjk_shengntccjsqs
func QuerySNSettlementTrendOne() *[]types.ClearandJiesuanParkingdata {
	db := utils.GormClient.Client
	//时间范围
	datetime := utils.Yesterdaydate()
	begin := datetime + " 00:00:00"
	end := datetime + " 23:59:59"
	var result []types.Result
	sqlstr4 := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and   F_VC_KAWLH =? GROUP BY F_VC_TINGCCBH `
	db.Raw(sqlstr4, begin, end, 3201).Scan(&result)
	log.Println("省内停车场结算趋势表 产生交易 result:", result)

	var qfresult []types.Result
	qfsqlstr := `select SUM(F_NB_JINE) as total,count(F_NB_JINE) as count ,F_VC_TINGCCBH  as  parkingid from b_js_jiessj where F_DT_JIAOYSJ >= ? and F_DT_JIAOYSJ <= ? and F_NB_QINGFJG = ? and not F_NB_ZHENGYCLJG  =? and  F_VC_KAWLH =? GROUP BY F_VC_TINGCCBH `
	db.Raw(qfsqlstr, begin, end, 1, 2, 3201).Scan(&qfresult)
	log.Println("省内停车场结算趋势表 已清分 result:", qfresult)
	log.Println("查询日期 datetime:", datetime)

	datas := make([]types.ClearandJiesuanParkingdata, 0)
	var d types.ClearandJiesuanParkingdata
	for _, r := range result {
		d.Datetime = datetime
		d.Parkingid = r.Parkingid
		d.JiesuanMoney = r.Total
		d.JiesuanCount = r.Count
		datas = append(datas, d)
	}
	for i, qfr := range qfresult {
		if qfr.Parkingid == datas[i].Parkingid {
			datas[i].ClearlingMoney = qfr.Total
			datas[i].ClearlingCount = qfr.Count
		}
	}
	return &datas
}

//新增省内停车场结算趋势表
func InsertSNSettlementTrendTable() error {
	db := utils.GormClient.Client
	data := new(types.BJsjkShengntccjsqs)
	data.FDtKaistjsj = utils.StrTimeToNowtime()      //开始
	data.FDtTongjwcsj = utils.StrTimeTodefaultdate() //
	if err := db.Table("b_jsjk_shengntccjsqs").Create(&data).Error; err != nil {
		// 错误处理...
		log.Println("新增省内停车场结算趋势表 error", err)
		return err
	}
	log.Println("新增省内停车场结算趋势表 成功！")
	return nil
}

//3、查询省内停车场结算趋势表最新一条数据
func QuerySNSettlementTrendTable() (error, *types.BJsjkShengntccjsqs) {
	db := utils.GormClient.Client
	shuju := new(types.BJsjkShengntccjsqs)
	if err := db.Table("b_jsjk_shengntccjsqs").Last(&shuju).Error; err != nil {
		log.Println("查询最新一条省内停车场结算趋势表 error :", err)
		return err, nil
	}
	log.Println("查询省内停车场结算趋势表:", shuju)
	return nil, shuju
}

//4、更新省内停车场结算趋势表最新一条
func UpdateSNSettlementTrendTable(data *types.BJsjkShengntccjsqs, id int) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jsjk_shengntccjsqs").Where("F_NB_ID=?", id).Updates(&data).Error; err != nil {
		log.Println("更新省内停车场结算趋势表时 error", err)
		return err
	}
	log.Println("更新省内停车场结算趋势表 完成++++++++++++++++++++++++++++++")
	return nil
}

//查询省外清分核对记录
func QuerySettlementclearlingcheck(ts int) (error, *[]types.BJsjkQingfhd) {
	db := utils.GormClient.Client
	hds := make([]types.BJsjkQingfhd, 0)
	//赋值Order("created_at desc")
	if err := db.Table("b_jsjk_qingfhd").Order("F_NB_ID desc").Limit(ts).Find(&hds).Error; err != nil {
		log.Println("查询最新的ts条省外清分核对数据时，QuerySettlementclearlingcheck error :", err)
		return err, nil
	}
	log.Println("查询最新的ts条省外清分核对结果:", hds)
	return nil, &hds
}

//查询 清分核对结果
func QueryClearlingcheck(req *dto.ReqQueryClarify) (error, *[]types.BJsjkQingfhd, int, int) {
	db := utils.GormClient.Client
	log.Println("req:", req)
	if req.BeginTime == "" {
		return errors.New("请输入开始查询时间"), nil, 0, 0
	}

	if req.EndTime == "" {
		return errors.New("请输入查询截止时间"), nil, 0, 0
	}

	if req.Prepage == 0 {
		return errors.New("请输入正确的每页展示记录数"), nil, 0, 0
	}
	//查询全部
	if req.CheckState == 0 {
		hmdtjs := make([]types.BJsjkQingfhd, 0)
		//赋值Order("created_at desc"),,FVcTongjrq
		//倒序默认
		if req.Orderstatus == 2 {
			//查询总数
			var result []types.Result
			sqlstr4 := `select count(F_NB_QINGFBXH) as count  from b_jsjk_qingfhd where F_VC_TONGJRQ >= ? and F_VC_TONGJRQ <= ?`
			db.Raw(sqlstr4, req.BeginTime, req.EndTime).Scan(&result)
			log.Println("查询总记录数:", result[0].Count)

			//根据总数，和prepage每页数量 生成分页总数
			totalpages := int(math.Ceil(float64(result[0].Count) / float64(req.Prepage))) //page总数
			if req.Currentpageid > totalpages {
				req.Currentpageid = totalpages
			}

			if req.Currentpageid <= 0 {
				req.Currentpageid = 1
			}

			if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Order("F_NB_ID desc").Offset((req.Currentpageid - 1) * req.Prepage).Limit(req.Prepage).Find(&hmdtjs).Error; err != nil {
				log.Println("查询最新的ts条省外清分核对数据时，QuerySettlementclearlingcheck error :", err)
				return err, nil, 0, 0
			}
			return nil, &hmdtjs, result[0].Count, totalpages
		} else {
			//查询总数
			var result []types.Result
			sqlstr4 := `select count(F_NB_QINGFBXH) as count  from b_jsjk_qingfhd where F_VC_TONGJRQ >= ? and F_VC_TONGJRQ <= ?`
			db.Raw(sqlstr4, req.BeginTime, req.EndTime).Scan(&result)
			log.Println("查询总记录数:", result[0].Count)

			//根据总数，和prepage每页数量 生成分页总数
			totalpages := int(math.Ceil(float64(result[0].Count) / float64(req.Prepage))) //page总数
			if req.Currentpageid > totalpages {
				req.Currentpageid = totalpages
			}
			if req.Currentpageid <= 0 {
				req.Currentpageid = 1
			}
			if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Order("F_NB_ID").Offset((req.Currentpageid - 1) * req.Prepage).Limit(req.Prepage).Find(&hmdtjs).Error; err != nil {
				log.Println("查询最新的ts条省外清分核对数据时，QuerySettlementclearlingcheck error :", err)
				return err, nil, 0, 0
			}
			return nil, &hmdtjs, result[0].Count, totalpages
		}
	}

	hmdtjs := make([]types.BJsjkQingfhd, 0)
	//赋值Order("created_at desc"),,FVcTongjrq
	//倒序默认
	if req.Orderstatus == 2 {
		//查询总数
		var result []types.Result
		sqlstr4 := `select count(F_NB_QINGFBXH) as count  from b_jsjk_qingfhd where F_VC_TONGJRQ >= ? and F_VC_TONGJRQ <= ? and F_NB_HEDJG =?`
		db.Raw(sqlstr4, req.BeginTime, req.EndTime, req.CheckState).Scan(&result)
		log.Println("查询总记录数:", result[0].Count)

		//根据总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(result[0].Count) / float64(req.Prepage))) //page总数
		if req.Currentpageid > totalpages {
			req.Currentpageid = totalpages
		}

		if req.Currentpageid <= 0 {
			req.Currentpageid = 1
		}

		if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Where("F_NB_HEDJG =?", req.CheckState).Order("F_NB_ID desc").Offset((req.Currentpageid - 1) * req.Prepage).Limit(req.Prepage).Find(&hmdtjs).Error; err != nil {
			log.Println("查询清分核对数据时 error :", err)
			return err, nil, 0, 0
		}
		return nil, &hmdtjs, result[0].Count, totalpages
	} else {
		//查询总数
		var result []types.Result
		sqlstr4 := `select count(F_NB_QINGFBXH) as count  from b_jsjk_qingfhd where F_VC_TONGJRQ >= ? and F_VC_TONGJRQ <= ?`
		db.Raw(sqlstr4, req.BeginTime, req.EndTime).Scan(&result)
		log.Println("查询总记录数:", result[0].Count)

		//根据总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(result[0].Count) / float64(req.Prepage))) //page总数
		if req.Currentpageid > totalpages {
			req.Currentpageid = totalpages
		}
		if req.Currentpageid <= 0 {
			req.Currentpageid = 1
		}
		if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Where("F_NB_HEDJG =?", req.CheckState).Order("F_NB_ID").Offset((req.Currentpageid - 1) * req.Prepage).Limit(req.Prepage).Find(&hmdtjs).Error; err != nil {
			log.Println("查询省外清分核对数据时 error :", err)
			return err, nil, 0, 0
		}
		return nil, &hmdtjs, result[0].Count, totalpages
	}
}

//根据条件查询 清分核对结果
func QueryClearlingcheckByConditions(req *dto.ReqClarifyExportExcel) (error, *[]types.BJsjkQingfhd) {
	db := utils.GormClient.Client
	log.Println("req:", req)
	if req.BeginTime == "" {
		return errors.New("请输入开始查询时间"), nil
	}
	if req.EndTime == "" {
		return errors.New("请输入查询截止时间"), nil
	}
	//查询全部
	if req.CheckState == 0 {
		hmdtjs := make([]types.BJsjkQingfhd, 0)
		//赋值Order("created_at desc"),,FVcTongjrq
		//倒序默认
		if req.Orderstatus == 2 {
			//查询总数
			if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Order("F_NB_ID desc").Find(&hmdtjs).Error; err != nil {
				log.Println("查询省外清分核对数据时  error :", err)
				return err, nil
			}
			return nil, &hmdtjs
		} else {
			//查询总数
			if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Order("F_NB_ID").Find(&hmdtjs).Error; err != nil {
				log.Println("查询省外清分核对数据时 error :", err)
				return err, nil
			}
			return nil, &hmdtjs
		}
	}

	hmdtjs := make([]types.BJsjkQingfhd, 0)
	//赋值Order("created_at desc"),,FVcTongjrq
	//倒序默认
	if req.Orderstatus == 2 {
		//查询总数
		if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Where("F_NB_HEDJG =?", req.CheckState).Order("F_NB_ID desc").Find(&hmdtjs).Error; err != nil {
			log.Println("查询清分核对数据时 error :", err)
			return err, nil
		}
		return nil, &hmdtjs
	} else {
		//查询总数
		if err := db.Table("b_jsjk_qingfhd").Where("F_VC_TONGJRQ >=?", req.BeginTime).Where("F_VC_TONGJRQ <=?", req.EndTime).Where("F_NB_HEDJG =?", req.CheckState).Order("F_NB_ID").Find(&hmdtjs).Error; err != nil {
			log.Println("查询省外清分核对数据时 error :", err)
			return err, nil
		}
		return nil, &hmdtjs
	}
}
