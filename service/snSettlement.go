package service

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/db"
	"settlementMonitoring/dto"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
	"strconv"
	"strings"
)

//查询省内结算总金额、总笔数
func QuerSNTotalSettlementData() (int, error, *dto.TotalSettlementData) {
	//查询数据库获取总金额、总笔数
	//conn := utils.RedisConn //初始化redis
	conn := utils.Pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	// key:"jiestotal"  value："金额｜总条数"
	rhgeterr, value := utils.RedisGet(&conn, "snjiesuantotal")
	if rhgeterr != nil {
		return types.Statuszero, rhgeterr, nil
	}
	if value == nil {
		log.Println("查询redis数据库获取省内结算总金额、总笔数为空 ")
		return types.Statuszero, errors.New("get redis value==nil"), nil
	}

	vstr := string(value.([]uint8))
	log.Println("The get redis value is ", vstr)

	if !utils.StringExist(vstr, "|") {
		return types.Statuszero, errors.New("get redis error"), nil
	}

	vs := strings.Split(vstr, `"`)

	v := strings.Split(vs[1], `|`)
	zje, _ := strconv.Atoi(v[0])
	zts, _ := strconv.Atoi(v[1])
	log.Println("++++++++++++++++++++++++++【实时】查询成功", "【实时】省内结算总金额: ", int64(zje), "省内结算总条数", zts)
	//返回数据赋值
	return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(int64(zje)), Count: zts}
}

//查询省内的已发送 总条数、总金额
func QuerySNSendTotalSettlemen() (int, error, *dto.TotalSettlementData) {
	//查询省内的已发送 总条数、总金额
	err, data := db.QueryShengnSendTable()
	if err != nil {
		return types.Statuszero, err, nil
	}
	if data.FVcKuaizsj == "" {
		err2, Data := db.QueryShengnSendTableByID(data.FNbId - 1)
		if err2 != nil {
			return types.Statuszero, err, nil
		}
		log.Println("查询成功", "省内已发送结算总金额: ", Data.FNbZongje, "省内已发送结算总条数", Data.FNbZongts)
		return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(Data.FNbZongje), Count: Data.FNbZongts}
	}
	//返回数据赋值
	log.Println("查询成功", "省内已发送结算总金额: ", data.FNbZongje, "省内已发送结算总条数", data.FNbZongts)

	return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(data.FNbZongje), Count: data.FNbZongts}
}

//QuerySNAlreadyPleaseData
func QuerySNAlreadyPleaseData() (int, error, *dto.TotalSettlementData) {
	//查询省内已请款的数据总条数、总金额
	err, data := db.QueryShengnAlreadyPleaseTable()
	if err != nil {
		return types.Statuszero, err, nil
	}
	if data.FVcTongjrq == "" {
		err2, Data := db.QueryShengnAlreadyPleaseTableByID(data.FNbId - 1)
		if err2 != nil {
			return types.Statuszero, err, nil
		}
		log.Println("查询成功", "查询省内已请款的数据总金额: ", Data.FNbQingkzje, "省内已请款的数据总条数", Data.FNbQingkzts)
		return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(Data.FNbQingkzje), Count: Data.FNbQingkzts}
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(data.FNbQingkzje), Count: data.FNbQingkzts}
}

//QuerySNRefusePayData
func QuerySNRefusePayData() (int, error, *dto.TotalSettlementData) {
	//查询坏账（拒付）数据 总条数、总金额
	err, data := db.QueryShengnRefusePayTable()
	if err != nil {
		return types.Statuszero, err, nil
	}
	if data.FVcTongjrq == "" {
		err2, Data := db.QueryShengnRefusePayTableByID(data.FNbId - 1)
		if err2 != nil {
			return types.Statuszero, err, nil
		}
		log.Println("查询成功", "查询省内已请款的数据总金额: ", Data.FNbJufzje, "省内已请款的数据总条数", Data.FNbJufzts)
		return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(Data.FNbJufzje), Count: Data.FNbJufzts}
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &dto.TotalSettlementData{Amount: utils.Fen2Yuan(data.FNbJufzje), Count: data.FNbJufzts}
}

//QuerySNRealTimeData
func QuerySNRealTimeData() (int, error, *[]dto.RealTimeSettlementData) {
	//查询省内结算实时数据监控 应该 144 条
	ts := types.Frequency
	Data := make([]dto.RealTimeSettlementData, ts)
	err, ds := db.QuerySNRealTimeSettlementData(ts)
	if err != nil {
		return types.Statuszero, err, nil
	}
	for i, d := range *ds {
		Data[i].Jizjine = utils.Fen2Yuan(d.FNbShengncsje) //省内产生金额
		Data[i].Fasjine = utils.Fen2Yuan(d.FNbShengnyfssjje)
		Data[i].Jizjine = utils.Fen2Yuan(d.FNbShengnyjzsjje)
		Data[i].Shengnjssl = d.FNbShengncsts
		Data[i].Fassl = d.FNbShengnyfssjts
		Data[i].Jizsl = d.FNbShengnyjzsjts
		Data[i].DateTime = d.FDtTongjwcsj.Format("2006-01-02 15:04:05")
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &Data
}

//QuerySNSettlementTrend
func QuerySNSettlementTrend() (int, error, *[]dto.SNClearandJiesuan) {
	//查询前30日省内结算趋势概览
	ts := 30
	Data := make([]dto.SNClearandJiesuan, ts)
	err, ds := db.QueryShengNSettlementTrendtable(ts)
	if err != nil {
		return types.Statuszero, err, nil
	}
	for i, d := range *ds {
		Data[i].JiesuanMoney = utils.Fen2Yuan(d.FNbShengnjyje) //省内产生金额
		Data[i].ClearlingMoney = utils.Fen2Yuan(d.FNbShengnqkje)
		Data[i].DiffMoney = utils.Fen2Yuan(d.FNbChae)
		Data[i].DiffCount = d.FNbJiaoyts - d.FNbQingkts
		Data[i].JiesuanCount = d.FNbJiaoyts
		Data[i].ClearlingCount = d.FNbQingkts
		Data[i].DateTime = d.FVcKuaizsj
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &Data
}

//QueryDataSync
func QueryDataSync() (int, error, *[]dto.DataSync) {
	//查询海岭数据同步监控
	ts := 12
	Data := make([]dto.DataSync, ts)
	err, ds := db.QueryDataSynctable(ts)
	if err != nil {
		return types.Statuszero, err, nil
	}
	for i, d := range *ds {
		Data[i].HailCount = d.FNbJiessjzl //海岭
		Data[i].JiesuanCount = d.FNbYitbsjl
		Data[i].DateTime = d.FDtTongjwcsj.Format("2006-01-02 15:04:05")
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &Data
}

func QuerySNDataClassification() (int, error, *dto.ShengNDataClassification) {
	//查询省内结算数据分类
	err, data := db.QuerySNDataClassificationTable()
	if err != nil {
		return types.Statuszero, err, nil
	}
	if data.FVcTongjrq == "" {
		err, Data := db.QuerySNDataClassificationTableByID(data.FNbId - 1)
		if err != nil {
			return types.Statuszero, err, nil
		}
		return types.StatusSuccessfully, nil, &dto.ShengNDataClassification{
			Data.FNbShengnzjysl,
			Data.FNbQingksl,
			Data.FNbWeifssl,
			Data.FNbFassjl,
			Data.FNbjufsjl,
			//Data.FNbQingksl,
			Data.FDtTongjwcsj.Format("2006-01-02 15:04:05"),
			utils.Fen2Yuan(Data.FNbShengnzjyje),
			utils.Fen2Yuan(Data.FNbShengnqkje),
			utils.Fen2Yuan(Data.FNbShengnjfje),
		}
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &dto.ShengNDataClassification{
		data.FNbShengnzjysl,
		data.FNbQingksl,
		data.FNbWeifssl,
		data.FNbFassjl,
		data.FNbjufsjl,
		//data.FNbQingksl,
		data.FDtTongjwcsj.Format("2006-01-02 15:04:05"),
		utils.Fen2Yuan(data.FNbShengnzjyje),
		utils.Fen2Yuan(data.FNbShengnqkje),
		utils.Fen2Yuan(data.FNbShengnjfje),
	}
}

//QueryAbnormalDataParking
func QueryAbnormalDataParking() (int, error, *[]dto.AbnormalDataOfParking) {
	//查询异常数据停车场top10
	ts := 10
	Data := make([]dto.AbnormalDataOfParking, ts)
	err, data := db.QueryAbnormalDataOfParkingTable()
	if err != nil {
		return types.Statuszero, err, nil
	}
	err1, yqdatas := db.QueryAbnormalDataOfParkingtable(data.FVcKuaizsj, ts)
	if err1 != nil {
		return types.Statuszero, err1, nil
	}
	for i, yqd := range *yqdatas {
		Data[i].AbnormalDatacount = yqd.FNbZongts
		Data[i].Parkingname = yqd.FVcTingccid
		Data[i].AbnormalDataAmount = utils.Fen2Yuan(yqd.FNbZongje)
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &Data
}

//QueryOverdueData
func QueryOverdueData() (int, error, *[]dto.Overduedata) {
	ts := 10
	Data := make([]dto.Overduedata, ts)
	err, data := db.QueryOverdueDataTable()
	if fmt.Sprint(err) == "record not found" {
		log.Println("QueryOverdueData err == `record not found`:", err)
		return types.StatusSuccessfully, nil, &Data
	}
	if err != nil {
		return types.Statuszero, err, nil
	}
	err1, yqdatas := db.QueryOverdueDatatable(data.FVcTongjrq, ts)
	if err1 != nil {
		return types.Statuszero, err1, nil
	}
	for i, yqd := range *yqdatas {
		Data[i].Overduecount = yqd.FNbYuqzts
		Data[i].Parkingname = yqd.FVcTingccid
		Data[i].OverdueAmount = utils.Fen2Yuan(yqd.FNbYuqzje)
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, &Data
}

//QueryHSDZData
func QueryHSDZData() (int, error, *[]db.Hsdzdata) {
	data, err := db.QueryHSDZData()
	if err != nil {
		return types.Statuszero, err, nil
	}
	//返回数据赋值
	return types.StatusSuccessfully, nil, data
}
