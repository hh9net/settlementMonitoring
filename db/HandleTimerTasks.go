package db

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"strconv"
	"strings"
	"time"
)

//goroutine1
//1定时任务 一天一次的【都去重了】
func HandleDayTasks() {
	for {
		now := time.Now()               //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时

		next = time.Date(next.Year(), next.Month(), next.Day(), 8, 0, 0, 0, next.Location()) //获取下一个凌晨的日期

		t := time.NewTimer(next.Sub(now)) //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		log.Println("执行线程1，处理一天一次的定时任务11111111111111111111111111111111111111111111111111111111111111111")

		//任务三
		//查询停车场的总金额、总笔数
		qterr := QueryTingccJieSuan()
		if qterr != nil {
			log.Println("+++++++++++++++++++++【1.1error】+++++++++++++++++=查询停车场的总金额、总笔数定时任务:", qterr)
		}
		//任务四
		//查询清分包、争议包的包号、接收时间  使用redis记录的
		qcderr := QueryClearlingAndDisputePackage()
		if qcderr != nil {
			log.Println("+++++++++++++++++++++【1.2error】+++++++++++++++++=查询清分包、争议包的包号、接收时间定时任务:", qcderr)
		}
		//任务五
		//清分核对[已经去重了]
		cherr := StatisticalClearlingcheck()
		if cherr != nil {
			log.Println("+++++++++++++++++++++【1.3error】+++++++++++++++++=清分核对定时任务:", cherr)
		}

		//任务七
		//省外结算趋势
		qserr := SettlementTrendbyDay()
		if qserr != nil {
			log.Println("+++++++++++++++++++++【1.4error】+++++++++++++++++=查询省外结算趋势定时任务 error:", qserr)
		}

		//省内业务
		//任务十三
		//省内结算趋势
		qsnqserr := QueryShengNSettlementTrenddata()
		if qsnqserr != nil {
			log.Println("+++++++++++++++++++++【1.5error】+++++++++++++++++=查询省内结算分类 定时任务 error:", qsnqserr)
		}

		//任务 十四
		//逾期数据
		yuqierr := Overduedata()
		if yuqierr != nil {
			log.Println("+++++++++++++++++++++【1.6error】+++++++++++++++++=查询逾期数据 定时任务 error:", yuqierr)
		}

		//任务 十五
		//省外停车场结算趋势
		qwqserr := SWSettlementTrendbyDay()
		if qwqserr != nil {
			log.Println("+++++++++++++++++++++【1.7error】+++++++++++++++++=查询省外停车场结算趋势 定时任务 error:", qwqserr)
		}
		//1.16 省内停车场结算趋势
		snqserr := SNSettlementTrendbyDay()
		if snqserr != nil {
			log.Println("+++++++++++++++++++++【1.8error】+++++++++++++++++=查询省内停车场结算趋势 定时任务 error:", snqserr)
		}

		log.Println("执行线程1，处理一天一次的定时任务【完成】11111111111111111111111111111111111111111111111111111111111111111")

	}
}

//goroutine2
//2定时任务 按小时的
func HandleHourTasks() {
	tiker := time.NewTicker(time.Minute * 60) //每15秒执行一下

	for {
		log.Println("执行线程2，处理按小时的定时任务222222222222222222222222222222222222222222222222")
		//任务一
		//查询省外结算总金额、总笔数
		qerr := QuerTotalSettlementData()
		if qerr != nil {
			log.Println("+++++++++++++++++++++【2.1error】+++++++++++++++++=查询省外结算总金额、总笔数定时任务:", qerr)
		}
		//任务二
		//查询省外已清分总金额、总笔数(不含坏账)
		qcerr := QuerTotalClarify()
		if qcerr != nil {
			log.Println("+++++++++++++++++++++【2.2error】+++++++++++++++++=查询省外已清分总金额、总笔数定时任务:", qcerr)
		}
		//任务1 待处理争议数据
		qderr := QueryShengwDisputedata()
		if qderr != nil {
			log.Println("查询省外存在争议的总金额、总笔数定时任务【2.3error】 error:", qderr)
		}
		//任务2 异常数据统计
		qycerr := QueryAbnormaldata()
		if qycerr != nil {
			log.Println("查询异常数据统计的总金额、总笔数定时任务 【2.4error】error:", qycerr)
		}

		//任务4 海玲数据同步
		tberr := QueryDataSyncdata()
		if tberr != nil {
			log.Println("+++++++++++++++++++++【2.5error】+++++++++++++++++查询海玲数据同步定时任务 error:", tberr)
		}

		//任务5 异常数据top10
		yctoperr := QueryAbnormalDataOfParkingdata()
		if yctoperr != nil {
			log.Println("+++++++++++++++++++++【2.6error】+++++++++++++++++查询异常数据top10定时任务 error:", yctoperr)
		}

		//任务六
		//查询省外数据分类查询
		qdcerr := DataClassification()
		if qdcerr != nil {
			log.Println("+++++++++++++++++++++【2.7error】+++++++++++++++++=数据分类查询定时任务 error:", qdcerr)
		}

		//任务 五
		//省内发送结算数据金额、条数
		snjsfserr := ShengnSendJieSuanData()
		if snjsfserr != nil {
			log.Println("+++++++++++++++++++++【2.8error】+++++++++++++++++=查询省内发送结算数据金额、条数 定时任务 error:", snjsfserr)
		}

		//任务 八
		//省内结算总金额、总条数监控
		snjserr := ShengnJieSuanData()
		if snjserr != nil {
			log.Println("+++++++++++++++++++++【2.9error】+++++++++++++++++=查询省内结算总金额、总条数-监控定时任务 error:", snjserr)
		}

		//任务 十二
		//查询省内结算分类
		qjsflerr := QuerySNDataClassificationData()
		if qjsflerr != nil {
			log.Println("+++++++++++++++++++++【2.10error】+++++++++++++++++=查询省内结算分类 定时任务 error:", qjsflerr)
		}

		//任务 十
		//省内拒付数据金额、条数
		snjferr := QueryShengnRefusePayData()
		if snjferr != nil {
			log.Println("+++++++++++++++++++++【2.11error】+++++++++++++++++=查询省内拒付数据金额、条数 定时任务 error:", snjferr)
		}

		//任务一 转结算24小时监控
		dterr := DataTurnMonitor()
		if dterr != nil {
			log.Println("省外之前24小时转结算定时任务 【2.12error】error:", dterr)
		}

		//任务3 处理黑名单统计
		qhmderr := QueryblacklistCount()
		if qhmderr != nil {
			log.Println("查询黑名单统计总数定时任务【2.13error】 error:", qhmderr)
		}
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程2，处理按小时的定时任务【完成】222222222222222222222222222222222222222222222222")

	}

}

//goroutine3
//3定时任务 按分钟的
func HandleMinutesTasks() {
	tiker := time.NewTicker(time.Minute * 10) //每15秒执行一下

	for {
		log.Println("执行线程3，处理按分钟的定时任务333333333333333333333333333333333333333333333333333333333333333333")

		//任务一 数据包监控
		perr := PacketMonitoring()
		if perr != nil {
			log.Println("省外结算数据包监控定时任务 【3.1error】error:", perr)
		}
		//任务二 省内实时数据监控[统计全库数据，做差值]
		snsserr := ShengNRealTimeSettlementData()
		if snsserr != nil {
			log.Println("省内实时数据监控 定时任务 【3.2error】error:", snsserr)
		}

		//任务 三
		//查询省内已请款数据金额、条数
		qqingkerr := QueryShengnAlreadyPleaseData()
		if qqingkerr != nil {
			log.Println("+++++++++++++++++++++【3.3error】+++++++++++++++++=f查询省内已请款数据金额、条数 定时任务 error:", qqingkerr)
		}

		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程3，处理按分钟的定时任务【完成】333333333333333333333333333333333333333333333333333333333333333333")

	}

}

//goroutine4
//3定时任务 按分钟的
func HandleKafka() {
	//tiker := time.NewTicker(time.Second * 10)
	//for {
KafkaI:
	log.Println("执行go程 处理kafka数据++++++++++++++++++++++++【kafka执行】+++++++++++++++++++++++++++++++++处理kafka数据")
	//处理kafka数据
	err := utils.ConsumerGroup()
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++【执行go程 处理kafka数据】 error :", err)
		goto KafkaI
	}
	//log.Println(<-tiker.C)
	//}

}

//1任务1
func QuerTotalSettlementData() error {
	//0、查询最新记录[插入之前先做校验该天是否有新增数据数据]
	qerr, sj := QueryTabledata(10000)
	if qerr != nil {
		log.Println("查询省外结算总金额、总笔数,查询最新的省外结算统计记录  error!", qerr)
		return qerr //不用返回前端
	}
	s1 := utils.DateNowFormat()
	if sj.FVcTongjrq == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}

	//1、新增开始统计记录
	inerr := InsertTabledata(10000)
	if inerr != nil {
		log.Println("查询省外结算总金额、总笔数,新增开始统计的记录  error!", inerr)
		return inerr //不用返回前端
	}
	//2、查询最新记录
	qerr, sj = QueryTabledata(10000)
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
	conn := utils.Pool.Get()
	//	conn := utils.RedisConn //初始化redis
	// key:"jiesuantotal"  value："金额｜总条数"
	rhseterr := utils.RedisSet(&conn, "swjiesuantotal", strconv.Itoa(int(zje))+"|"+strconv.Itoa(zts))
	if rhseterr != nil {
		return rhseterr
	}
	defer func() {
		_ = conn.Close()

	}()
	log.Println("更新省外结算统计最新统计记录成功++++++++++++++++++++++++++++[2.1]++++++++++++++++++++++++++++++++++")
	//返回数据赋值
	return nil
}

//1任务2
func QuerTotalClarify() error {
	//0、查询最新记录[插入之前先做校验该天是否有新增数据数据]
	qerr, sj := QueryShengwClearingdata()
	if qerr != nil {
		log.Println("查查询省外已清分总金额、总笔数,查询最新数据时  error!", qerr)
		return qerr //不用返回前端
	}
	s1 := utils.DateNowFormat()
	if sj.FVcTongjrq == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}

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
	//3、查询不包含坏账的已清分
	count, amount := QueryShengwClearingJieSuan()
	if count == 0 || amount == 0 {
		log.Println("查询省外已清分总金额、总笔数   error!", count, amount)
		return errors.New("查询省外已清分总金额、总笔数为0值") //不用返回前端
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

	log.Println("更新清分监控表数据成功+++++++++++++++++++++++[2.2]++++++++++++++++++++++++++")
	//返回数据赋值
	return nil
}

//1任务3 获取停车场总金额、总笔数
func QueryTingccJieSuan() error {
	qterr, sj := QueryTingjiesuan()
	if qterr != nil {
		return qterr
	}
	s1 := utils.DateNowFormat()
	if sj.FVcTongjrq == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}
	//获取停车场总金额、总笔数
	result := QueryTingccJieSuandata()
	conn := utils.Pool.Get()

	defer func() {
		_ = conn.Close()
	}()
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
		//	conn := utils.RedisConn //初始化redis
		// key:"jiesstatistical"  item: 停车场id  value："金额｜总条数"
		rhseterr := utils.RedisHSet(&conn, "jiesstatistical", r.Parkingid, strconv.Itoa(int(r.Total))+"|"+strconv.Itoa(r.Count))
		if rhseterr != nil {
			return rhseterr
		}
		log.Println("按停车场获取-停车场总金额、总笔数 -【RedisHSet】jiesstatistical 成功 【++++++++++++[1.1]++++++++++++++++++】")
	}
	return nil
}

//1任务4 查询清分、争议处理包
func QueryClearlingAndDisputePackage() error {
	conn := utils.Pool.Get()

	defer func() {
		_ = conn.Close()

	}()

	//1、获取清分包、争议包数据
	Yesterday := utils.Yesterdaydate()
	qcerr, clears := QueryClearlingdata(Yesterday)
	if qcerr != nil {
		return qcerr
	}
	//	Clears:=make( []types.ClearlingAndDispute,0)
	//var Clear types.ClearlingAndDispute
	if clears == nil {
		Clear := types.ClearlingAndDispute{
			DataType:  "clear",
			PackageNo: "",
			DateTime:  "",
		}
		m := make(map[string]string, 0)
		// key:日期    value:"包号"｜"时间"
		m[Yesterday] = Clear.PackageNo + "|" + Clear.DateTime
		//2、把数据存储于redis  接收时间、包号
		hmseterr := utils.RedisHMSet(&conn, Clear.DataType, m)
		if hmseterr != nil {
			return hmseterr
		}
		log.Println("获取清分包-【RedisHSet】 v:=clear 成功 【++++++++++++[1.2]++++++++++++++++++】")
	} else {
		for _, clear := range *clears {
			Clear := types.ClearlingAndDispute{
				DataType:  "clear",
				PackageNo: strconv.Itoa(int(clear.FNbXiaoxxh)),
				DateTime:  clear.FDtChulsj.Format("2006-01-02 15:04:05"),
			}
			// key:日期    value:"包号"｜"时间"
			m1 := make(map[string]string, 0)
			sj := strings.Split(clear.FVcQingfmbr, "T")
			m1[sj[0]] = Clear.PackageNo + "|" + Clear.DateTime
			//2、把数据存储于redis  接收时间、包号
			hmseterr := utils.RedisHMSet(&conn, Clear.DataType, m1)
			if hmseterr != nil {
				return hmseterr
			}
			log.Println("获取清分包-【RedisHSet】 v:=clear 成功 【++++++++++++[1.2]++++++++++++++++++】")

		}
	}

	//1查询争议处理数据
	qderr, dispute := QueryDisputedata(Yesterday)
	if qderr != nil {
		return qderr
	}
	var Disput types.ClearlingAndDispute
	if dispute == nil {
		Disput = types.ClearlingAndDispute{
			DataType:  "disput",
			PackageNo: "",
			DateTime:  "",
		}
	} else {
		Disput = types.ClearlingAndDispute{
			DataType:  "disput",
			PackageNo: strconv.Itoa(int(dispute.FNbXiaoxxh)),
			DateTime:  utils.DateTimeFormat(dispute.FDtZhengyclsj),
		}
	}
	m2 := make(map[string]string, 0)
	//2、把数据存储于redis  接收时间、包号
	m2[Yesterday] = Disput.PackageNo + "|" + Disput.DateTime

	dishmseterr := utils.RedisHMSet(&conn, Disput.DataType, m2)
	if dishmseterr != nil {
		return dishmseterr
	}
	log.Println("获取争议包-【RedisHSet】 v:=disput 成功 【++++++++++++[1.2]++++++++++++++++++】")
	log.Println("  【++++++++++++【1.2 是 ok的】++++++++++++++++++】")

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

	log.Println("更新最新存在争议数据的统计结果成功++++++++++++++++++++【2.3】++++++++++++++++++++++++")
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

	log.Println("插入异常数据统计执行定时任务 成功+++++++++++++++++++++++++++++【2.4】++++++++++++++++++++++")

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
	log.Println("插入最新的黑名单的数据记录成功+++++++++++++++++++++++++【2.13】+++++++++++++++++++++++ ")
	return nil
}

//1.6数据分类
func DataClassification() error {
	//1、插入数据分类记录
	inerr := InsertSWDataClassification()
	if inerr != nil {
		return inerr
	}

	//2、查询数据分类记录
	qerr, dataclassification := QuerySWDataClassificationTable()
	if qerr != nil {
		return qerr
	}

	data := QuerySWDataClassification()

	//4、赋值
	dc := new(types.BJsjkShengwjssjfl)
	dc.FNbJiaoyzts = data.Shengwzcount         //   `F_NB_JIAOYZTS` int DEFAULT NULL COMMENT '交易总条数',
	dc.FNbQingfsjts = data.Yiqfcount           //   `F_NB_QINGFSJTS` int DEFAULT NULL COMMENT '清分数据条数',
	dc.FNbJizsjts = data.Jizcount              //   `F_NB_JIZSJTS` int DEFAULT NULL COMMENT '记账数据条数',
	dc.FNbZhengysjts = data.Zhengycount        //   `F_NB_ZHENGYSJTS` int DEFAULT NULL COMMENT '争议数据条数 待处理',
	dc.FNbWeidbsjts = data.Weidbcount          //   `F_NB_WEIDBSJTS` int DEFAULT NULL COMMENT '未打包数据条数',
	dc.FNbYidbsjts = data.Yidbcount            //   `F_NB_YIDBSJTS` int DEFAULT NULL COMMENT '已打包数据条数',
	dc.FNbYifssjts = data.Yifscount            //   `F_NB_YIFSSJTS` int DEFAULT NULL COMMENT '已发送数据条数',
	dc.FNbHuaizsjts = data.Huaizcount          //   `F_NB_HUAIZSJTS` int DEFAULT NULL COMMENT '坏账数据条数',
	dc.FDtTongjwcsj = utils.StrTimeToNowtime() //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	dc.FVcTongjrq = utils.DateNowFormat()      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	//5、更新
	uperr := UpdateSWDataClassificationTable(dc, dataclassification.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Println("更新 最新的省外结算数据分类 记录 成功+++++++++++++++++++++++[2.7]+++++++++++++++++++++++")

	return nil
}

//3.1转结算监控
func DataTurnMonitor() error {
	//1、新增转结算记录  '统计类型 1:单点、2:总对总',
	inerr := InsertDataTurnMonitor(2)
	if inerr != nil {
		return inerr
	}

	//2、查询转结算数据
	turndata := QueryDataTurnMonitor()

	//3、查询最新结算记录   '统计类型 1:单点、2:总对总',
	qerr, tabledata := QueryDataTurnMonitorTable(2)
	if qerr != nil {
		return qerr
	}
	//4、赋值
	zdzdata := new(types.BJsjkZhuanjssjjk)
	zdzdata.FNbChedyssjts = turndata.ZDZcount       //  `F_NB_CHEDYSSJTS` int DEFAULT NULL COMMENT '车道原始数据条数',
	zdzdata.FNbJiesbsjts = turndata.Jieszcount      //  `F_NB_JIESBSJTS` int DEFAULT NULL COMMENT '结算表数据条数',
	zdzdata.FDtTongjwcsj = utils.StrTimeToNowtime() //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	zdzdata.FVcKuaizsj = utils.DateTimeNowFormat()  //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',

	//5、根据id 更新数据
	uperr := UpdateDataTurnMonitorTable(zdzdata, tabledata.FNbId)
	if uperr != nil {
		return uperr
	}
	//1、新增转结算记录 '统计类型 1:单点、2:总对总',
	ddinerr := InsertDataTurnMonitor(1)
	if ddinerr != nil {
		return ddinerr
	}
	//2、查询最新结算记录   '统计类型 1:单点、2:总对总',
	ddqerr, ddtabledata := QueryDataTurnMonitorTable(1)
	if ddqerr != nil {
		return ddqerr
	}

	//3、赋值
	data := new(types.BJsjkZhuanjssjjk)
	data.FNbChedyssjts = turndata.DDzcount       //  `F_NB_CHEDYSSJTS` int DEFAULT NULL COMMENT '车道原始数据条数',
	data.FNbJiesbsjts = turndata.Jieszcount      //  `F_NB_JIESBSJTS` int DEFAULT NULL COMMENT '结算表数据条数',
	data.FDtTongjwcsj = utils.StrTimeToNowtime() //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	data.FVcKuaizsj = utils.DateTimeNowFormat()  //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',

	//4、根据id 更新数据
	dduperr := UpdateDataTurnMonitorTable(data, ddtabledata.FNbId)
	if dduperr != nil {
		return dduperr
	}
	log.Println("插入之前24小时转结算表数据 记录 完成【按分钟统计】+++++++++++++++++++++++++【2.12】+++++++++++++++++++++++++")

	return nil
}

//1.7 省外结算趋势
func SettlementTrendbyDay() error {
	qsjerr, sj := QuerySettlementTrendbyDayTable()
	if qsjerr != nil {
		return qsjerr
	}
	s1 := utils.DateNowFormat()
	s2 := sj.FDtTongjwcsj.Format("2006-01-02")
	if s2 == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}

	qsdatas := QuerySettlementTrendbyDay()
	for _, qsdata := range *qsdatas {
		//1、新增省外结算趋势
		inerr := InsertSettlementTrendbyDayTable()
		if inerr != nil {
			return inerr
		}
		//2、查询最新一条
		qerr, qsOnedata := QuerySettlementTrendbyDayTable()
		if qerr != nil {
			return qerr
		}
		//3、赋值
		qushijl := new(types.BJsjkShengwjsqs)

		qushijl.FNbJiaoye = qsdata.JiesuanMoney                       //   `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
		qushijl.FNbQingdje = qsdata.ClearlingMoney                    //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
		qushijl.FNbChae = qsdata.JiesuanMoney - qsdata.ClearlingMoney //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
		qushijl.FNbJiaoyts = qsdata.JiesuanCount                      //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
		qushijl.FNbQingfts = qsdata.ClearlingCount                    //   `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
		qushijl.FDtTongjwcsj = utils.StrTimeToNowtime()               //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
		qushijl.FVcTongjrq = qsdata.Datetime                          //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

		//4、更新数据
		uperr := UpdateSettlementTrendbyDayTable(qushijl, qsOnedata.FNbId)
		if uperr != nil {
			return uperr
		}

	}

	log.Println("更新之前30天 省外结算趋势表数据 记录 成功+++++++【1.4】++++++++++")

	return nil
}

//3.1 省外数据包实时监控
func PacketMonitoring() error {
	//1、新增数据包表
	inerr := InsertPacketMonitoringTable()
	if inerr != nil {
		return inerr
	}
	//2、查询最新一次
	qerr, qdata := QueryPacketMonitoringTable()
	if qerr != nil {
		return qerr
	}
	//3、查询数据包数据
	data := QueryPacketMonitoring()
	//4、赋值
	pkgdata := new(types.BJsjkShujbjk)
	pkgdata.FNbDabsl = data.Dabaosl                    //   `F_NB_DABSL` int DEFAULT NULL COMMENT '打包数量',
	pkgdata.FNbDabje = data.Dabaojine                  //   `F_NB_DABJE` bigint DEFAULT NULL COMMENT '打包金额',
	pkgdata.FNbFasysjybsl = data.Fasbsl                //   `F_NB_FASYSJYBSL` int DEFAULT NULL COMMENT '已发送原始交易消息包数量',
	pkgdata.FNbFasysjybje = data.Fasbjine              //   `F_NB_FASYSJYBJE` bigint DEFAULT NULL COMMENT '已发送原始交易消息包金额',
	pkgdata.FNbJizbsl = data.Jizbsl                    //   `F_NB_JIZBSL` int DEFAULT NULL COMMENT '记账包数量',
	pkgdata.FNbJizbje = data.Jizbjine                  //   `F_NB_JIZBJE` bigint DEFAULT NULL COMMENT '记账包金额',
	pkgdata.FNbYuansjyydbsl = data.Yuansbsl            //   `F_NB_YUANSJYYDBSL` int DEFAULT NULL COMMENT '原始交易消息应答包数量',
	pkgdata.FDtTongjwcsj = utils.StrTimeToNowtime()    //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	pkgdata.FVcKuaizsj = utils.KuaizhaoTimeNowFormat() //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',

	//5、更新数据包监控
	uperr := UpdatePacketMonitoringTable(pkgdata, qdata.FNbId)
	if uperr != nil {
		return uperr
	}

	log.Println("插入省外结算数据包表数据 记录 完成++++++++++++++++++++++++++++++【3.1】++++++++++++++++++++++++++++++++++++")

	return nil
}

// 1.8 省内结算总金额、总笔数监控
func ShengnJieSuanData() error {

	//1、新增省内结算监控记录
	inerr := InsertShengnJieSuanTable()
	if inerr != nil {
		return inerr
	}
	//2、查询省内结算监控最新记录
	qerr, data := QueryShengnJieSuanTable()
	if qerr != nil {
		return qerr
	}

	//3、查询省内结算监控数据
	count, amount := QueryShengnJieSuan()
	//4、赋值
	Data := new(types.BJsjkJiestj)
	Data.FNbZongje = amount
	Data.FNbZongts = count
	Data.FDtTongjwcsj = utils.StrTimeToNowtime()
	Data.FVcTongjrq = utils.DateNowFormat()
	//5、更新数据
	uperr := UpdateShengnJieSuanTable(Data, data.FNbId)
	if uperr != nil {
		return uperr
	}

	//6、把数据更新到redis
	//conn := utils.RedisConn //初始化redis
	conn := utils.Pool.Get()

	defer func() {
		_ = conn.Close()
	}()
	// key:"snjiesuantotal"  value："金额｜总条数"
	rseterr := utils.RedisSet(&conn, "snjiesuantotal", strconv.Itoa(int(amount))+"|"+strconv.Itoa(count))
	if rseterr != nil {
		return rseterr
	}
	log.Println("更新省内结算总金额、总笔数到redis 成功+++++++++++++【2.9】+++++++")
	return nil
}

//1.9 省内发送结算数据金额、条数
func ShengnSendJieSuanData() error {

	//1、新增记录
	inerr := InsertShengnSendTable()
	if inerr != nil {
		return inerr
	}
	//2、查询最新记录
	qerr, data := QueryShengnSendTable()
	if qerr != nil {
		return qerr
	}

	//3、查询省内结算数据
	count, amount := QueryShengnSendjiessj()
	//4、赋值
	Data := new(types.BJsjkShengnyfssjtj)
	Data.FNbZongje = amount
	Data.FNbZongts = count
	Data.FDtTongjwcsj = utils.StrTimeToNowtime()
	Data.FVcKuaizsj = utils.KuaizhaoTimeNowFormat()
	//5、更新数据
	uperr := UpdateShengnSendTable(Data, data.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Println("插入省内发送结算数据金额、条数成功+++++++++++++++++++++【2.8】++++++++++")
	return nil
}

//1.10 省内拒付数据金额、条数
func QueryShengnRefusePayData() error {
	//1、新增拒付数据记录
	inerr := InsertShengnRefusePayTable()
	if inerr != nil {
		return inerr
	}
	//2、查询拒付数据最新记录
	qerr, data := QueryShengnRefusePayTable()
	if qerr != nil {
		return qerr
	}

	//3、查询拒付数据数据
	count, amount := QueryShengnRefusePay()
	//4、赋值
	Data := new(types.BJsjkShengnjfsjtj)
	Data.FNbJufzje = amount
	Data.FNbJufzts = count
	Data.FDtTongjwcsj = utils.StrTimeToNowtime()
	Data.FVcTongjrq = utils.DateNowFormat()
	//5、更新拒付数据数据
	uperr := UpdateShengnRefusePayTable(Data, data.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Println("更新拒付数据金额、条数成功+++++++++++++++[2.11]+++++++++++++++++")
	return nil
}

//1.11 查询省内已请款数据金额、条数
func QueryShengnAlreadyPleaseData() error {
	//1、新增省内已请款记录
	inerr := InsertShengnAlreadyPleaseTable()
	if inerr != nil {
		return inerr
	}
	//2、查询省内已请款最新记录
	qerr, data := QueryShengnAlreadyPleaseTable()
	if qerr != nil {
		return qerr
	}

	//3、查询省内已请款数据
	count, amount := QueryAlreadyPlease()
	//4、赋值
	Data := new(types.BJsjkShengnqktj)
	Data.FNbQingkzje = amount
	Data.FNbQingkzts = count
	Data.FDtTongjwcsj = utils.StrTimeToNowtime()
	Data.FVcTongjrq = utils.DateNowFormat()
	//5、更新省内已请款数据
	uperr := UpdateShengnAlreadyPleaseTable(Data, data.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Println("更新省内已请款金额、条数成功+++++++++++++++++++【3.3]+++++++++++++++++++")
	return nil
}

//1.12  查询省内结算分类
func QuerySNDataClassificationData() error {
	//1、新增省内结算分类
	inerr := InsertSNDataClassificationTable()
	if inerr != nil {
		return inerr
	}
	//2、查询省内结算分类
	qerr, data := QuerySNDataClassificationTable()
	if qerr != nil {
		return qerr
	}

	//3、查询省内结算分类
	flshuju := QuerySNDataClassification()
	//4、赋值
	Data := new(types.BJsjkShengnjssjfl)
	Data.FNbShengnzjysl = flshuju.Shengnzcount
	Data.FNbQingksl = flshuju.Yiqkcount
	Data.FNbWeifssl = flshuju.Weifscount
	Data.FNbFassjl = flshuju.Yifscount
	Data.FNbjufsjl = flshuju.Jufuzcount
	Data.FDtTongjwcsj = utils.StrTimeToNowtime()
	Data.FVcTongjrq = utils.DateNowFormat()
	//5、更新省内结算分类
	uperr := UpdateSNDataClassificationTable(Data, data.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Println("更新省内结算分类成功++++++++++++++++++++++[2.10]+++++++++++++++++++")
	return nil
}

//3.2 	省内实时数据
func ShengNRealTimeSettlementData() error {
	//1、新增实时数据表
	inerr := InsertSNRealTimeSettlementDataTable()
	if inerr != nil {
		return inerr
	}
	//2、查询实时数据最新一次
	qerr, xdata := QuerySNRealTimeSettlementDataTable()
	if qerr != nil {
		return qerr
	}
	//get redis
	//conn := utils.RedisConn //初始化redis
	conn := utils.Pool.Get()

	defer func() {
		_ = conn.Close()

	}()
	// key:"snshishishuju"  value："金额｜总条数"
	rhgeterr, value := utils.RedisGet(&conn, "snshishishuju")
	if rhgeterr != nil {
		return rhgeterr
	}
	if value == nil {
		log.Println("查询 获取省内实时数据为空 ")
		return errors.New("get redis value==nil")
	}

	vstr := string(value.([]uint8))
	log.Println("The get redis value is ", vstr)

	if !utils.StringExist(vstr, "|") {
		return errors.New("get redis error")
	}

	vs := strings.Split(vstr, `"`)
	v := strings.Split(vs[1], `|`)
	zje, _ := strconv.Atoi(v[0])    //省内产生金额
	zts, _ := strconv.Atoi(v[1])    //省内产生条数
	fsjine, _ := strconv.Atoi(v[2]) //'省内已发送数据金额'
	fszts, _ := strconv.Atoi(v[3])  //省内已发送数据条数
	jzjine, _ := strconv.Atoi(v[4]) //省内已记账数据金额
	jzzts, _ := strconv.Atoi(v[5])  //省内已记账数据条数
	log.Println("获取省内实时成功+++++++++：", "省内产生金额: ", int64(zje), "省内产生条数:", zts, "省内已发送数据金额:", fsjine, "省内已发送数据条数:", fszts, "省内已记账数据金额:", jzjine, "省内已记账数据条数:", jzzts)

	//3、查询实时数据
	data := QueryRealTimeSettlementData()

	//4、赋值
	ssdata := new(types.BJsjkShengnsssjjk)
	ssdata.FNbShengncsje = data.Shengnjsjine - int64(zje)  //   `F_NB_SHENGNCSJE` bigint DEFAULT NULL COMMENT '省内产生金额',
	ssdata.FNbShengnyfssjje = data.Fasjine - int64(fsjine) //   `F_NB_SHENGNYFSSJJE` bigint DEFAULT NULL COMMENT '省内已发送数据金额',
	ssdata.FNbShengnyjzsjje = data.Jizjine - int64(jzjine) //   `F_NB_SHENGNYJZSJJE` bigint DEFAULT NULL COMMENT '省内已记账数据金额',
	ssdata.FNbShengncsts = data.Shengnjssl - zts           //   `F_NB_SHENGNCSTS` int DEFAULT NULL COMMENT '省内产生条数',
	ssdata.FNbShengnyfssjts = data.Fassl - fszts           //   `F_NB_SHENGNYFSSJTS` int DEFAULT NULL COMMENT '省内已发送数据条数',
	ssdata.FNbShengnyjzsjts = data.Jizsl - jzzts           //   `F_NB_SHENGNYJZSJTS` int DEFAULT NULL COMMENT '省内已记账数据条数',
	ssdata.FDtTongjwcsj = utils.StrTimeToNowtime()         //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	ssdata.FVcTongjrq = utils.KuaizhaoTimeNowFormat()      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	//5、更新实时数据监控
	uperr := UpdateSNRealTimeSettlementDataTable(ssdata, xdata.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Println("插入省内今日实时数据 完成++++++++++++++++++++++【3.2】++++++++++++++++++++++++")

	//redis set新值
	s := strconv.Itoa(int(data.Shengnjsjine)) + "|" + strconv.Itoa(data.Shengnjssl) + "|" + strconv.Itoa(int(data.Fasjine)) + "|" + strconv.Itoa(data.Fassl) + "|" + strconv.Itoa(int(data.Jizjine)) + "|" + strconv.Itoa(data.Jizsl)
	rseterr := utils.RedisSet(&conn, "snshishishuju", s)
	if rseterr != nil {
		return rseterr
	}
	log.Println("set redis 成功 ")
	log.Println("set redis 省内今日实时数据 完成 [ snshishishuju ]  ++++++++++++++++++++++【3.2】++++++++++++++++++++++++")

	return nil
}

//1.13 省内结算趋势
func QueryShengNSettlementTrenddata() error {
	qsjerr, sj := QueryShengNSettlementTrendTable()
	if qsjerr != nil {
		return qsjerr
	}
	s1 := utils.DateNowFormat()
	s2 := sj.FDtTongjwcsj.Format("2006-01-02")
	if s2 == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}

	// 查询省内结算趋势
	qsshujus := QueryShengNSettlementTrend()
	for i, qsshuju := range *qsshujus {
		//1、新增省内结算趋势
		inerr := InsertShengNSettlementTrendTable()
		if inerr != nil {
			return inerr
		}
		//2、查询省内结算趋势
		qerr, data := QueryShengNSettlementTrendTable()
		if qerr != nil {
			return qerr
		}
		//4、赋值
		Data := new(types.BJsjkShengnjsqs)
		Data.FNbShengnjyje = qsshuju.JiesuanMoney                    //   `F_NB_SHENGNJYJE` bigint DEFAULT NULL COMMENT '省内交易金额',
		Data.FNbShengnqkje = qsshuju.ClearlingMoney                  //   `F_NB_SHENGNQKJE` bigint DEFAULT NULL COMMENT '省内请款金额',
		Data.FNbChae = qsshuju.JiesuanMoney - qsshuju.ClearlingMoney //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
		Data.FNbJiaoyts = qsshuju.JiesuanCount                       //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
		Data.FNbQingkts = qsshuju.ClearlingCount                     //   `F_NB_QINGKTS` int DEFAULT NULL COMMENT '请款条数',
		Data.FVcKuaizsj = qsshuju.Datetime                           //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
		Data.FDtTongjwcsj = utils.StrTimeToNowtime()
		//5、更新省内结算趋势
		uperr := UpdateShengNSettlementTrendTable(Data, data.FNbId)
		if uperr != nil {
			return uperr
		}
		log.Printf("更新第%d天省内结算趋势成功+++++++++++++", i+1)
	}
	log.Println("更新省内之前30天结算趋势成功+++++++++++++++++++++++++【1.5】+++++++++++++++++++++++++")

	return nil
}

//2.4 海玲数据同步
func QueryDataSyncdata() error {
	// 查询海玲数据同步
	hlnum, tongbcount := QueryDataSync()

	//1、新增海玲数据同步
	inerr := InsertDataSyncTable()
	if inerr != nil {
		return inerr
	}
	//2、查询最新一条海玲数据同步
	qerr, data := QueryDataSyncTable()
	if qerr != nil {
		return qerr
	}
	//4、赋值
	Data := new(types.BJsjkShujtbjk)
	Data.FNbJiessjzl = hlnum                     //   `F_NB_JIESJZL` int DEFAULT NULL COMMENT '结算数据总量',F_NB_JIESSJZL
	Data.FNbYitbsjl = tongbcount                 //   `F_NB_YITBSJL` int DEFAULT NULL COMMENT '已同步数据量',
	Data.FDtTongjwcsj = utils.StrTimeToNowtime() //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	Data.FVcTongjrq = utils.DateNowFormat()      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

	//5、更新最新一条海玲数据同步记录
	uperr := UpdateDataSyncTable(Data, data.FNbId)
	if uperr != nil {
		return uperr
	}
	log.Printf("更新最新一条海玲数据同步成功")
	log.Println("插入 oracle 海玲数据同步定时任务 成功+++++++++++++++++++++++++++【2.5】+++++++++++++++++++++++++++++++")
	return nil
}

//AbnormalDataOfParking
//2.5 异常数据top10
func QueryAbnormalDataOfParkingdata() error {
	// 查询异常数据top10
	dddata, zdzdata := QueryAbnormalDataOfParking()

	for _, dd := range *dddata {
		//1、新增异常数据top10
		inerr := InsertAbnormalDataOfParkingTable(1)
		if inerr != nil {
			return inerr
		}
		//2、查询最新一条异常数据top10
		qerr, data := QueryAbnormalDataOfParkingTable()
		if qerr != nil {
			return qerr
		}
		//4、赋值
		Data := new(types.BJsjkYicsjtcctj)

		Data.FNbZongts = dd.Count                       //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
		Data.FNbZongje = dd.Total                       //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额（分）',
		Data.FVcTingccid = dd.Parkingid                 //   `F_NB_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
		Data.FVcKuaizsj = utils.KuaizhaoTimeNowFormat() //   `F_VC_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
		Data.FDtTongjwcsj = utils.StrTimeToNowtime()    //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',

		//5、更新最新一条异常数据top10
		uperr := UpdateAbnormalDataOfParkingTable(Data, data.FNbId)
		if uperr != nil {
			return uperr
		}
		log.Printf("更新单点最新一条异常数据top10成功")
	}

	for _, zdz := range *zdzdata {
		//1、新增异常数据top10
		inerr := InsertAbnormalDataOfParkingTable(1)
		if inerr != nil {
			return inerr
		}
		//2、查询最新一条异常数据top10
		qerr, zdzdata := QueryAbnormalDataOfParkingTable()
		if qerr != nil {
			return qerr
		}
		//4、赋值
		Data := new(types.BJsjkYicsjtcctj)

		Data.FNbZongts = zdz.Count                      //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
		Data.FNbZongje = zdz.Total                      //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额（分）',
		Data.FVcTingccid = zdz.Parkingid                //   `F_NB_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
		Data.FVcKuaizsj = utils.KuaizhaoTimeNowFormat() //   `F_VC_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
		Data.FDtTongjwcsj = utils.StrTimeToNowtime()    //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',

		//5、更新最新一条异常数据top10
		uperr := UpdateAbnormalDataOfParkingTable(Data, zdzdata.FNbId)
		if uperr != nil {
			return uperr
		}
		log.Printf("更新异常数据top10成功+++++++++++++++ ")
	}
	log.Printf("插入异常数据top10成功++++++++++++++++++【2.6】++++++++++++++++")

	return nil
}

// QueryOverdueData
//1.14 逾期数据
func Overduedata() error {
	//逾期数据
	shujus := QueryOverdueData()
	if len(*shujus) == 0 {
		log.Printf("查询逾期数据成功，没有逾期数据+++++++++++++++++++++++++++[1.6]++++++++++++++++++++++++++++")
		return nil
	}
	for _, yqdata := range *shujus {
		//1、新增逾期数据
		inerr := InsertOverdueDataTable()
		if inerr != nil {
			return inerr
		}
		//2、查询逾期数据
		qerr, data := QueryOverdueDataTable()
		if qerr != nil {
			return qerr
		}
		//4、赋值
		Data := new(types.BJsjkYuqsjtj)
		Data.FNbYuqzts = yqdata.Count           //   `F_NB_YUQZTS` int DEFAULT NULL COMMENT '逾期总条数',
		Data.FNbYuqzje = yqdata.Total           //   `F_NB_YUQZJE` bigint DEFAULT NULL COMMENT '逾期总金额 （分）',
		Data.FVcTingccid = yqdata.Parkingid     //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
		Data.FVcTongjrq = utils.DateNowFormat() //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
		Data.FDtTongjwcsj = utils.StrTimeToNowtime()
		//5、更新逾期数据
		uperr := UpdateOverdueDataTable(Data, data.FNbId)
		if uperr != nil {
			return uperr
		}
		log.Printf("插入逾期数据成功+++++++++++")
	}
	log.Printf("更新逾期数据成功+++++++++++++++++++++++++++[1.6]++++++++++++++++++++++++++++")
	return nil
}

//1.15 省外停车场结算趋势
func SWSettlementTrendbyDay() error {
	qoneerr, qsOnesqsj := QuerySWSettlementTrendTable()
	if qoneerr != nil {
		return qoneerr
	}

	s1 := utils.DateNowFormat()
	s2 := qsOnesqsj.FDtTongjwcsj.Format("2006-01-02")
	if s2 == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}

	qsdatas := QuerySWSettlementTrendOne()
	for _, qsdata := range *qsdatas {
		//1、新增省外结算趋势
		inerr := InsertSWSettlementTrendTable()
		if inerr != nil {
			return inerr
		}
		//2、查询最新一条
		qerr, qsOnedata := QuerySWSettlementTrendTable()
		if qerr != nil {
			return qerr
		}
		//3、赋值
		qushijl := new(types.BJsjkShengwtccjsqs)
		qushijl.FVcTingccid = qsdata.Parkingid                        //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
		qushijl.FNbJiaoyje = qsdata.JiesuanMoney                      //   `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
		qushijl.FNbQingfje = qsdata.ClearlingMoney                    //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
		qushijl.FNbChae = qsdata.JiesuanMoney - qsdata.ClearlingMoney //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
		qushijl.FNbJiaoyts = qsdata.JiesuanCount                      //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
		qushijl.FNbQingfts = qsdata.ClearlingCount                    //   `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
		qushijl.FDtTongjwcsj = utils.StrTimeToNowtime()               //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
		qushijl.FVcTongjrq = qsdata.Datetime                          //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

		//4、更新数据
		uperr := UpdateSWSettlementTrendTable(qushijl, qsOnedata.FNbId)
		if uperr != nil {
			return uperr
		}
	}
	log.Printf("更新省外停车场结算趋势+++++++++++++++++++++++++++[1.7]++++++++++++++++++++++++++++")
	return nil
}

//1.16 省内停车场昨日结算趋势
func SNSettlementTrendbyDay() error {
	qsnqserr, snqsonedata := QuerySNSettlementTrendTable()
	if qsnqserr != nil {
		return qsnqserr
	}
	s1 := utils.DateNowFormat()
	s2 := snqsonedata.FDtTongjwcsj.Format("2006-01-02")
	if s2 == s1 {
		log.Println("这一天已经插入数据了，不需要重复统计")
		return nil
	}

	qsdatas := QuerySNSettlementTrendOne()
	for _, qsdata := range *qsdatas {
		//1、新增省内结算趋势
		inerr := InsertSNSettlementTrendTable()
		if inerr != nil {
			return inerr
		}
		//2、查询最新一条
		qerr, qsOnedata := QuerySNSettlementTrendTable()
		if qerr != nil {
			return qerr
		}
		//3、赋值
		qushijl := new(types.BJsjkShengntccjsqs)

		qushijl.FVcTingccid = qsdata.Parkingid                        //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
		qushijl.FNbShengnjyje = qsdata.JiesuanMoney                   //   `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
		qushijl.FNbShengnqkje = qsdata.ClearlingMoney                 //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
		qushijl.FNbChae = qsdata.JiesuanMoney - qsdata.ClearlingMoney //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
		qushijl.FNbJiaoyts = qsdata.JiesuanCount                      //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
		qushijl.FNbQingkts = qsdata.ClearlingCount                    //   `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
		qushijl.FDtTongjwcsj = utils.StrTimeToNowtime()               //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
		qushijl.FVcKuaizsj = qsdata.Datetime                          //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

		//4、更新数据
		uperr := UpdateSNSettlementTrendTable(qushijl, qsOnedata.FNbId)
		if uperr != nil {
			return uperr
		}
	}

	log.Printf("更新省内停车场结算趋势ok+++++++++++++++++++++++++++【1.8】++++++++++++++++++++++++++++")
	return nil
}

//
