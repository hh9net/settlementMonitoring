package service

import (
	"github.com/sirupsen/logrus"
	"log"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/utils"
)

//查询省外结算总金额、总笔数
func QuerTotalSettlementData() (int, error, *dto.TotalSettlementData) {
	qerr, sj := db.QueryTabledata(10000)
	if qerr != nil {
		logrus.Println("查询省外结算总金额、总笔数,查询最新的省外结算统计记录  error!", qerr)
		return 0, qerr, nil //不用返回前端
	}
	logrus.Println("查询成功")
	//返回数据赋值
	return 203, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(sj.FNbZongje), Count: sj.FNbZongts}
}

//查询省外已清分总金额、总笔数
func QuerTotalClarify() (int, error, *dto.TotalClarifyData) {
	//包含坏账的已清分
	qerr, qingfjg := db.QueryShengwClearingdata()
	if qerr != nil {
		log.Println("查询省外已清分总金额、总笔数,查询最新数据时  error!", qerr)
		return 0, qerr, nil //不用返回前端
	}
	logrus.Println("查询省外已清分总金额、总笔数 (包含坏账的)成功")
	//返回数据赋值
	return 204, nil, &dto.TotalClarifyData{Amount: utils.Fen2Yuan(qingfjg.FNbZongje), Count: qingfjg.FNbZongts}
}

//查询省外坏账总金额、总笔数
func QuerTotalBaddebts() (int, error, *dto.TotalBaddebtsData) {
	//坏账
	qerr, qingfjg := db.QueryShengwClearingdata()
	if qerr != nil {
		log.Println("查询省外坏账总金额、总笔数,查询最新数据时  error!", qerr)
		return 0, qerr, nil //不用返回前端
	}
	logrus.Println("查询省外坏账总金额、总笔数 成功")
	//返回数据赋值
	return 205, nil, &dto.TotalBaddebtsData{Amount: utils.Fen2Yuan(qingfjg.FNbHuaizje), Count: qingfjg.FNbHuaizts}
}

//查询存在争议的数据总金额、总笔数
func QueryDisputedata() (int, error, *dto.TotalDisputeData) {
	//存在争议的数据
	qerr, zyjg := db.QueryShengwDispute()
	if qerr != nil {
		log.Println("查询存在争议的数据总金额、总笔数,查询最新数据时  error!", qerr)
		return 0, qerr, nil //不用返回前端
	}
	logrus.Println("查询存在争议的数据总金额、总笔数 (包含坏账的)成功")
	//返回数据赋值
	return 206, nil, &dto.TotalDisputeData{Amount: utils.Fen2Yuan(zyjg.FNbZongje), Count: zyjg.FNbZongts}
}

//查询异常的数据总金额、总笔数
func QueryAbnormaldata() (int, error, *dto.TotalAbnormalData) {

	//查询异常数据的统计结果[最新的数据]
	yccount, ycamount, qycerr := db.QueryAbnormalData(1)
	if qycerr != nil {
		log.Println(qycerr)
		return 0, qycerr, nil
	}
	ddyccount, ddycamount, ddqycerr := db.QueryAbnormalData(2)
	if ddqycerr != nil {
		log.Println(ddqycerr)
		return 0, ddqycerr, nil
	}

	logrus.Println("查询异常的数据总金额、总笔数  成功")
	//返回数据赋值
	return 207, nil, &dto.TotalAbnormalData{Amount: utils.Fen2Yuan(ycamount + ddycamount), Count: yccount + ddyccount}
}

//Queryblacklistdata
//查询异常的数据总金额、总笔数
func Queryblacklistdata() (int, error, *dto.TotalBlacklistData) {

	//查询黑名单总数、较2小时前变化值[最新的数据]
	qerr, hmdjl := db.QueryBlacklisttable()

	if qerr != nil {
		log.Println(qerr)
		return 0, qerr, nil
	}
	id := hmdjl.FNbId
	if hmdjl.FNbHeimdzs == 0 {
		id = id - 1
	}
	ts := 3 //需要查询条数
	qdterr, hmdjls := db.QueryBlacklistTiaoshutable(id, ts)

	if qdterr != nil {
		log.Println(qdterr)
		return 0, qdterr, nil
	}

	changecount := (*hmdjls)[2].FNbHeimdzs - (*hmdjls)[0].FNbHeimdzs
	logrus.Println("查询黑名单总数、较2小时前变化值  成功", (*hmdjls)[2].FNbHeimdzs, changecount)
	//返回数据赋值
	return 208, nil, &dto.TotalBlacklistData{Blacklistcount: (*hmdjls)[2].FNbHeimdzs, ChangeCount: changecount}
}
