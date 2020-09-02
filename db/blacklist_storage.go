package db

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"settlementMonitoring/types"
	"settlementMonitoring/utils"
)

//查表是否存在[测试]
func QueryblacklistTable() error {
	//连接黑名单的数据库
	HmdDBInit()
	//查表是否存在
	db := GormClient.Client
	log.Println("表数量", len(types.Blacklist)) //表数量 64
	for _, b := range types.Blacklist {
		if db.HasTable(b) == true {
			log.Println(b, "表存在")
		} else {
			log.Println(b, "表不存在")
		}
	}
	return nil
}

//获取黑名单总数
func QueryBlacklistcount() (error, int) {
	//连接黑名单的数据库
	HmdDBInit()
	var Count int
	bs := 0
	for _, b := range types.Blacklist {
		qerr, count := Queryblacklistcount(b)
		if qerr != nil {
			return qerr, 0
		}
		bs = bs + 1
		Count = count + Count
	}
	if bs != 64 {
		log.Printf("表数%d获取黑名单总数:%s", bs, utils.IntToChinesenum(Count))
		log.Fatal("统计表数不对，没有统计完")
		return errors.New("统计表数不对，没有统计完"), 0
	}
	log.Printf("表数%d获取黑名单总数:%s", bs, utils.IntToChinesenum(Count))
	return nil, Count
}

//获取某一个表的黑名单总记录数
func Queryblacklistcount(bname string) (error, int) {
	db := HmdGormClient.HmdClient
	var count int
	if err := db.Table(bname).Count(&count).Error; err != nil {
		log.Printf("查询此表：%s黑名单总数 失败", bname)
		return err, 0
	}
	log.Printf("查询此表：%s黑名单总数 %d", bname, count)
	return nil, count
}

//2、新增黑名单总记录的统计开始记录
func BlacklistDataInsert() error {
	db := utils.GormClient.Client
	yctj := new(types.BJsjkHeimdjk)
	//赋值
	yctj.FDtKaistjsj = utils.StrTimeToNowtime()      //开始统计时间
	yctj.FDtTongjwcsj = utils.StrTimeTodefaultdate() //统计完成时间

	if err := db.Table("b_jsjk_heimdjk").Create(&yctj).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_jsjk_heimdjk error", err)
		return err
	}
	log.Println("新增黑名单总记录的统计开始记录成功！", yctj.FDtKaistjsj)
	return nil
}

//3、查询最新的黑名单总记录统计记录
func QueryBlacklisttable() (error, *types.BJsjkHeimdjk) {
	db := utils.GormClient.Client
	hmdtjs := new(types.BJsjkHeimdjk)
	//赋值
	if err := db.Table("b_jsjk_heimdjk").Last(&hmdtjs).Error; err != nil {
		log.Println("查询最新的黑名单数据的数据记录时，QueryBlacklisttable error :", err)
		return err, nil
	}
	log.Println("查询最新的黑名单数据的数据记录结果:", hmdtjs)
	return nil, hmdtjs
}

func QueryBlacklisttableByID(id int) (error, *types.BJsjkHeimdjk) {
	db := utils.GormClient.Client
	hmdtjs := new(types.BJsjkHeimdjk)
	//赋值
	if err := db.Table("b_jsjk_heimdjk").Where("F_NB_ID=?", id).Last(&hmdtjs).Error; err != nil {
		log.Println("查询最新的黑名单数据的数据记录时，QueryBlacklisttable error :", err)
		return err, nil
	}
	log.Println("查询最新的黑名单数据的数据记录结果:", hmdtjs)
	return nil, hmdtjs
}

//4、更新最新的黑名单总记录统计记录
func UpdateBlacklistlData(data *types.BJsjkHeimdjk, id int) error {
	//Newdb()
	db := utils.GormClient.Client
	hmdtj := new(types.BJsjkHeimdjk)

	hmdtj.FDtTongjwcsj = data.FDtTongjwcsj //统计完成时间
	hmdtj.FVcKuaizsj = data.FVcKuaizsj     //快照时间
	hmdtj.FNbHeimdzs = data.FNbHeimdzs     //黑名单总数
	if err := db.Table("b_jsjk_heimdjk").Where("F_NB_ID=?", id).Updates(&hmdtj).Error; err != nil {
		log.Println("最新的黑名单的数据记录 error", err)
		return err
	}
	log.Println("更新最新的黑名单的数据记录成功+++++++++++++++++++++++ ")

	return nil
}

//5、查询最新的ts条黑名单总记录统计记录
func QueryBlacklistTiaoshutable(id, ts int) (error, *[]types.BJsjkHeimdjk) {
	db := utils.GormClient.Client
	hmdtjs := make([]types.BJsjkHeimdjk, 0)
	//赋值
	if err := db.Table("b_jsjk_heimdjk").Where("F_NB_ID <= ? ", id).Order("F_NB_ID desc").Limit(ts).Find(&hmdtjs).Error; err != nil {
		log.Println("查询最新的黑名单数据的数据记录时，QueryBlacklisttable error :", err)
		return err, nil
	}
	log.Println("查询最新的黑名单数据的数据记录结果:", hmdtjs)
	return nil, &hmdtjs
}
