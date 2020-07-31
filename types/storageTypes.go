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
	FNbId int `gorm:"column:F_NB_ID; AUTO_INCREMENT ;primary_key"` //`F_NB_ID` '唯一自增id',

	FNbKawlh     int       `gorm:"column:F_NB_KAWLH"`     // `F_NB_KAWLH` int DEFAULT NULL COMMENT '卡网络号',
	FNbZongje    int64     `gorm:"column:F_NB_ZONGJE"`    // `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int       `gorm:"column:F_NB_ZONGTS"`    // `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  time.Time `gorm:"column:F_DT_KAISTJSJ"`  // `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time `gorm:"column:F_DT_TONGJWCSJ"` // `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string    `gorm:"column:F_VC_TONGJRQ"`   //`F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//2清分核对表 b_jsjk_qingfhd
type BJsjkQingfhd struct {
	FNbId int `gorm:"AUTO_INCREMENT ;primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbQingfbxh  int64  `gorm:"column:F_NB_QINGFBXH"`  //   `F_NB_QINGFBXH` bigint DEFAULT NULL COMMENT '清分包序号',
	FNbQingfje   int64  `gorm:"column:F_NB_QINGFJE"`   //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	FNbTongjqfje int64  `gorm:"column:F_NB_TONGJQFJE"` //   `F_NB_TONGJQFJE` bigint DEFAULT NULL COMMENT '统计清分金额',
	FNbHedjg     int    `gorm:"column:F_NB_HEDJG"`     //   `F_NB_HEDJG` int DEFAULT NULL COMMENT '核对结果 是否一致,1:一致，2:不一致',
	FDtTongjrq   string `gorm:"column:F_DT_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//3 省内拒付数据统计表 `b_jsjk_shengnjfsjtj`
type BJsjkShengnjfsjtj struct {
	FNbId int `gorm:"AUTO_INCREMENT ;primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbJufzts    int    `gorm:"column:F_NB_JufZTS"`    //   `F_NB_JufZTS` int DEFAULT NULL COMMENT '拒付总条数',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string `gorm:"column:F_DT_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//4 黑名单监控表  `b_jsjk_heimdjk`
type BJsjkHeimdjk struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbHeimdzs   int       `gorm:"column:F_NB_HEIMDZS"`   //   `F_NB_HEIMDZS` int DEFAULT NULL COMMENT '黑名单总数',
	FDtKaistjsj  time.Time `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcKuaizsj   string    `gorm:"column:F_VC_KUAIZSJ"`   //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//5 省内结算趋势表 CREATE TABLE `b_jsjk_shengnjsqs`
type BJsjkShengnjsqs struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbShengnjyje int64  `gorm:"column:F_NB_SHENGNJYJE "` //   `F_NB_SHENGNJYJE` bigint DEFAULT NULL COMMENT '省内交易金额',
	FNbShengnqksj int64  `gorm:"column:F_NB_SHENGNQKJE "` //   `F_NB_SHENGNQKJE` bigint DEFAULT NULL COMMENT '省内请款金额',
	FNbChae       int64  `gorm:"column:F_NB_CHAE"`        //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts    int    `gorm:"column:F_NB_JIAOYTS"`     //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingkts    int    `gorm:"column:F_NB_QINGKTS"`     //   `F_NB_QINGKTS` int DEFAULT NULL COMMENT '请款条数',
	FDtKaistjsj   string `gorm:"column:F_DT_KAISTJSJ"`    //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string `gorm:"column:F_DT_TONGJWCSJ"`   //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj    string `gorm:"column:F_DT_KUAIZSJ"`     //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//6 省内结算数据分类表`b_jsjk_shengnjssjfl`
type BJsjkShengnjssjfl struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbShengnzjysl int    `gorm:"column:F_NB_SHENGNZJYSL "` //   `F_NB_SHENGNZJYSL` int DEFAULT NULL COMMENT '省内总交易数量',
	FNbQingksl     int    `gorm:"column:F_NB_QINGKSL"`      //   `F_NB_QINGKSL` int DEFAULT NULL COMMENT '请款数量',
	FNbWeifssl     int    `gorm:"column:F_NB_WEIFSSL"`      //   `F_NB_WEIFSSL` int DEFAULT NULL COMMENT '未发送数据量',
	FNbJufsjl      int    `gorm:"column:F_NB_FASSJL"`       //   `F_NB_FASSJL` int DEFAULT NULL COMMENT '发送数据量',
	FNbjufsjl      int    `gorm:"column:F_NB_JUFSJL"`       //   `F_NB_JUFSJL` int DEFAULT NULL COMMENT '拒付数据量',
	FDtKaistjsj    string `gorm:"column:F_DT_KAISTJSJ"`     //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj   string `gorm:"column:F_DT_TONGJWCSJ"`    //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq     string `gorm:"column:F_DT_TONGJRQ"`      //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//7 省内请款统计表 `b_jsjk_shengnqktj`
type BJsjkShengnqktj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbQingkzje  int64  `gorm:"column:F_NB_QINGKZJE"`  //   `F_NB_QINGKZJE` bigint DEFAULT NULL COMMENT '请款总金额 （分）',
	FNbQingkzts  int    `gorm:"column:F_NB_QINGKZTS"`  //   `F_NB_QINGKZTS` int DEFAULT NULL COMMENT '请款总条数',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string `gorm:"column:F_DT_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//8  省内实时数据监控表 b_jsjk_shengnsssjjk
type BJsjkShengnsssjjk struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key; column:F_NB_ID"` //`F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbShengncsje    int64  `gorm:"column:F_NB_SHENGNCSJE"`    //   `F_NB_SHENGNCSJE` bigint DEFAULT NULL COMMENT '省内产生金额',
	FNbShengnyfssjje int64  `gorm:"column:F_NB_SHENGNYFSSJJE"` //   `F_NB_SHENGNYFSSJJE` bigint DEFAULT NULL COMMENT '省内已发送数据金额',
	FNbShengnyjzsjje int64  `gorm:"column:F_NB_SHENGNYJZSJJE"` //   `F_NB_SHENGNYJZSJJE` bigint DEFAULT NULL COMMENT '省内已记账数据金额',
	FNbShengncsts    int    `gorm:"column:F_NB_SHENGNCSTS"`    //   `F_NB_SHENGNCSTS` int DEFAULT NULL COMMENT '省内产生条数',
	FNbShengnyfssjts int    `gorm:"column:F_NB_SHENGNYFSSJTS"` //   `F_NB_SHENGNYFSSJTS` int DEFAULT NULL COMMENT '省内已发送数据条数',
	FNbShengnyjzsjts int    `gorm:"column:F_NB_SHENGNYJZSJTS"` //   `F_NB_SHENGNYJZSJTS` int DEFAULT NULL COMMENT '省内已记账数据条数',
	FDtKaistjsj      string `gorm:"column:F_DT_KAISTJSJ"`      //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj     string `gorm:"column:F_DT_TONGJWCSJ"`     //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq       string `gorm:"column:F_DT_TONGJRQ"`       //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//9 省内停车场结算趋势表 `b_jsjk_shengntccjsqs`
type BJsjkShengntccjsqs struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbShengnjyje int64  `gorm:"column:F_NB_SHENGNJYJE"` //   `F_NB_SHENGNJYJE` bigint DEFAULT NULL COMMENT '省内交易金额',
	FNbShengnqkje int64  `gorm:"column:F_NB_SHENGNQKJE"` //   `F_NB_SHENGNQKJE` bigint DEFAULT NULL COMMENT '省内请款金额',
	FNbChae       int64  `gorm:"column:F_NB_CHAE"`       //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts    int    `gorm:"column:F_NB_JIAOYTS"`    //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingkts    int    `gorm:"column:F_NB_QINGKTS"`    //   `F_NB_QINGKTS` int DEFAULT NULL COMMENT '请款条数',
	FVcTingccid   string `gorm:"column:F_VC_TINGCCID"`   //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FVcGongsid    string `gorm:"column:F_VC_GONGSID"`    //   `F_VC_GONGSID` varchar(32) DEFAULT NULL COMMENT '公司id',
	FDtKaistjsj   string `gorm:"column:F_DT_KAISTJSJ"`   //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string `gorm:"column:F_DT_TONGJWCSJ"`  //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj    string `gorm:"column:F_DT_KUAIZSJ"`    //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//10 省内已发送数据统计表  `b_jsjk_shengnyfssjtj`
type BJsjkShengnyfssjtj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbZongje    int64  `gorm:"column:F_NB_ZONGJE"`    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int    `gorm:"column:F_NB_ZONGTS"`    //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj   string `gorm:"column:F_DT_KUAIZSJ"`   //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//11 省外结算趋势表  `b_jsjk_shengwjsqs`
type BJsjkShengwjsqs struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbJiaoye    int64  `gorm:"column:F_NB_JIAOYJE"`   //   `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
	FNbQingdje   int64  `gorm:"column:F_NB_QINGFJE"`   //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	FNbChae      int64  `gorm:"column:F_NB_CHAE"`      //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts   int    `gorm:"column:F_NB_JIAOYTS"`   //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingfts   int    `gorm:"column:F_NB_QINGFTS"`   //   `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq   string `gorm:"column:F_DT_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//12 省外结算数据分类  `b_jsjk_shengwjssjfl`
type BJsjkShengwjssjfl struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbJiaoyzts   int    `gorm:"column:F_NB_JIAOYZTS"`   //   `F_NB_JIAOYZTS` int DEFAULT NULL COMMENT '交易总条数',
	FNbQingfsjts  int    `gorm:"column:F_NB_QINGFSJTS"`  //   `F_NB_QINGFSJTS` int DEFAULT NULL COMMENT '清分数据条数',
	FNbJizsjts    int    `gorm:"column:F_NB_JIZSJTS"`    //   `F_NB_JIZSJTS` int DEFAULT NULL COMMENT '记账数据条数',
	FNbZhengysjts int    `gorm:"column:F_NB_ZHENGYSJTS"` //   `F_NB_ZHENGYSJTS` int DEFAULT NULL COMMENT '争议数据条数 待处理',
	FNbWeidbsjts  int    `gorm:"column:F_NB_WEIDBSJTS"`  //   `F_NB_WEIDBSJTS` int DEFAULT NULL COMMENT '未打包数据条数',
	FNbYidbsjts   int    `gorm:"column:F_NB_YIDBSJTS"`   //   `F_NB_YIDBSJTS` int DEFAULT NULL COMMENT '已打包数据条数',
	FNbYifssjts   int    `gorm:"column:F_NB_YIFSSJTS"`   //   `F_NB_YIFSSJTS` int DEFAULT NULL COMMENT '已发送数据条数',
	FNbHuaizsjts  int    `gorm:"column:F_NB_HUAIZSJTS"`  //   `F_NB_HUAIZSJTS` int DEFAULT NULL COMMENT '坏账数据条数',
	FDtKaistjsj   string `gorm:"column:F_DT_KAISTJSJ"`   //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string `gorm:"column:F_DT_TONGJWCSJ"`  //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtTongjrq    string `gorm:"column:F_DT_TONGJRQ"`    //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//13 省外结算争议数据统计表 `b_jsjk_shengwjszysjtj`
type BJsjkShengwjszysjtj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbZongje    int64     `gorm:"column:F_NB_ZONGJE"`    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int       `gorm:"column:F_NB_ZONGTS"`    //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  time.Time `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string    `gorm:"column:F_VC_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//14 省外结算清分统计表  `b_jsjk_shengwqftj`
type BJsjkShengwqftj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbZongje    int64     `gorm:"column:F_NB_ZONGJE"`    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int       `gorm:"column:F_NB_ZONGTS"`    //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  time.Time `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string    `gorm:"column:F_VC_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
	FNbHuaizje   int64     `gorm:"column:F_NB_HUAIZJE"`   //`F_NB_HUAIZJE` bigint DEFAULT NULL COMMENT '坏账金额',
	FNbHuaizts   int       `gorm:"column:F_NB_HUAIZTS"`   //`F_NB_HUAIZTS` bigint DEFAULT NULL COMMENT '坏账条数',
}

//15 省外停车场结算趋势表 `b_jsjk_shengwtccjsqs`
type BJsjkShengwtccjsqs struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbJiaoyje   int64  `gorm:"column:F_NB_JIAOYJE"`   //   `F_NB_JIAOYJE` bigint DEFAULT NULL COMMENT '交易金额',
	FNbQingfje   int64  `gorm:"column:F_NB_QINGFJE"`   //   `F_NB_QINGFJE` bigint DEFAULT NULL COMMENT '清分金额',
	FNbChae      int64  `gorm:"column:F_NB_CHAE"`      //   `F_NB_CHAE` bigint DEFAULT NULL COMMENT '差额',
	FNbJiaoyts   int    `gorm:"column:F_NB_JIAOYTS"`   //   `F_NB_JIAOYTS` int DEFAULT NULL COMMENT '交易条数',
	FNbQingfts   int    `gorm:"column:F_NB_QINGFTS"`   //   `F_NB_QINGFTS` int DEFAULT NULL COMMENT '清分条数',
	FVcGongsid   string `gorm:"column:F_VC_GONGSID"`   //   `F_VC_GONGSID` varchar(32) DEFAULT NULL COMMENT '公司id',
	FVcTingccid  string `gorm:"column:F_VC_TINGCCID"`  //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string `gorm:"column:F_VC_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',
}

//16 结算数据包监控表 `b_jsjk_shujbjk`
type BJsjkShujbjk struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbDabsl        int    `gorm:"column:F_NB_DABSL"`        //   `F_NB_DABSL` int DEFAULT NULL COMMENT '打包数量',
	FNbDabje        int    `gorm:"column:F_NB_DABJE"`        //   `F_NB_DABJE` bigint DEFAULT NULL COMMENT '打包金额',
	FNbFasysjybsl   int    `gorm:"column:F_NB_FASYSJYBSL"`   //   `F_NB_FASYSJYBSL` int DEFAULT NULL COMMENT '已发送原始交易消息包数量',
	FNbFasysjybje   int64  `gorm:"column:F_NB_FASYSJYBJE"`   //   `F_NB_FASYSJYBJE` bigint DEFAULT NULL COMMENT '已发送原始交易消息包金额',
	FNbJizbsl       int    `gorm:"column:F_NB_JIZBSL"`       //   `F_NB_JIZBSL` int DEFAULT NULL COMMENT '记账包数量',
	FNbJizbje       int64  `gorm:"column:F_NB_JIZBJE"`       //   `F_NB_JIZBJE` bigint DEFAULT NULL COMMENT '记账包金额',
	FNbYuansjyydbsl int    `gorm:"column:F_NB_YUANSJYYDBSL"` //   `F_NB_YUANSJYYDBSL` int DEFAULT NULL COMMENT '原始交易消息应答包数量',
	FDtKaistjsj     string `gorm:"column:F_DT_KAISTJSJ"`     //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj    string `gorm:"column:F_DT_TONGJWCSJ"`    //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj      string `gorm:"column:F_DT_KUAIZSJ"`      //   `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//17 数据同步监控表 `b_jsjk_shujtbjk`
type BJsjkShujtbjk struct {
	FNbId int `gorm:"AUTO_INCREMENT ;primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbJiessjzl  int    `gorm:"column:F_NB_JIESJZL"`   //   `F_NB_JIESJZL` int DEFAULT NULL COMMENT '结算数据总量',F_NB_JIESSJZL
	FNbYitbsjl   int    `gorm:"column:F_NB_YITBSJL"`   //   `F_NB_YITBSJL` int DEFAULT NULL COMMENT '已同步数据量',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string `gorm:"column:F_VC_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//18 停车场结算数据统计表 `b_jsjk_tingccjssjtj`
type BJsjkTingccjssjtj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FVcTingccid  string    `gorm:"column:F_VC_TINGCCID"`  //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FNbkawlh     int       `gorm:"column:F_NB_KAWLH"`     //   `F_NB_KAWLH` int DEFAULT NULL COMMENT '卡网络号',
	FNbZongje    int64     `gorm:"column:F_NB_ZONGJE"`    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int       `gorm:"column:F_NB_ZONGTS"`    //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FDtKaistjsj  time.Time `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string    `gorm:"column:F_VC_TONGJRQ"`   //   `F_VC_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//19 异常数据停车场统计表`b_jsjk_yicsjtcctj`
type BJsjkYicsjtcctj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbZongts    int    `gorm:"column:F_NB_ZONGTS"`    //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FNbZongje    int64  `gorm:"column:F_NB_ZONGJE"`    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额（分）',
	FVcTingccid  string `gorm:"column:F_NB_TINGCCID"`  //   `F_NB_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FNbTongjlx   int    `gorm:"column:F_NB_TONGJLX"`   //   `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:单点、2:总对总',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcKuaizsj   string `gorm:"column:F_VC_KUAIZSJ"`   //   `F_VC_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',

}

//20 结算待处理异常数据统计表 `b_jsjk_yicsjtj`
type BJsjkYicsjtj struct {
	FNbId int `gorm:"AUTO_INCREMENT; primary_key ;column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbZongje    int64     `gorm:"column:F_NB_ZONGJE"`    //   `F_NB_ZONGJE` bigint DEFAULT NULL COMMENT '总金额 （分）',
	FNbZongts    int       `gorm:"column:F_NB_ZONGTS"`    //   `F_NB_ZONGTS` int DEFAULT NULL COMMENT '总条数',
	FNbTongjlx   int       `gorm:"column:F_NB_TONGJLX"`   //   `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:总对总、2:单点',
	FDtKaistjsj  time.Time `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj time.Time `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string    `gorm:"column:F_VC_TONGJRQ"`   //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//21 逾期数据统计表 `b_jsjk_yuqsjtj`
type BJsjkYuqsjtj struct {
	FNbId int `gorm:"AUTO_INCREMENT ;primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbYuqzts    int    `gorm:"column:F_NB_YUQZTS"`    //   `F_NB_YUQZTS` int DEFAULT NULL COMMENT '逾期总条数',
	FNbYuqzje    int64  `gorm:"column:F_NB_YUQZJE"`    //   `F_NB_YUQZJE` bigint DEFAULT NULL COMMENT '逾期总金额 （分）',
	FVcTingccid  string `gorm:"column:F_VC_TINGCCID"`  //   `F_VC_TINGCCID` varchar(32) DEFAULT NULL COMMENT '停车场id',
	FDtKaistjsj  string `gorm:"column:F_DT_KAISTJSJ"`  //   `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj string `gorm:"column:F_DT_TONGJWCSJ"` //   `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FVcTongjrq   string `gorm:"column:F_ VC_TONGJRQ"`  //   `F_DT_TONGJRQ` date DEFAULT NULL COMMENT '统计日期',

}

//22 转结算数据监控表 `b_jsjk_zhuanjssjjk`
type BJsjkZhuanjssjjk struct {
	FNbId int `gorm:"AUTO_INCREMENT ;primary_key; column:F_NB_ID"` //  `F_NB_ID` int NOT NULL AUTO_INCREMENT COMMENT '唯一自增id',

	FNbChedyssjts int    `gorm:"column:F_NB_CHEDYSSJTS"` //  `F_NB_CHEDYSSJTS` int DEFAULT NULL COMMENT '车道原始数据条数',
	FNbJiesbsjts  int    `gorm:"column:F_NB_JIESBSJTS"`  //  `F_NB_JIESBSJTS` int DEFAULT NULL COMMENT '结算表数据条数',
	FNbTongjlx    int    `gorm:"column:F_NB_TONGJLX"`    //  `F_NB_TONGJLX` int DEFAULT NULL COMMENT '统计类型 1:单点、2:总对总',
	FDtKaistjsj   string `gorm:"column:F_DT_KAISTJSJ"`   //  `F_DT_KAISTJSJ` datetime DEFAULT NULL COMMENT '开始统计时间',
	FDtTongjwcsj  string `gorm:"column:F_DT_TONGJWCSJ"`  //  `F_DT_TONGJWCSJ` datetime DEFAULT NULL COMMENT '统计完成时间',
	FDtKuaizsj    string `gorm:"column:F_DT_KUAIZSJ"`    //  `F_DT_KUAIZSJ` datetime DEFAULT NULL COMMENT '快照时间',
}

//23 用户表 b_jsjk_jiesjkptyhb
type BJsjkJiesjkptyhb struct {
	FNbYonghid string `gorm:"column:F_NB_YONGHID"` //  `F_NB_YONGHID` varchar(32) DEFAULT NULL COMMENT '用户id', '用户id',//手机号 F_NB_YONGHID
	FVcYonghmm string `gorm:"column:F_VC_YONGHMM"` //  `F_VC_YONGHMM` varchar(32) DEFAULT NULL COMMENT '用户密码', '用户密码',
	FVcShoujh  string `gorm:"column:F_VC_Shoujh"`  //  `F_VC_Shoujh` varchar(32) DEFAULT NULL COMMENT '手机号', '手机号',
	FVcYoux    string `gorm:"column:F_VC_YOUX"`    //  `F_VC_YOUX` varchar(32) DEFAULT NULL COMMENT '邮箱', '邮箱',
	FVcYonghnc string `gorm:"column:F_VC_YONGHNC"` //  `F_VC_YONGHNC` varchar(32) DEFAULT NULL COMMENT '用户昵称''用户昵称',
}

// 24   B_JS_QINGFTJXX【清分统计消息】b_js_qingftjxx
type BJsQingftjxx struct {
	FVcBanbh         string    `gorm:"column:F_VC_BANBH"`                              //F_VC_BANBH	版本号	VARCHAR(32)
	FNbXiaoxlb       int       `gorm:"column:F_NB_XIAOXLB"`                            //F_NB_XIAOXLB	消息类别	INT
	FNbXiaoxlx       int       `gorm:"column:F_NB_XIAOXLX"`                            //F_NB_XIAOXLX	消息类型	INT
	FVcFaszid        string    `gorm:"column:F_VC_FASZID"`                             //F_VC_FASZID	发送者ID	VARCHAR(32)
	FVcJieszid       string    `gorm:"column:F_VC_JIESZID"`                            //F_VC_JIESZID	接收者ID	VARCHAR(32)
	FNbXiaoxxh       int64     `gorm:"column:F_NB_XIAOXXH"`                            //F_NB_XIAOXXH	消息序号	BIGINT
	FDtJiessj        time.Time `gorm:"column:F_DT_JIESSJ"`                             //F_DT_JIESSJ	接收时间	DATETIME
	FVcQingfmbr      string    `gorm:"column:F_VC_QINGFMBR"`                           //F_VC_QINGFMBR	清分目标日	DATE
	FNbQingfzje      int       `gorm:"column:F_NB_QINGFZJE"`                           //F_NB_QINGFZJE	清分总金额	INT
	FNbQingfsl       int       `gorm:"column:F_NB_QINGFSL"`                            //F_NB_QINGFSL	清分数量	INT
	FDtQingftjclsj   time.Time `gorm:"column:F_DT_QINGFTJCLSJ"`                        //F_DT_QINGFTJCLSJ	清分统计处理时间	DATETIME
	FNbYuansjysl     int       `gorm:"column:F_NB_YUANSJYSL"`                          //F_NB_YUANSJYSL	原始包交易数量	INT
	FNbZhengycljgbsl int       `gorm:"column:F_NB_ZHENGYCLJGBSL"`                      //F_NB_ZHENGYCLJGBSL	争议处理结果包数量	INT
	FVcXiaoxwjlj     string    `gorm:"column:F_VC_XIAOXWJLJ"`                          //F_VC_XIAOXWJLJ	消息文件路径	VARCHAR(512)
	FDtChulsj        time.Time `gorm:"column:F_DT_CHULSJ"`                             //`F_DT_CHULSJ` datetime DEFAULT NULL COMMENT '处理时间',
	FNbWeiyid        int       `gorm:"AUTO_INCREMENT; primary_key;column:F_NB_WEIYID"` //`F_NB_WEIYID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
}

// 25  B_JS_ZHENGYCLXX【争议交易处理消息】b_js_zhengyjyclxx
type BJsZhengyjyclxx struct {
	FVcBanbh        string    `gorm:"column:F_VC_BANBH"`                              //F_VC_BANBH	版本号	VARCHAR(32)
	FNbXiaoxlb      int       `gorm:"column:F_NB_XIAOXLB"`                            //F_NB_XIAOXLB	消息类别	INT
	FNbXiaoxlx      int       `gorm:"column:F_NB_XIAOXLX"`                            //F_NB_XIAOXLX	消息类型	INT
	FVcFaszid       string    `gorm:"column:F_VC_FASZID"`                             //F_VC_FASZID	发送者ID	VARCHAR(32)
	FVcJieszid      string    `gorm:"column:F_VC_JIESZID"`                            //F_VC_JIESZID	接收者ID	VARCHAR(32)
	FNbXiaoxxh      int64     `gorm:"column:F_NB_XIAOXXH"`                            //F_NB_XIAOXXH	消息序号	BIGINT
	FDtJiessj       time.Time `gorm:"column:F_DT_JIESSJ"`                             //F_DT_JIESSJ	接收时间	DATETIME
	FVcQingffid     string    `gorm:"column:F_VC_QINGFFID"`                           //F_VC_QINGFFID	清分方ID	VARCHAR(32)
	FVcLianwzxid    string    `gorm:"column:F_VC_LIANWZXID"`                          //F_VC_LIANWZXID	联网中心ID	VARCHAR(32)
	FVcFaxfid       string    `gorm:"column:F_VC_FAXFID"`                             //F_VC_FAXFID	发行方ID	VARCHAR(32)
	FVcZhengyjgwjid int       `gorm:"column:F_VC_ZHENGYJGWJID"`                       //F_VC_ZHENGYJGWJID	争议结果文件ID	INT
	FDtZhengyclsj   time.Time `gorm:"column:F_DT_ZHENGYCLSJ"`                         //F_DT_ZHENGYCLSJ	争议处理时间	DATETIME
	FNbZhengysl     int       `gorm:"column:F_NB_ZHENGYSL"`                           //F_NB_ZHENGYSL	争议数量	INT
	FNbQuerxyjzdzje int       `gorm:"column:F_NB_QUERXYJZDZJE"`                       //F_NB_QUERXYJZDZJE	确认需要记账的总金额	INT
	FNbZhixjg       int       `gorm:"column:F_NB_ZHIXJG"`                             //F_NB_ZHIXJG` int DEFAULT NULL COMMENT '执行结果 1：消息已正常接收（用于Advice Response时含已接受建议）、2：消息头错误，如MessageClass或MessageType不符合定义，SenderId不存在等、3：消息格式不正确，即XML Schema验证未通过、4：消息格式正确但内容错误，包括数量不符，内容重复等、5：消息重复、6：消息正常接收，但不接受建议（仅用于Advice Response）、7：消息版本错误',
	FVcXiaoxwjlj    string    `gorm:"column:F_VC_XIAOXWJLJ"`                          //F_VC_XIAOXWJLJ	消息文件路径	VARCHAR(512)
	FNbWeiyid       int       `gorm:"AUTO_INCREMENT; primary_key;column:F_NB_WEIYID"` //`F_NB_WEIYID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
}

// 26    B_JS_ZHENGYJYCLMX【争议交易处理明细】
type BJsZhengyjyclmx struct {
	FNbZhengyjyclxxxh int64 //F_NB_ZHENGYJYCLXXXH	争议交易处理消息序号	BIGINT
	FNbFenzxh         int   //F_NB_FENZXH	分组序号	INT
	FNbYuansjyxxxh    int64 //F_NB_YUANSJYXXXH	原始交易消息序号	BIGINT
	FNbZunjlsl        int   //F_NB_ZUNJLSL	组内记录数量	INT
	FNbZunjezh        int   //F_NB_ZUNJEZH	组内金额总和	INT
	FNbYuansbnxh      int   //F_NB_YUANSBNXH	原始包内序号	INT
	FNbChuljg         int   //F_NB_CHULJG	处理结果	INT
	FNbWeiyid         int   //`F_NB_WEIYID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
}

// 27 B_JS_JIZCLXX【记账处理消息】
type BJsJizclxx struct {
	FVcBanbh       string    //F_VC_BANBH	版本号	VARCHAR(32)
	FNbXiaoxlb     int       //F_NB_XIAOXLB	消息类别	INT
	FNbXiaoxlx     int       //F_NB_XIAOXLX	消息类型	INT
	FVcFaszid      string    //F_VC_FASZID	发送者ID	VARCHAR(32)
	FVcJieszid     string    //F_VC_JIESZID	接收者ID	VARCHAR(32)
	FNbXiaoxxh     int64     //F_NB_XIAOXXH	消息序号	BIGINT
	FDtJiessj      string    //F_DT_JIESSJ	接收时间	DATETIME
	FNbYuansjyxxxh int64     //F_NB_YUANSJYXXXH	原始交易消息序号	BIGINT
	FNbJilsl       int       //F_NB_JILSL	记录数量	INT
	FNbZongje      int64     //F_NB_ZONGJE	总金额	INT
	FNbZhengysl    int       //F_NB_ZHENGYSL	争议数量	INT
	FNbZhixjg      int       //F_NB_ZHIXJG	执行结果	INT
	FDtChulsj      time.Time //F_DT_CHULSJ	处理时间	DATETIME
	FVcXiaoxwjlj   string    //F_VC_XIAOXWJLJ	消息文件路径	VARCHAR(512)
	FNbWeiyid      int       //`F_NB_WEIYID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
}

//29   B_JS_JIZCLMX【记账处理明细】
type BJsJizclmx struct {
	FNbYuansjyxxxh int64 //F_NB_YUANSJYXXXH	原始交易消息序号	BIGINT
	FNbBaonxh      int   //F_NB_BAONXH	包内序号	INT
	FNbChuljg      int   //F_NB_CHULJG	处理结果	INT
	FNbWeiyid      int   //`F_NB_WEIYID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
}

//    B_JS_QINGFTONGJIMX【清分统计明细】b_js_qingftjmx
type BJsQingftjmx struct {
	FNbQingftjxxxh    int64  //F_NB_QINGFTJXXXH	清分统计消息序号	BIGINT
	FNbFenzxh         int    //F_NB_FENZXH	分组序号	INT
	FVcTongxbzxxtid   string //F_VC_TONGXBZXXTID	通行宝中心系统ID	VARCHAR(32)
	FNbYuansjyxxxh    int64  //F_NB_YUANSJYXXXH	原始交易消息序号	BIGINT
	FNbZhengycljgwjid int    //F_NB_ZHENGYCLJGWJID	争议处理结果文件ID	INT
	FNbWeiyid         int    //`F_NB_WEIYID` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
}

type Result struct {
	Total     int64
	Count     int
	Parkingid string
}

type ClearlingAndDispute struct {
	DataType  string
	DateTime  string
	PackageNo string
}

type ClearlingAndDisputeData struct {
	ClearPacgNo    string
	Cleardatetime  string
	DisputPacgeNo  string
	Disputdatetime string
}
