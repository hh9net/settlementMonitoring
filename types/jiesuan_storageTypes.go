package types

import "time"

//  B_JS_JIESSJ【结算数据】`b_js_jiessj`
type BJsJiessj struct {
	FVcJiaoyjlid   string    `xorm:"pk"` //F_VC_JIAOYJLID	交易记录ID	VARCHAR(128)
	FVcTingccbh    string    //F_VC_TINGCCBH	停车场编号	VARCHAR(32)
	FVcChedid      string    //F_VC_CHEDID	车道ID	VARCHAR(32)
	FVcGongsjtid   string    //F_VC_GONGSJTID	公司/集团ID	VARCHAR(32)
	FNbTingcclx    int       //F_NB_TINGCCLX	停车场类型	INT 1单点，2总对总
	FNbYuansjybxh  int64     //F_NB_YUANSJYBXH	原始交易包序号	BIGINT
	FNbJiaoybnxh   int       //F_NB_JIAOYBNXH	交易包内序号	INT
	FNbJizjg       int       //F_NB_JIZJG	记账结果	INT "0：未记账  1：已记账    2：争议数据"
	FNbZhengylx    int       //F_NB_ZHENGYLX	争议类型	INT 0，不是争议，1，验证未通过
	FNbJizbxh      int64     //F_NB_JIZBXH	记账包序号	INT
	FNbZhengyclbxh int64     //F_NB_ZHENGYCLBXH	争议处理包序号	BIGINT  记账结果：争议放过、坏账时
	FNbQingfbxh    int64     //F_NB_QINGFBXH	清分包序号	BIGINT
	FVcXiaofjlbh   string    //F_VC_XIAOFJLBH	消费记录编号	VARCHAR(128)
	FVcJiamkh      string    //F_VC_JIAMKH	加密卡号	VARCHAR(32)终端
	FVcKajmjyxh    string    //F_VC_KAJMJYXH	加密卡交易序号	VARCHAR(32)终端
	FVcObuid       string    //F_VC_OBUID	Obuid	VARCHAR(32)
	FVcObufxf      string    //F_VC_OBUFXF	obu发行方	VARCHAR(32)
	FVcObucp       string    //F_VC_OBUCP	obu内车牌	VARCHAR(32)
	FVcObucpys     string    //F_VC_OBUCPYS	obu车牌颜色	VARCHAR(32)
	FVcKah         string    //F_VC_KAH	    卡号	VARCHAR(32)
	FVcKawlh       string    //F_VC_KAWLH	卡网络号	VARCHAR(32)
	FVcKajyxh      string    //F_VC_KAJYXH	卡交易序号	VARCHAR(32)
	FVcKafxf       string    //F_VC_KAFXF	卡发行方	VARCHAR(32)
	FNbKalx        int       //F_NB_KALX	卡类型	INT  储值卡22，23 记账卡
	FVcCheph       string    // F_VC_CHEPH   卡内车牌号
	FNbJiaoyqye    int64     //F_NB_JIAOYQYE	交易前余额	分转元 INT
	FNbJiaoyhye    int64     //F_NB_JIAOYHYE	交易后余额	分转元 INT
	FNbJine        int64     //F_NB_JINE	金额	INT         分转元
	FVcTacm        string    //F_VC_TACM	TAC码	VARCHAR(32)
	FDtJiaoysj     time.Time //F_DT_JIAOYSJ	交易时间	DATETIME   2020-05-13 14:34:34
	FDtJiaoylx     string    //F_DT_JIAOYLX	交易类型	VARCHAR(32)
	FVcChex        string    //F_VC_CHEX	车型	VARCHAR(32)
	FVcObuzt       string    //F_VC_OBUZT	OBu状态	VARCHAR(32)
	FVcSuanfbs     string    //F_VC_SUANFBS	算法标识	VARCHAR(32)     【交易标识】
	FDtYonghrksj   time.Time //F_DT_YONGHRKSJ	用户入口时间	DATETIME
	FNbYonghtcsc   int       //F_NB_YONGHTCSC	用户停车时长(分)	INT  天时分秒
	FVcZhangdms    string    //F_VC_ZHANGDMS	账单描述（给用户通知的信息）	VARCHAR(32)
	FVcMiybbh      string    //F_VC_MIYBBH	密钥版本号	VARCHAR(32)
	FVcObuyyxlh    string    //F_VC_OBUYYXLH	obu应用序列号	VARCHAR(32)
	FVcChedjyxh    string    //F_VC_CHDJYXH	车道交易序号	VARCHAR(32)
	FNbQingfjg     int       //F_NB_QINGFJG  '清分结果 0：未清分、1：已清分'
	FNbDabzt       int       //F_NB_DABZT	打包状态	INT   0 初始 ；1打包中； 2已打包
	FNbZhengycljg  int       //F_NB_ZHENGYCLJG  '争议处理结果 0:未处理、1：争议放过、2：坏账'
	FNbJusbj       int       //`F_NB_JUSBJ`   DEFAULT '0' COMMENT '拒收标记 0、正常，1、拒收',
	FVcQingfmbr    string    // `F_NB_QINGFMBR` int(11) DEFAULT NULL COMMENT '清分目标日',
}
