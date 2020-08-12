package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/utils"
	"strconv"
	"strings"
)

//查询省内结算总金额、总笔数
func QuerSNTotalSettlementData() (int, error, *dto.TotalSettlementData) {
	//查询数据库获取总金额、总笔数
	conn := utils.RedisInit() //初始化redis
	// key:"jiestotal"  value："金额｜总条数"
	rhgeterr, value := utils.RedisGet(conn, "snjiesuantotal")
	if rhgeterr != nil {
		return 0, rhgeterr, nil
	}
	if value == nil {
		log.Println("查询redis数据库获取省内结算总金额、总笔数为空 ")
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
	log.Println("查询成功", "省内结算总金额: ", int64(zje), "省内结算总条数", zts)
	//返回数据赋值
	return 301, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(int64(zje)), Count: zts}
}

//查询省内的已发送 总条数、总金额
func QuerySNSendTotalSettlemen() (int, error, *dto.TotalSettlementData) {
	//查询省内的已发送 总条数、总金额
	err, data := db.QueryShengnSendTable()
	if err != nil {
		return 0, err, nil
	}
	if data.FVcKuaizsj == "" {
		err2, Data := db.QueryShengnSendTableByID(data.FNbId - 1)
		if err2 != nil {
			return 0, err, nil
		}
		log.Println("查询成功", "省内已发送结算总金额: ", Data.FNbZongje, "省内已发送结算总条数", Data.FNbZongts)
		return 302, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(Data.FNbZongje), Count: Data.FNbZongts}
	}
	//返回数据赋值
	return 302, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(data.FNbZongje), Count: data.FNbZongts}
}
