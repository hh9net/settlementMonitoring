package db

import (
	"github.com/jinzhu/gorm"
	"log"
	"settlementMonitoring/config"
	"settlementMonitoring/utils"
	"time"
)

var HSDZGormClientDB *gorm.DB

//接收汇总表 由省内结算平台维护 恒生 `msg_recv_statistics` (
type MsgRecvStatistics struct {
	RecordId       int       `gorm:"column:RECORD_ID"`        //	`RECORD_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '汇总记录ID',
	MsgDate        time.Time `gorm:"column:MSG_DATE"`         //	`MSG_DATE` date DEFAULT NULL COMMENT '数据日期',
	SendCount      int       `gorm:"column:SEND_COUNT"`       //	`SEND_COUNT` int(11) DEFAULT NULL COMMENT '已发送的数量',//恒生接收
	SendSum        int       `gorm:"column:SEND_SUM"`         //	`SEND_SUM` int(11) DEFAULT NULL COMMENT '已发送的金额',
	AcceptedCount  int       `gorm:"column:ACCEPTED_COUNT"`   //	`ACCEPTED_COUNT` int(11) DEFAULT NULL COMMENT '接收数量 原始交易数据ACK为0的数量',
	AcceptedSum    int       `gorm:"column:ACCEPTED_SUM"`     //	`ACCEPTED_SUM` int(11) DEFAULT NULL COMMENT '接收金额',
	RefusedCount   int       `gorm:"column:REFUSED_COUNT"`    //	`REFUSED_COUNT` int(11) DEFAULT NULL COMMENT '拒收数量 原始交易数据ACK非0的数量',
	RefusedSum     int       `gorm:"column:REFUSED_SUM"`      //	`REFUSED_SUM` int(11) DEFAULT NULL COMMENT '拒收金额',
	CreateTime     time.Time `gorm:"column:CREATE_TIME"`      //	`CREATE_TIME` datetime DEFAULT NULL COMMENT '记录生成时间',
	LastUpdateTime time.Time `gorm:"column:LAST_UPDATE_TIME"` //	`LAST_UPDATE_TIME` datetime DEFAULT NULL COMMENT '最后更新时间',
}

//发送汇总表 由停车场接入平台维护 `msg_send_statistics` (
type MsgSendStatistics struct {
	RecordId       int       `gorm:"column:RECORD_ID"`        //	`RECORD_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '汇总记录ID',
	MsgDate        time.Time `gorm:"column:MSG_DATE"`         //	`MSG_DATE` date DEFAULT NULL COMMENT '数据日期',
	SendCount      int       `gorm:"column:SEND_COUNT"`       //	`SEND_COUNT` int(11) DEFAULT NULL COMMENT '已发送的数量',//恒生接收
	SendSum        int       `gorm:"column:SEND_SUM"`         //	`SEND_SUM` int(11) DEFAULT NULL COMMENT '已发送的金额',
	AcceptedCount  int       `gorm:"column:ACCEPTED_COUNT"`   //	`ACCEPTED_COUNT` int(11) DEFAULT NULL COMMENT '接收数量 原始交易数据ACK为0的数量',
	AcceptedSum    int       `gorm:"column:ACCEPTED_SUM"`     //	`ACCEPTED_SUM` int(11) DEFAULT NULL COMMENT '接收金额',
	RefusedCount   int       `gorm:"column:REFUSED_COUNT"`    //	`REFUSED_COUNT` int(11) DEFAULT NULL COMMENT '拒收数量 原始交易数据ACK非0的数量',
	RefusedSum     int       `gorm:"column:REFUSED_SUM"`      //	`REFUSED_SUM` int(11) DEFAULT NULL COMMENT '拒收金额',
	CreateTime     time.Time `gorm:"column:CREATE_TIME"`      //	`CREATE_TIME` datetime DEFAULT NULL COMMENT '记录生成时间',
	LastUpdateTime time.Time `gorm:"column:LAST_UPDATE_TIME"` //    `LAST_UPDATE_TIME` datetime DEFAULT NULL COMMENT '最后更新时间',
}

type Hsdzdata struct {
	SendCount  int    `json:"send_count"  example:"123"`
	SendAmount string `json:"send_amount"  example:"123"`
	RecvCount  int    `json:"recv_count"  example:"123"`
	RecvAmount string `json:"recv_amount"  example:"123"`
	Datetime   string `json:"datetime"  example:"123"`
}

func NewHSZDDB() {
	conf := config.ConfigInit() //初始化配置
	HSDZstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + "localsettle" + "?charset=utf8&parseTime=true&loc=Local"
	log.Println("++++++++++HSDZstr :=", HSDZstr)
	HSDZGormClientDB = utils.HSDZInitGormDB(HSDZstr)
}

func QueryHSDZData() (*[]Hsdzdata, error) {
	db := HSDZGormClientDB
	//发送数量  最近7天
	ts := 7
	hsjss := make([]MsgRecvStatistics, 0)
	fss := make([]MsgSendStatistics, 0)

	// 发送
	if err := db.Table("msg_send_statistics").Order("RECORD_ID desc").Limit(ts).Find(&fss).Error; err != nil {
		log.Println("查询最新的7条发送数据 error :", err)
	}
	log.Println("fss:", fss)
	//恒生接收
	if err := db.Table("msg_recv_statistics").Order("RECORD_ID desc").Limit(ts).Find(&hsjss).Error; err != nil {
		log.Println("查询最新的7条恒生接收数据 error :", err)
	}
	log.Println("hsjs:", hsjss)
	hsdzs := make([]Hsdzdata, ts)
	for i, fs := range fss {
		hsdzs[i].SendCount = fs.SendCount
		fssum := utils.Fen2Yuan(int64(fs.SendSum))
		hsdzs[i].SendAmount = fssum
		hsdzs[i].Datetime = fs.MsgDate.Format("2006-01-02")
	}

	for i, hsjs := range hsjss {
		if hsdzs[i].Datetime == hsjs.MsgDate.Format("2006-01-02") {
			hsdzs[i].RecvCount = hsjs.SendCount
			fssum := utils.Fen2Yuan(int64(hsjs.SendSum))
			hsdzs[i].RecvAmount = fssum
		}
	}
	log.Println("恒生对帐数据：", hsdzs)
	return &hsdzs, nil
}
