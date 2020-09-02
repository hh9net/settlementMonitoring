package service

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
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
	rhgeterr, value := utils.RedisGet(conn, "swjiesuantotal")
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
		hmerr, hmdsj := db.QueryBlacklisttableByID(id)
		if hmerr != nil {
			log.Println(hmerr)
			return 0, hmerr, nil
		}

		hmerr2, hmdsj2 := db.QueryBlacklisttableByID(id - 2)
		if hmerr2 != nil {
			log.Println(hmerr2)
			return 0, hmerr2, nil
		}
		ts := 24 //需要查询条数
		qdterr, hmdjls := db.QueryBlacklistTiaoshutable(id, ts)
		if qdterr != nil {
			log.Println(qdterr)
			return 0, qdterr, nil
		}
		bs := make([]dto.BlackList, 24)
		for i, b := range *hmdjls {
			bs[i].Blacklistcount = b.FNbHeimdzs
			bs[i].DateTime = b.FDtTongjwcsj.Format("2006-01-02 15:04:05")
		}
		bs12 := make([]dto.BlackList, 0)
		for i, blist := range bs {
			if i == 0 || i == 2 || i == 4 || i == 6 || i == 8 || i == 10 || i == 12 || i == 14 || i == 16 || i == 18 || i == 20 || i == 22 {
				bs12 = append(bs12, blist)
			}
		}
		changecount := hmdsj.FNbHeimdzs - hmdsj2.FNbHeimdzs
		log.Println("查询黑名单总数、较2小时前变化值  成功")
		return 208, nil, &dto.TotalBlacklistData{Blacklistcount: hmdsj.FNbHeimdzs,
			ChangeCount: changecount,
			DateTime:    hmdsj.FDtTongjwcsj.Format("2006-01-02 15:04:05"),
			Blacklist:   bs12}
	}

	hmerr, hmdsj := db.QueryBlacklisttableByID(id)
	if hmerr != nil {
		log.Println(hmerr)
		return 0, hmerr, nil
	}

	hmerr2, hmdsj2 := db.QueryBlacklisttableByID(id - 2)
	if hmerr2 != nil {
		log.Println(hmerr2)
		return 0, hmerr2, nil
	}
	ts := 24 //需要查询条数
	qdterr, hmdjls := db.QueryBlacklistTiaoshutable(id, ts)
	if qdterr != nil {
		log.Println(qdterr)
		return 0, qdterr, nil
	}
	bs := make([]dto.BlackList, 24)
	for i, b := range *hmdjls {
		bs[i].Blacklistcount = b.FNbHeimdzs
		bs[i].DateTime = b.FDtTongjwcsj.Format("2006-01-02 15:04:05")
	}
	bs12 := make([]dto.BlackList, 0)
	for i, blist := range bs {
		if i == 0 || i == 2 || i == 4 || i == 6 || i == 8 || i == 10 || i == 12 || i == 14 || i == 16 || i == 18 || i == 20 || i == 22 {
			bs12 = append(bs12, blist)
		}
	}
	changecount := hmdsj.FNbHeimdzs - hmdsj2.FNbHeimdzs
	log.Println("查询黑名单总数、较2小时前变化值  成功")
	//返回数据赋值
	return 208, nil, &dto.TotalBlacklistData{Blacklistcount: hmdsj.FNbHeimdzs,
		ChangeCount: changecount,
		DateTime:    hmdsj.FDtTongjwcsj.Format("2006-01-02 15:04:05"),
		Blacklist:   bs12}
}

//查询清分包、争议包的接收时间、包号 14天 存redis中的
func QueryClearlingAndDisputePackagedata() (int, error, *dto.ClearlAndDisputeData) {

	//查询清分包、争议包的接收时间、包号[最新的数据]前14天数据[1天]
	date := utils.OldData(14)
	chmgeterr, cleardata := utils.RedisHMGet(utils.RedisInit(), "clear", date)
	if chmgeterr != nil {
		return 0, chmgeterr, nil
	}

	dhmgeterr, disputdata := utils.RedisHMGet(utils.RedisInit(), "disput", date)
	if dhmgeterr != nil {
		return 0, dhmgeterr, nil
	}
	data := make([]types.ClearlingAndDisputeData, 0)

	for _, clearData := range *cleardata {
		clearAndDis := new(types.ClearlingAndDisputeData)
		if clearData == "no data" {
			data = append(data, types.ClearlingAndDisputeData{ClearPacgNo: ""})
			continue
		}
		vs := strings.Split(clearData, `"`)
		str := strings.Split(vs[1], `|`)
		clearAndDis.ClearPacgNo = str[0]
		clearAndDis.Cleardatetime = str[1]
		data = append(data, *clearAndDis)
	}

	for i, disputData := range *disputdata {
		if disputData == "no data" {
			//data = append(data, types.ClearlingAndDisputeData{ClearPacgNo: "no clear package", DisputPacgeNo: "no disput package"})
			data[i].DisputPacgeNo = ""
			continue
		}
		dvs := strings.Split(disputData, `"`)
		disstr := strings.Split(dvs[1], `|`)
		data[i].Disputdatetime = disstr[1]
		data[i].DisputPacgeNo = disstr[0]
	}
	for i, da := range date {
		data[i].Date = da
	}
	log.Println("查询清分包、争议包的接收时间、包号  成功。data数组长度:", len(data))
	//if len(data)!=
	//返回数据赋值
	return 209, nil, &dto.ClearlAndDisputeData{data}
}

//省外清分核对 StatisticalClearlingcheck
func StatisticalClearlingcheck() (int, error, *[]dto.ClearlingcheckData) {
	//清分核对
	ts := 100
	err, checkdata := db.QueryCheckResultbyTs(ts)
	if err != nil {
		return 0, err, nil
	}
	log.Println("清分核对结果:", checkdata)
	//返回数据赋值
	Data := make([]dto.ClearlingcheckData, len(*checkdata))
	for i, data := range *checkdata {
		Data[i].Clearlingpakgje = utils.Fen2Yuan(data.FNbQingfje)
		Data[i].Clearlingpakgxh = data.FNbQingfbxh
		Data[i].Clearlingpakgts = data.FNbQingfts
		Data[i].Tongjqfje = utils.Fen2Yuan(data.FNbTongjqfje)
		Data[i].Tongjqfts = data.FNbTongjqfts
		Data[i].Hedjg = data.FNbHedjg
		Data[i].Tongjrq = data.FVcTongjrq
		Data[i].Qingfbjssj = utils.DateTimeFormat(data.FDtQingfbjssj)
	}
	return 210, nil, &Data
}

//省外数据分类
func Dataclassification() (int, error, *dto.Dataclassification) {

	//查询记录
	err, dataclassification := db.QuerySWDataClassificationTable()
	if err != nil {
		return 0, err, nil
	}
	log.Println("清分核对结果:", dataclassification)

	if dataclassification.FNbJiaoyzts == 0 {
		err2, data := db.QuerySWDataClassificationTableByID(dataclassification.FNbId - 1)
		if err2 != nil {
			return 0, err2, nil
		}
		return 211, nil, &dto.Dataclassification{
			Shengwzcount: data.FNbJiaoyzts,
			Yiqfcount:    data.FNbQingfsjts,  //已清分总条数（不含坏账）
			Jizcount:     data.FNbJizsjts,    //记账
			Zhengycount:  data.FNbZhengysjts, //争议
			Weidbcount:   data.FNbWeidbsjts,  //未打包
			Yidbcount:    data.FNbYidbsjts,   //已打包
			Yifscount:    data.FNbYifssjts,   //已发送
			Huaizcount:   data.FNbHuaizsjts,  //坏账

		}
	}

	//返回数据赋值
	return 211, nil, &dto.Dataclassification{Shengwzcount: dataclassification.FNbJiaoyzts,
		Yiqfcount:   dataclassification.FNbQingfsjts,  //已清分总条数（不含坏账）
		Jizcount:    dataclassification.FNbJizsjts,    //记账
		Zhengycount: dataclassification.FNbZhengysjts, //争议
		Weidbcount:  dataclassification.FNbWeidbsjts,  //未打包
		Yidbcount:   dataclassification.FNbYidbsjts,   //已打包
		Yifscount:   dataclassification.FNbYifssjts,   //已发送
		Huaizcount:  dataclassification.FNbHuaizsjts,  //坏账
	}
}

//转结算
func QueryDataTurnMonitordata() (int, error, *[]dto.TurnDataResponse) {
	ts := 24
	//响应数据 list TurnDataResponse
	TurndataResps := make([]dto.TurnDataResponse, ts)
	datas := make([]dto.TurnData, ts)
	//查询数据  '统计类型 1:单点、2:总对总',
	//dd 1:单点、2:总对总'
	ddqerr, ddtds := db.QueryDataTurnMonitortable(ts, 1)
	if ddqerr != nil {
		return 0, ddqerr, nil
	}
	for i, dd := range *ddtds {
		datas[i].Jieszcount = dd.FNbJiesbsjts
		datas[i].DDzcount = dd.FNbChedyssjts //dd
		datas[i].DateTime = dd.FDtTongjwcsj.Format("2006-01-02 15:04:05")
	}

	//zdz 1:单点、2:总对总'
	zdzqerr, zdztds := db.QueryDataTurnMonitortable(ts, 2)
	if zdzqerr != nil {
		return 0, zdzqerr, nil
	}
	for i, dd := range *zdztds {
		datas[i].ZDZcount = dd.FNbChedyssjts //zdz
	}
	log.Println("datas", datas)

	//处理数据
	for i, r := range datas {
		TurndataResps[i].JieszCount = r.Jieszcount
		TurndataResps[i].YuansCount = r.ZDZcount + r.DDzcount
		TurndataResps[i].DifferCount = TurndataResps[i].YuansCount - r.Jieszcount
		TurndataResps[i].DateTime = r.DateTime
	}
	log.Println("响应数据：", TurndataResps)
	//返回数据
	return 212, nil, &TurndataResps
}

//结算趋势
func QuerySettlementTrend() (int, error, *[]dto.SettlementTrend) {
	ts := 30
	//响应数据 list TurnDataResponse
	Datas := make([]dto.SettlementTrend, ts)
	//查询数据
	qerr, ds := db.QuerySettlementTrendbyday(ts)
	if qerr != nil {
		return 0, qerr, nil
	}

	for i, d := range *ds {
		Datas[i].JiesuanAmount = utils.Fen2Yuan(d.FNbJiaoye)
		Datas[i].QingfAmount = utils.Fen2Yuan(d.FNbQingdje)
		Datas[i].DifferAmount = utils.Fen2Yuan(d.FNbChae)
		Datas[i].JiesuanCount = d.FNbJiaoyts
		Datas[i].QingfCount = d.FNbQingfts
		Datas[i].DateTime = d.FDtTongjwcsj.Format("2006-01-02 15:04:05")
	}

	log.Println("响应数据：", Datas)
	//返回数据
	return 213, nil, &Datas
}

// 查询省外数据包监控  144条
func QueryPacketMonitoring() (int, error, *[]dto.PacketMonitoringdata) {
	ts := types.Frequency
	//响应数据 list TurnDataResponse
	Datas := make([]dto.PacketMonitoringdata, ts)

	//查询数据
	qerr, ds := db.QueryPacketMonitoringtable(ts)
	if qerr != nil {
		return 0, qerr, nil
	}

	for i, d := range *ds {
		Datas[i].Yuansyingdbsl = d.FNbYuansjyydbsl
		Datas[i].Dabaojine = utils.Fen2Yuan(d.FNbDabje)
		Datas[i].Dabaosl = d.FNbDabsl
		Datas[i].Fasbjine = utils.Fen2Yuan(d.FNbFasysjybje)
		Datas[i].Fasbsl = d.FNbFasysjybsl
		Datas[i].Jizbjine = utils.Fen2Yuan(d.FNbJizbje)
		Datas[i].Jizbsl = d.FNbJizbsl
		Datas[i].DateTime = d.FDtTongjwcsj.Format("2006-01-02 15:04:05")
	}

	log.Println("响应数据：", Datas)
	//返回数据
	return 214, nil, &Datas
}

//查询最近15天清分包数据差额
func Clarifydifference() (int, error, *[]dto.DifferAmount) {
	ts := 15
	//响应数据 list TurnDataResponse
	Datas := make([]dto.DifferAmount, ts)

	//查询数据
	qerr, ds := db.QuerySettlementclearlingcheck(ts)
	if qerr != nil {
		return 0, qerr, nil
	}

	for i, d := range *ds {
		Datas[i].Differamount = utils.Fen2Yuan(d.FNbQingfje - d.FNbTongjqfje)
		Datas[i].DateTime = d.FVcTongjrq
	}
	log.Println("响应数据：", Datas)
	//返回数据
	return 215, nil, &Datas
}

func ClarifyQuery(req dto.ReqQueryClarify) (int, error, *dto.Clearlingcheckdata) {
	log.Print("清分核对请求参数：", req)
	//获取请求数据
	//var err error
	//var qfhdreqs *[]types.BJsjkQingfhd

	err, qfhdreqs, zongjls, zongys := db.QueryClearlingcheck(&req)
	if err != nil {
		if fmt.Sprint(err) == "请输入开始查询时间" {
			//查询用户是否被注册，查询失败
			return 0, err, nil
		}
		if fmt.Sprint(err) == "请输入查询截止时间" {
			//查询用户是否被注册，查询失败
			return 0, err, nil
		}

		if fmt.Sprint(err) == "请输入正确的每页展示记录数" {
			//查询用户是否被注册，查询失败
			return 0, err, nil
		}

		//查询用户是否被注册，查询失败
		return 0, err, nil
	}
	//响应数据 list
	Datas := make([]dto.ClearlingcheckData, len(*qfhdreqs))
	for i, d := range *qfhdreqs {
		Datas[i].Clearlingpakgxh = d.FNbQingfbxh
		Datas[i].Clearlingpakgje = utils.Fen2Yuan(d.FNbQingfje)
		Datas[i].Clearlingpakgts = d.FNbQingfts
		Datas[i].Tongjqfje = utils.Fen2Yuan(d.FNbTongjqfje)
		Datas[i].Tongjqfts = d.FNbTongjqfts
		Datas[i].Hedjg = d.FNbHedjg
		Datas[i].Tongjrq = d.FVcTongjrq
		Datas[i].Qingfbjssj = utils.DateTimeFormat(d.FDtQingfbjssj)
	}
	Data := dto.Clearlingcheckdata{
		Clearlingcheck: Datas,
		ZongTS:         zongjls,
		ZongYS:         zongys,
	}
	log.Println("响应数据：", Data)
	//返回数据
	return 216, nil, &Data
}

//清分确认【未实现】
func Clarifyconfirm() (int, error, *[]dto.PacketMonitoringdata) {
	ts := 30
	//响应数据 list TurnDataResponse
	Datas := make([]dto.PacketMonitoringdata, ts)

	//查询数据
	qerr, ds := db.QueryPacketMonitoringtable(ts)
	if qerr != nil {
		return 0, qerr, nil
	}

	for i, d := range *ds {
		Datas[i].Yuansyingdbsl = d.FNbYuansjyydbsl
	}

	log.Println("响应数据：", Datas)
	//返回数据
	return 217, nil, &Datas
}

func ExportExcel(req dto.ReqClarifyExportExcel) (int, error, []byte, string) {
	log.Print("清分核对导出请求参数：", req)
	//获取请求数据

	err, qfhdreqs := db.QueryClearlingcheckByConditions(&req)
	if err != nil {
		if fmt.Sprint(err) == "请输入开始查询时间" {
			//查询用户是否被注册，查询失败
			return 0, err, nil, ""
		}
		if fmt.Sprint(err) == "请输入查询截止时间" {
			//查询用户是否被注册，查询失败
			return 0, err, nil, ""
		}
	}
	// ExportExcel 导出Excel文件
	// sheetName 工作表名称
	// columns 列名切片
	// rows 数据切片，是一个二维数组
	//  ExportExcel(sheetName string, columns []string, rows [][]interface{})
	columns := []string{"统计日期", "统计清分数据(条)", "统计清分金额(元)", "清分包接收(条)", "清分包金额(元)", "清分包接收时间", "清分包编号", "校验状态"}
	rows := make([][]interface{}, 0)
	for _, qfhdsj := range *qfhdreqs {
		row := make([]interface{}, 0)
		row = append(row, qfhdsj.FVcTongjrq)
		row = append(row, qfhdsj.FNbTongjqfts)
		row = append(row, utils.Fen2Yuan(qfhdsj.FNbTongjqfje))
		row = append(row, qfhdsj.FNbQingfts)
		row = append(row, utils.Fen2Yuan(qfhdsj.FNbQingfje))
		row = append(row, qfhdsj.FDtQingfbjssj)
		row = append(row, qfhdsj.FNbQingfbxh)
		if qfhdsj.FNbHedjg == 1 {
			row = append(row, "校验成功")
		}
		if qfhdsj.FNbHedjg == 2 {
			row = append(row, "校验失败")
		}
		rows = append(rows, row)
	}
	log.Println("导出文件名为：清分包数据核对记录表.xlsx 成功", rows[0])

	path := utils.ExportExcel("清分包数据核对记录", columns, rows)

	log.Println("要发送excle 文件的path:=", path)
	file, oserr := os.Open("./" + path)
	if oserr != nil {
		log.Println("os.Open error:", oserr)
		return 0, oserr, nil, ""
	}
	data, rerr := ioutil.ReadAll(file)
	if rerr != nil {
		return 0, rerr, nil, ""
	}
	defer file.Close()

	//返回数据
	return 218, nil, data, path
}
