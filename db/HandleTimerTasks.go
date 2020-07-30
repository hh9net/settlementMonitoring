package db

import (
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"strconv"
	"time"
)

//goroutine1
//1定时任务 一天一次的
func HandleDayTasks() {
	tiker := time.NewTicker(time.Second * 30) //每15秒执行一下 一天一次的

	for {
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程1，处理一天一次的定时任务")
		//任务一
		//查询省外结算总金额、总笔数
		qerr := QuerTotalSettlementData()
		if qerr != nil {
			log.Println("查询省外结算总金额、总笔数定时任务:", qerr)
		}
		//任务二
		//查询省外已清分总金额、总笔数
		qcerr := QuerTotalClarify()
		if qcerr != nil {
			log.Println("查询省外已清分总金额、总笔数定时任务:", qcerr)
		}
		//任务三
		//查询停车场的总金额、总笔数
		qterr := QueryTingccJieSuan()
		if qterr != nil {
			log.Println("查询停车场的总金额、总笔数定时任务:", qterr)
		}
		//任务四
		//查询清分包、争议包的包号、接收时间
		qcderr := QueryClearlingAndDisputePackage()
		if qcderr != nil {
			log.Println("查询清分包、争议包的包号、接收时间定时任务:", qcderr)
		}
		//
	}
}

//goroutine2
//2定时任务 按小时的
func HandleHourTasks() {
	tiker := time.NewTicker(time.Second * 15) //每15秒执行一下

	for {
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程2，处理按小时的定时任务")
		////任务1 待处理争议数据
		//qderr := QueryShengwDisputedata()
		//if qderr != nil {
		//	log.Println("查询省外存在争议的总金额、总笔数定时任务 error:", qderr)
		//}
		////任务2 异常数据统计
		//qycerr := QueryAbnormaldata()
		//if qycerr != nil {
		//	log.Println("查询异常数据统计的总金额、总笔数定时任务 error:", qycerr)
		//}
		//任务3 处理黑名单统计
		qhmderr := QueryblacklistCount()
		if qhmderr != nil {
			log.Println("查询黑名单统计总数定时任务 error:", qhmderr)
		}

	}

}

//goroutine3
//3定时任务 按分钟的
func HandleMinutesTasks() {
	tiker := time.NewTicker(time.Second * 10) //每15秒执行一下

	for {
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程3，处理按分钟的定时任务")

	}

}

//1任务1
func QuerTotalSettlementData() error {
	//1、新增开始统计记录
	inerr := InsertTabledata(10000)
	if inerr != nil {
		log.Println("查询省外结算总金额、总笔数,新增开始统计的记录  error!", inerr)
		return inerr //不用返回前端
	}
	//2、查询最新记录
	qerr, sj := QueryTabledata(10000)
	if qerr != nil {
		log.Println("查询省外结算总金额、总笔数,查询最新的省外结算统计记录  error!", qerr)
		return qerr //不用返回前端
	}
	//3、获取省外结算统计数据
	zts, zje := QueryShengwJieSuan()
	//赋值
	data := new(types.BJsjkJiestj)
	data.FNbZongje = zje // `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	data.FNbZongts = zts // `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',

	data.FDtTongjwcsj = utils.StrTimeTotime(utils.DateTimeNowFormat()) // `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',

	data.FVcTongjrq = utils.DateNowFormat() //`F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	//4、更新最新统计记录
	uperr := UpdateTabledata(data, 10000, sj.FNbId)
	if uperr != nil {
		log.Println("db.UpdateTabledata error!", uperr)
		return uperr
	}

	//5、把数据更新到redis【覆盖】
	conn := utils.RedisInit() //初始化redis
	// key:"jiesuantotal"  value："金额｜总条数"
	rhseterr := utils.RedisSet(conn, "jiesuantotal", strconv.Itoa(int(zje))+"|"+strconv.Itoa(zts))
	if rhseterr != nil {
		return rhseterr
	}

	log.Println("更新省外结算统计最新统计记录成功")
	//返回数据赋值
	return nil
}

//1任务2
func QuerTotalClarify() error {
	//1、新增清分监控，开始统计记录
	inerr := ShengwClearingInsert()
	if inerr != nil {
		log.Println("查询省外已清分总金额、总笔数,新增清分监控开始统计的记录  error!", inerr)
		return inerr //不用返回前端
	}
	//2、查询最新数据
	qerr, cxjg := QueryShengwClearingdata()
	if qerr != nil {
		log.Println("查询省外已清分总金额、总笔数,查询最新数据时  error!", qerr)
		return qerr //不用返回前端
	}
	//3、查询包含坏账的已清分
	count, amount := QueryShengwClearingJieSuan()
	if count == 0 || amount == 0 {
		log.Println("查询省外已清分总金额、总笔数   error!", count, amount)
		return nil //不用返回前端
	}
	//4、查询坏账
	badcount, badamount := QueryShengwBadDebtsJieSuan()
	//赋值
	qingftj := new(types.BJsjkShengwqftj)
	qingftj.FNbZongje = amount                      //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	qingftj.FNbZongts = count                       //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	qingftj.FDtTongjwcsj = utils.StrTimeToNowtime() //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	qingftj.FVcTongjrq = utils.DateNowFormat()      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
	qingftj.FNbHuaizje = badamount                  //`F_NB_HUAIZJE` bigint DEFAULT NULL COMMENT '坏账金额',
	qingftj.FNbHuaizts = badcount                   //`F_NB_HUAIZTS` bigint DEFAULT NULL COMMENT '坏账条数',
	//5、更新清分监控表数据  更具查询结果cxjg的id更新记录
	uperr := UpdateShengwClearingdata(qingftj, cxjg.FNbId)
	if uperr != nil {
		log.Println("查询省外已清分总金额、总笔数,查询最新数据时  error!", uperr)
		return uperr //不用返回前端
	}

	log.Println("更新清分监控表数据成功")
	//返回数据赋值
	return nil
}

//1任务3 获取停车场总金额、总笔数
func QueryTingccJieSuan() error {
	//获取停车场总金额、总笔数
	result := QueryTingccJieSuandata()
	for _, r := range *result {
		//1、插入表新数据
		inerr := InsertTingjiesuan()
		if inerr != nil {
			return inerr
		}
		//2、查询 停车场结算数据统计表最新数据
		qterr, tingjs := QueryTingjiesuan()
		if qterr != nil {
			return qterr
		}
		//赋值
		jiestjsj := new(types.BJsjkTingccjssjtj)
		jiestjsj.FVcTingccid = tingjs.FVcTingccid        //停车场id 插入 redis
		jiestjsj.FNbZongje = r.Total                     //总金额   插入 redis
		jiestjsj.FNbZongts = r.Count                     //总数     插入 redis
		jiestjsj.FDtTongjwcsj = utils.StrTimeToNowtime() //统计完成时间
		jiestjsj.FVcTongjrq = utils.DateNowFormat()      //统计日期

		//3、更新停车场
		uperr := UpdateTingjiesuan(jiestjsj, r.Parkingid, tingjs.FNbId)
		if uperr != nil {
			return uperr
		}

		//4、更新到redis中
		conn := utils.RedisInit() //初始化redis
		// key:"jiesstatistical"  item: 停车场id  value："金额｜总条数"
		rhseterr := utils.RedisHSet(conn, "jiesstatistical", r.Parkingid, strconv.Itoa(int(r.Total))+"|"+strconv.Itoa(r.Count))
		if rhseterr != nil {
			return rhseterr
		}
	}
	return nil
}

//1任务4 查询清分、争议处理包
func QueryClearlingAndDisputePackage() error {

	//1、获取清分包、争议包数据
	qcerr, clear := QueryClearlingdata()
	if qcerr != nil {
		return qcerr
	}
	Clear := types.ClearlingAndDispute{
		DataType:  "clear",
		PackageNo: strconv.Itoa(int(clear.FNbXiaoxxh)),
		DateTime:  utils.DateTimeFormat(clear.FDtJiessj),
	}
	m := make(map[string]string, 0)
	// key:日期    value:"包号"｜"时间"

	m[utils.Yesterdaydate()] = Clear.PackageNo + "|" + Clear.DateTime
	//2、把数据存储于redis  接收时间、包号
	hmseterr := utils.RedisHMSet(utils.RedisInit(), Clear.DataType, m)
	if hmseterr != nil {
		return hmseterr
	}

	qderr, dispute := QueryDisputedata()
	if qderr != nil {
		return qderr
	}

	Disput := types.ClearlingAndDispute{
		DataType:  "disput",
		PackageNo: strconv.Itoa(int(dispute.FNbXiaoxxh)),
		DateTime:  utils.DateTimeFormat(dispute.FDtJiessj),
	}
	//2、把数据存储于redis  接收时间、包号
	m[utils.Yesterdaydate()] = Disput.PackageNo + "|" + Disput.DateTime

	dishmseterr := utils.RedisHMSet(utils.RedisInit(), Disput.DataType, m)
	if dishmseterr != nil {
		return dishmseterr
	}
	return nil
}

//2任务1存在争议数据统计
func QueryShengwDisputedata() error {
	//1、新增存在争议监控，开始统计记录
	inerr := ShengwDisputeInsert()
	if inerr != nil {
		log.Println("查询省外争议数据统计总金额、总笔数,新增在争议数据开始统计的记录  error!", inerr)
		return inerr //不用返回前端
	}
	//2、查询最新数据
	qerr, zytjjg := QueryShengwDispute()
	if qerr != nil {
		log.Println("查询省外存在争议总金额、总笔数,查询最新数据时  error!", qerr)
		return qerr //不用返回前端
	}

	//3、查询争议数据的统计结果
	zycount, zyamount := QueryDisputeJieSuanData()
	//赋值
	zytj := new(types.BJsjkShengwjszysjtj)
	zytj.FNbZongje = zyamount                    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	zytj.FNbZongts = zycount                     //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	zytj.FDtTongjwcsj = utils.StrTimeToNowtime() //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	zytj.FVcTongjrq = utils.DateNowFormat()      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	//4、更新争议数据的统计结果  更具查询结果的id更新记录
	uperr := UpdateShengwDispute(zytj, zytjjg.FNbId)
	if uperr != nil {
		log.Println("更新省外争议数据的总金额、总笔数,更新最新数据时  error!", uperr)
		return uperr //不用返回前端
	}

	log.Println("更新最新争议数据的统计结果成功")
	//返回数据赋值
	return nil
}

//2任务2 异常数据统计执行
func QueryAbnormaldata() error {
	qzdzycerr := QueryAbnormaltabledata(1)
	if qzdzycerr != nil {
		log.Println("查询总对总异常数据统计的总金额、总笔数定时任务 error:", qzdzycerr)
		return qzdzycerr
	}

	qddycerr := QueryAbnormaltabledata(2)
	if qddycerr != nil {
		log.Println("查询单点异常数据统计的总金额、总笔数定时任务 error:", qddycerr)
		return qddycerr
	}
	return nil
}

//2任务2.1 统计异常数据方法 1:zdz 2:dd
func QueryAbnormaltabledata(lx int) error {
	//1、新增异常数据，开始统计记录
	inerr := AbnormalDataInsert(lx)
	if inerr != nil {
		log.Println("查询异常数据统计总金额、总笔数,新增在争议数据开始统计的记录  error!", inerr)
		return inerr //不用返回前端
	}
	//2、查询最新数据
	qerr, yctjjg := QueryAbnormaltable(lx)
	if qerr != nil {
		log.Println("查询异常数据总金额、总笔数,查询最新数据时  error!", qerr)
		return qerr //不用返回前端
	}
	//3、查询异常数据统计结果[最新的数据]
	zycount, zyamount, qzyerr := QueryAbnormalData(lx)
	if qzyerr != nil {
		log.Println(qzyerr)
		return qzyerr
	}
	//赋值
	zytj := new(types.BJsjkYicsjtj)
	zytj.FNbZongje = zyamount                    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	zytj.FNbZongts = zycount                     //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	zytj.FDtTongjwcsj = utils.StrTimeToNowtime() //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	zytj.FVcTongjrq = utils.DateNowFormat()      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	//4、更新异常数据的统计结果  更具查询结果的id更新记录
	uperr := UpdateAbnormalData(zytj, yctjjg.FNbId)
	if uperr != nil {
		log.Println("更新异常数据的总金额、总笔数,更新最新数据时  error!", uperr)
		return uperr //不用返回前端
	}

	log.Println("更新最新异常数据的统计结果成功")
	//返回数据赋值
	return nil
}

//2任务3 统计黑名单
func QueryblacklistCount() error {
	//1 新增黑名单总记录的统计开始记录
	inerr := BlacklistDataInsert()
	if inerr != nil {
		return inerr
	}
	//2查询最新的黑名单总记录统计记录
	qhmderr, hmdjl := QueryBlacklisttable()
	if qhmderr != nil {
		return qhmderr
	}
	//3 获取黑名单总数
	qerr, count := QueryBlacklistcount()
	if qerr != nil {
		return qerr
	}
	//赋值
	hmdtj := new(types.BJsjkHeimdjk)
	hmdtj.FNbHeimdzs = count
	hmdtj.FDtTongjwcsj = utils.StrTimeToNowtime()
	hmdtj.FVcKuaizsj = utils.KuaizhaoTimeNowFormat()

	//4、更新 更新最新的黑名单总记录统计记录
	uperr := UpdateBlacklistlData(hmdtj, hmdjl.FNbId)
	if uperr != nil {
		return uperr
	}

	return nil
}
