package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"strconv"
	"strings"
)

//查询省外结算总金额、总笔数
func QuerTotalSettlementData() (int, error, *dto.TotalSettlementData) {
	//查询数据库获取总金额、总笔数
	conn := utils.RedisInit() //初始化redis
	// key:"jiestotal"  value："金额｜总条数"
	rhgeterr, value := utils.RedisGet(conn, "jiesuantotal")
	if rhgeterr != nil {
		return 0, rhgeterr, nil
	}
	if value == nil {
		log.Println("查询数据库获取总金额、总笔数为空 ")
		return 0, errors.New("get redis value==nil"), nil
	}

	vstr := string(value.([]uint8))
	log.Println("The get redis value is ", vstr)

	if !utils.StringExist(vstr, "|") {
		return 0, errors.New("get redis error"), nil
	}

	vs := strings.Split(vstr, `"`)

	v := strings.Split(vs[1], `|`)
	zje, _ := strconv.Atoi(v[0])
	zts, _ := strconv.Atoi(v[1])
	log.Println("查询成功", "jine: ", int64(zje), "Count", zts)
	//返回数据赋值
	return 203, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(int64(zje)), Count: zts}
}

//查询省外已清分总金额、总笔数
func QuerTotalClarify() (int, error, *dto.TotalClarifyData) {
	//包含坏账的已清分
	qerr, qingfjg := db.QueryShengwClearingdata()
	if qerr != nil {
		log.Println("查询省外已清分总金额、总笔数,查询最新数据时  error!", qerr)
		return 0, qerr, nil //不用返回前端
	}
	log.Println("查询省外已清分总金额、总笔数 (包含坏账的)成功")
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
	log.Println("查询省外坏账总金额、总笔数 成功")
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
	log.Println("查询存在争议的数据总金额、总笔数 (包含坏账的)成功")
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

	log.Println("查询异常的数据总金额、总笔数  成功")
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
	ts := 3 //需要查询条数【后面可以改】
	qdterr, hmdjls := db.QueryBlacklistTiaoshutable(id, ts)

	if qdterr != nil {
		log.Println(qdterr)
		return 0, qdterr, nil
	}

	changecount := (*hmdjls)[2].FNbHeimdzs - (*hmdjls)[0].FNbHeimdzs
	log.Println("查询黑名单总数、较2小时前变化值  成功", (*hmdjls)[2].FNbHeimdzs, changecount)
	//返回数据赋值
	return 208, nil, &dto.TotalBlacklistData{Blacklistcount: (*hmdjls)[2].FNbHeimdzs, ChangeCount: changecount}
}

//查询清分包、争议包的接收时间、包号 14天
func QueryClearlingAndDisputePackagedata() (int, error, *dto.ClearlAndDisputeData) {

	//查询清分包、争议包的接收时间、包号[最新的数据]前14天数据

	chmgeterr, cleardata := utils.RedisHMGet(utils.RedisInit(), "clear", utils.OldData(14))
	if chmgeterr != nil {
		return 0, chmgeterr, nil
	}

	dhmgeterr, disputdata := utils.RedisHMGet(utils.RedisInit(), "disput", utils.OldData(14))
	if dhmgeterr != nil {
		return 0, dhmgeterr, nil
	}
	data := make([]types.ClearlingAndDisputeData, 0)

	for _, clearData := range *cleardata {
		clearAndDis := new(types.ClearlingAndDisputeData)
		if clearData == "no data" {
			continue
		}
		str := strings.Split(clearData, "|")
		clearAndDis.ClearPacgNo = str[0]
		clearAndDis.Cleardatetime = str[1]
		data = append(data, *clearAndDis)
	}

	for i, disputData := range *disputdata {
		if disputData == "no data" {
			continue
		}
		disstr := strings.Split(disputData, "|")
		data[i].Disputdatetime = disstr[1]
		data[i].DisputPacgeNo = disstr[0]
	}

	log.Println("查询清分包、争议包的接收时间、包号  成功。data数组长度:", len(data))
	//返回数据赋值
	return 209, nil, &dto.ClearlAndDisputeData{data}
}
