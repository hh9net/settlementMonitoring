package types

import "time"

//省外数据
type JieSuanWssj struct {
	Id      int64 `gorm:"AUTO_INCREMENT primary_key"` // 设置 id 为自增类型
	JiaoyId int64 `gorm:"unique;not null"`
	Jine    int64
}

//本省结算数据
type JieSuanJiangssj struct {
	Id      int64 `gorm:"AUTO_INCREMENT primary_key"` // 设置 id 为自增类型
	JiaoyId int64 `gorm:"unique;not null"`
	Jine    int64
}

//1 结算统计监控表 b_jsjk_jiestj
type BJsjkJiestj struct {
	FNbId        int       `gorm:"AUTO_INCREMENT primary_key"` //`F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbKawlh     int       //`F_NB_KAWLH` int DEFAULT NULL COMMENT '卡网络号',
	FNbZongje    int64     //`F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int       //`F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  time.Time //`F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time //`F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   time.Time //`F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//2清分核对表 b_jsjk_qingfhd
type BJsjkQingfhd struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbQingfbxh  int64  //  `F_NB_QINGFBXH` bigint DEFAULT NULL COMMENT '清分包序号',
	FNbQingfje   int64  //  `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	FNbTongjqfje int64  //  `F_NB_TONGJQFJE` bigint DEFAULT NULL COMMENT '统计清分金额',
	FNbHedjg     int    //  `F_NB_HEDJG` int DEFAULT NULL COMMENT '核对结果 是否一致,1:一致，2:不一致',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//3 省内拒付数据统计表 `b_jsjk_shengnjfsjtj`
type BJsjkShengnjfsjtj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbJufzje    int64  //  `F_NB_JufZJE` bigint DEFAULT NULL COMMENT '拒付总金额 （分）',
	FNbJufzts    int    //  `F_NB_JufZTS` int DEFAULT NULL COMMENT '拒付总条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//4 黑名单监控表  `b_jsjk_heimdjk`
type BJsjkHeimdjk struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbHeimdzs   int    //  `F_NB_HEIMDZS` int DEFAULT NULL COMMENT '黑名单总数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj   string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//5 省内结算趋势表 CREATE TABLE `b_jsjk_shengnjsqs`
type BJsjkShengnjsqs struct {
	FNbId         int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbShengnjyje int64  //  `F_NB_SHENGNJYJE` bigint DEFAULT NULL COMMENT '省内交易金额',
	FNbShengnqksj int64  //  `F_NB_SHENGNQKJE` bigint DEFAULT NULL COMMENT '省内请款金额',
	FNbChae       int64  //  `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts    int    //  `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingkts    int    //  `F_NB_QINGKTS` int DEFAULT NULL COMMENT '请款条数',
	FDtKaistjsj   string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj    string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//6 省内结算数据分类表`b_jsjk_shengnjssjfl`
type BJsjkShengnjssjfl struct {
	FNbId          int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbShengnzjysl int    //  `F_NB_SHENGNZJYSL` int DEFAULT NULL COMMENT '省内总交易数量',
	FNbQingksl     int    //  `F_NB_QINGKSL` int DEFAULT NULL COMMENT '请款数量',
	FNbWeifssl     int    //  `F_NB_WEIFSSL` int DEFAULT NULL COMMENT '未发送数据量',
	FNbJufsjl      int    //  `F_NB_FASSJL` int DEFAULT NULL COMMENT '发送数据量',
	FNbjufsjl      int    //  `F_NB_JUFSJL` int DEFAULT NULL COMMENT '拒付数据量',
	FDtKaistjsj    string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj   string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq     string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//7 省内请款统计表 `b_jsjk_shengnqktj`
type BJsjkShengnqktj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbQingkzje  int64  //  `F_NB_QINGKZJE` bigint DEFAULT NULL COMMENT '请款总金额 （分）',
	FNbQingkzts  int    //  `F_NB_QINGKZTS` int DEFAULT NULL COMMENT '请款总条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//8  省内实时数据监控表 b_jsjk_shengnsssjjk
type BJsjkShengnsssjjk struct {
	FNbId            int    `gorm:"AUTO_INCREMENT primary_key"` //`F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbShengncsje    int64  //  `F_NB_SHENGNCSJE` bigint DEFAULT NULL COMMENT '省内产生金额',
	FNbShengnyfssjje int64  //  `F_NB_SHENGNYFSSJJE` bigint DEFAULT NULL COMMENT '省内已发送数据金额',
	FNbShengnyjzsjje int64  //  `F_NB_SHENGNYJZSJJE` bigint DEFAULT NULL COMMENT '省内已记账数据金额',
	FNbShengncsts    int    //  `F_NB_SHENGNCSTS` int DEFAULT NULL COMMENT '省内产生条数',
	FNbShengnyfssjts int    //  `F_NB_SHENGNYFSSJTS` int DEFAULT NULL COMMENT '省内已发送数据条数',
	FNbShengnyjzsjts int    //  `F_NB_SHENGNYJZSJTS` int DEFAULT NULL COMMENT '省内已记账数据条数',
	FDtKaistjsj      string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj     string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq       string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//9 省内停车场结算趋势表 `b_jsjk_shengntccjsqs`
type BJsjkShengntccjsqs struct {
	FNbId         int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbShengnjyje int64  //  `F_NB_SHENGNJYJE` bigint DEFAULT NULL COMMENT '省内交易金额',
	FNbShengnqkje int64  //  `F_NB_SHENGNQKJE` bigint DEFAULT NULL COMMENT '省内请款金额',
	FNbChae       int64  //  `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts    int    //  `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingkts    int    //  `F_NB_QINGKTS` int DEFAULT NULL COMMENT '请款条数',
	FVcTingccid   string //  `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FVcGongsid    string //  `F_VC_GONGSID` varchar(32) DEFAULT NULL COMMENT '公司id',
	FDtKaistjsj   string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj    string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//10 省内已发送数据统计表  `b_jsjk_shengnyfssjtj`
type BJsjkShengnyfssjtj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbZongje    int64  //  `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int    //  `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj   string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//11 省外结算趋势表  `b_jsjk_shengwjsqs`
type BJsjkShengwjsqs struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbJiaoye    int64  //  `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
	FNbQingdje   int64  //  `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	FNbChae      int64  //  `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts   int    //  `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingfts   int    //  `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//12 省外结算数据分类  `b_jsjk_shengwjssjfl`
type BJsjkShengwjssjfl struct {
	FNbId         int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbJiaoyzts   int    //  `F_NB_JIAOYZTS` int DEFAULT NULL COMMENT '交易总条数',
	FNbQingfsjts  int    //  `F_NB_QINGFSJTS` int DEFAULT NULL COMMENT '清分数据条数',
	FNbJizsjts    int    //  `F_NB_JIZSJTS` int DEFAULT NULL COMMENT '记账数据条数',
	FNbZhengysjts int    //  `F_NB_ZHENGYSJTS` int DEFAULT NULL COMMENT '争议数据条数 待处理',
	FNbWeidbsjts  int    //  `F_NB_WEIDBSJTS` int DEFAULT NULL COMMENT '未打包数据条数',
	FNbYidbsjts   int    //  `F_NB_YIDBSJTS` int DEFAULT NULL COMMENT '已打包数据条数',
	FNbYifssjts   int    //  `F_NB_YIFSSJTS` int DEFAULT NULL COMMENT '已发送数据条数',
	FNbHuaizsjts  int    //  `F_NB_HUAIZSJTS` int DEFAULT NULL COMMENT '坏账数据条数',
	FDtKaistjsj   string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq    string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//13 省外结算争议数据统计表 `b_jsjk_shengwjszysjtj`
type BJsjkShengwjszysjtj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbZongje    int64  //  `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int    //  `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//14 省外结算清分统计表  `b_jsjk_shengwqftj`
type BJsjkShengwqftj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbZongje    int64  //  `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int    //  `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//15 省外停车场结算趋势表 `b_jsjk_shengwtccjsqs`
type BJsjkShengwtccjsqs struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbJiaoyje   int64  //  `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
	FNbQingfje   int64  //  `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	FNbChae      int64  //  `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts   int    //  `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingfts   int    //  `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
	FVcGongsid   string //  `F_VC_GONGSID` varchar(32) DEFAULT NULL COMMENT '公司id',
	FVcTingccid  string //  `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//16 结算数据包监控表 `b_jsjk_shujbjk`
type BJsjkShujbjk struct {
	FNbId           int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbDabsl        int    //  `F_NB_DABSL` int DEFAULT NULL COMMENT '打包数量',
	FNbDabje        int    //  `F_NB_DABJE` bigint DEFAULT NULL COMMENT '打包金额',
	FNbFasysjybsl   int    //  `F_NB_FASYSJYBSL` int DEFAULT NULL COMMENT '已发送原始交易消息包数量',
	FNbFasysjybje   int64  //  `F_NB_FASYSJYBJE` bigint DEFAULT NULL COMMENT '已发送原始交易消息包金额',
	FNbJizbsl       int    //  `F_NB_JIZBSL` int DEFAULT NULL COMMENT '记账包数量',
	FNbJizbje       int64  //  `F_NB_JIZBJE` bigint DEFAULT NULL COMMENT '记账包金额',
	FNbYuansjyydbsl int    //  `F_NB_YUANSJYYDBSL` int DEFAULT NULL COMMENT '原始交易消息应答包数量',
	FDtKaistjsj     string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj    string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj      string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//17 数据同步监控表 `b_jsjk_shujtbjk`
type BJsjkShujtbjk struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbJiessjzl  int    //  `F_NB_JIESJZL` int DEFAULT NULL COMMENT '结算数据总量',F_NB_JIESSJZL
	FNbYitbsjl   int    //  `F_NB_YITBSJL` int DEFAULT NULL COMMENT '已同步数据量',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//18 停车场结算数据统计表 `b_jsjk_tingccjssjtj`
type BJsjkTingccjssjtj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FVcTingccid  string //  `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FNbkawlh     int    //  `F_NB_KAWLH` int DEFAULT NULL COMMENT '卡网络号',
	FNbZongje    int64  //  `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int    //  `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//19 异常数据停车场统计表`b_jsjk_yicsjtcctj`
type BJsjkYicsjtcctj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbZongts    int    //  `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FNbZongje    int64  //  `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额（分）',
	FVcTingccid  string //  `F_NB_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FNbTongjlx   int    //  `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:单点、2:总对总',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj   string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',

}

//20 结算待处理异常数据统计表 `b_jsjk_yicsjtj`
type BJsjkYicsjtj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbZongje    int64  //  `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int    //  `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FNbTongjlx   int    //  `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:总对总、2:单点',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//21 逾期数据统计表 `b_jsjk_yuqsjtj`
type BJsjkYuqsjtj struct {
	FNbId        int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbYuqzts    int    //  `F_NB_YUQZTS` int DEFAULT NULL COMMENT '逾期总条数',
	FNbYuqzje    int64  //  `F_NB_YUQZJE` bigint DEFAULT NULL COMMENT '逾期总金额 （分）',
	FVcTingccid  string //  `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FDtKaistjsj  string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string //  `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//22 转结算数据监控表 `b_jsjk_zhuanjssjjk`
type BJsjkZhuanjssjjk struct {
	FNbId         int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbChedyssjts int    //  `F_NB_CHEDYSSJTS` int DEFAULT NULL COMMENT '车道原始数据条数',
	FNbJiesbsjts  int    //  `F_NB_JIESBSJTS` int DEFAULT NULL COMMENT '结算表数据条数',
	FNbTongjlx    int    //  `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:单点、2:总对总',
	FDtKaistjsj   string //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj    string //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//23 用户表 b_jsjk_jiesjkptyhb
type BJsjkJiesjkptyhb struct {
	//	FNbId      int    `gorm:"AUTO_INCREMENT primary_key"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',
	FNbYonghid string //   '用户id',//手机号 F_NB_YONGHID
	FVcYonghmm string //   '用户密码',
	FVcShoujh  string //   '手机号',
	FVcYoux    string //   '邮箱',
	FVcYonghnc string //   '用户昵称',
}
