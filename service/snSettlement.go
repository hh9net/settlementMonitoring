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

//QuerySNAlreadyPleaseData
func QuerySNAlreadyPleaseData() (int, error, *dto.TotalSettlementData) {
	//查询省内已请款的数据总条数、总金额
	err, data := db.QueryShengnAlreadyPleaseTable()
	if err != nil {
		return 0, err, nil
	}
	if data.FVcTongjrq == "" {
		err2, Data := db.QueryShengnAlreadyPleaseTableByID(data.FNbId - 1)
		if err2 != nil {
			return 0, err, nil
		}
		log.Println("查询成功", "查询省内已请款的数据总金额: ", Data.FNbQingkzje, "省内已请款的数据总条数", Data.FNbQingkzts)
		return 303, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(Data.FNbQingkzje), Count: Data.FNbQingkzts}
	}
	//返回数据赋值
	return 303, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(data.FNbQingkzje), Count: data.FNbQingkzts}
}

//QuerySNRefusePayData
func QuerySNRefusePayData() (int, error, *dto.TotalSettlementData) {
	//查询坏账（拒付）数据 总条数、总金额
	err, data := db.QueryShengnRefusePayTable()
	if err != nil {
		return 0, err, nil
	}
	if data.FVcTongjrq == "" {
		err2, Data := db.QueryShengnRefusePayTableByID(data.FNbId - 1)
		if err2 != nil {
			return 0, err, nil
		}
		log.Println("查询成功", "查询省内已请款的数据总金额: ", Data.FNbJufzje, "省内已请款的数据总条数", Data.FNbJufzts)
		return 304, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(Data.FNbJufzje), Count: Data.FNbJufzts}
	}
	//返回数据赋值
	return 304, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(data.FNbJufzje), Count: data.FNbJufzts}
}

//QuerySNRealTimeData
func QuerySNRealTimeData() (int, error, *[]dto.RealTimeSettlementData) {
	//查询省内结算实时数据监控
	ts := 30
	Data := make([]dto.RealTimeSettlementData, ts)

	err, ds := db.QuerySNRealTimeSettlementData(ts)
	if err != nil {
		return 0, err, nil
	}
	for i, d := range *ds {
		Data[i].Jizjine = d.FNbShengncsje //省内产生金额
		Data[i].Fasjine = d.FNbShengnyfssjje
		Data[i].Jizjine = d.FNbShengnyjzsjje
		Data[i].Shengnjssl = d.FNbShengncsts
		Data[i].Fassl = d.FNbShengnyfssjts
		Data[i].Jizsl = d.FNbShengnyjzsjts
	}
	//返回数据赋值
	return 305, nil, &Data
}

//QuerySNSettlementTrend
func QuerySNSettlementTrend() (int, error, *[]dto.SNClearandJiesuan) {
	//查询前30日省内结算趋势概览
	ts := 30
	Data := make([]dto.SNClearandJiesuan, ts)
	err, ds := db.QueryShengNSettlementTrendtable(ts)
	if err != nil {
		return 0, err, nil
	}
	for i, d := range *ds {
		Data[i].JiesuanMoney = d.FNbShengnjyje //省内产生金额
		Data[i].ClearlingMoney = d.FNbShengnqkje
		Data[i].DiffMoney = d.FNbChae
		Data[i].JiesuanCount = d.FNbJiaoyts
		Data[i].ClearlingCount = d.FNbQingkts
	}
	//返回数据赋值
	return 306, nil, &Data
}

//QueryDataSync
func QueryDataSync() (int, error, *[]dto.DataSync) {
	//查询海岭数据同步监控
	ts := 12
	Data := make([]dto.DataSync, ts)
	err, ds := db.QueryDataSynctable(ts)
	if err != nil {
		return 0, err, nil
	}
	for i, d := range *ds {
		Data[i].HailCount = d.FNbJiessjzl
		Data[i].JiesuanCount = d.FNbYitbsjl
	}
	//返回数据赋值
	return 307, nil, &Data
}

//QuerySNDataClassification
func QuerySNDataClassification() (int, error, *dto.ShengNDataClassification) {
	//查询省内结算数据分类
	err, data := db.QuerySNDataClassificationTable()
	if err != nil {
		return 0, err, nil
	}
	if data.FVcTongjrq == "" {
		err, Data := db.QuerySNDataClassificationTableByID(data.FNbId - 1)
		if err != nil {
			return 0, err, nil
		}
		return 308, nil, &dto.ShengNDataClassification{
			Data.FNbShengnzjysl,
			Data.FNbQingksl,
			Data.FNbWeifssl,
			Data.FNbFassjl,
			Data.FNbjufsjl,
		}
	}
	//返回数据赋值
	return 308, nil, &dto.ShengNDataClassification{
		data.FNbShengnzjysl,
		data.FNbQingksl,
		data.FNbWeifssl,
		data.FNbFassjl,
		data.FNbjufsjl,
	}
}

//QueryAbnormalDataParking
func QueryAbnormalDataParking() (int, error, *[]dto.AbnormalDataOfParking) {
	//查询异常数据停车场top10
	ts := 10
	Data := make([]dto.AbnormalDataOfParking, ts)
	err, data := db.QueryAbnormalDataOfParkingTable()
	if err != nil {
		return 0, err, nil
	}
	err1, yqdatas := db.QueryAbnormalDataOfParkingtable(data.FVcKuaizsj, ts)
	if err1 != nil {
		return 0, err1, nil
	}
	for i, yqd := range *yqdatas {
		Data[i].AbnormalDatacount = yqd.FNbZongts
		Data[i].Parkingname = yqd.FVcTingccid
		Data[i].AbnormalDataAmount = yqd.FNbZongje
	}
	//返回数据赋值
	return 309, nil, &Data
}

//QueryOverdueData
func QueryOverdueData() (int, error, *[]dto.Overduedata) {
	//查询海岭数据同步监控
	ts := 10
	Data := make([]dto.Overduedata, ts)
	err, data := db.QueryOverdueDataTable()
	if err != nil {
		return 0, err, nil
	}
	err1, yqdatas := db.QueryOverdueDatatable(data.FVcTongjrq, ts)
	if err1 != nil {
		return 0, err1, nil
	}
	for i, yqd := range *yqdatas {
		Data[i].Overduecount = yqd.FNbYuqzts
		Data[i].Parkingname = yqd.FVcTingccid
		Data[i].OverdueAmount = yqd.FNbYuqzje
	}
	//返回数据赋值
	return 310, nil, &Data
}
