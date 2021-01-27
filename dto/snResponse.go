package dto

type QueryResponse struct {
	Code    int `json:"code"  example:"200"`
	CodeMsg string
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}

type RealTimeSettlementData struct {
	Shengnjssl   int    `json:"num"  example:"123"`               //省内结算数量
	Shengnjsjine string `json:"money"  example:"123"`             //省内结算金额
	Fassl        int    `json:"send_num"  example:"123"`          //已发送 数量
	Fasjine      string `json:"send_money"  example:"123"`        //已发送 金额
	Jizsl        int    `json:"keepaccount_num"  example:"123"`   //记账数量
	Jizjine      string `json:"keepaccount_money"  example:"123"` //记账金额
	DateTime     string `json:"datetime"  example:"123"`          //完成时间
}

type SNClearandJiesuan struct {
	JiesuanMoney   string `json:"total_money"  example:"123"`     //交易总金额
	ClearlingMoney string `json:"clearl_money"  example:"123"`    //清分总金额
	DiffMoney      string `json:"diff_money"  example:"123"`      //差额
	DiffCount      int    `json:"diff_count"  example:"123"`      //差数
	JiesuanCount   int    `json:"count"  example:"123"`           //交易结算条数
	ClearlingCount int    `json:"clear_count"  example:"123"`     //清分总笔数
	DateTime       string `json:"datetime"  example:"2020-08-18"` //完成时间

}

type DataSync struct {
	JiesuanCount int    `json:"sync_count"  example:"123"`      //结算表
	HailCount    int    `json:"source_count"  example:"123"`    //海岭
	DateTime     string `json:"datetime"  example:"2020-08-18"` //完成时间

}

type ShengNDataClassification struct {
	Shengnzcount int    `json:"count"`                          //省内结算总数据
	Yiqkcount    int    `json:"clear_count"`                    //已清分总条数（含坏账）
	Weifscount   int    `json:"no_send_count"`                  //未打包
	Yifscount    int    `json:"send_count"`                     //已发送
	Jufuzcount   int    `json:"bad_debts_count"`                //拒付数量
	DateTime     string `json:"datetime"  example:"2020-08-18"` //完成时间
	ShengnzMoney string `json:"money"`                          //省内结算总金额元
	YiqkMoney    string `json:"clear_money"`                    //已清分总金额元
	JufuzMoney   string `json:"bad_debts_money"`                //拒付金额元

}

type Overduedata struct {
	Overduecount  int    `json:"overdue_count"  example:"123"`  //逾期数量
	Parkingname   string `json:"parking_name"  example:"南京南站"`  //停车场名称
	OverdueAmount string `json:"overdue_amount"  example:"123"` //逾期金额
}

//AbnormalDataOfParking
type AbnormalDataOfParking struct {
	AbnormalDatacount  int    `json:"abnormal_data_count"  example:"123"`
	Parkingname        string `json:"parking_name"  example:"南京南站"` //停车场名称
	AbnormalDataAmount string `json:"abnormal_data_amount"  example:"123"`
}

type SyncRequest struct {
	TradeTimeStart string `json:"tradeTimeStart"  example:"123"`
	SyncStatus     int    `json:"syncStatus"  example:"123"`
}

type SyncResponse struct {
	Code     int      `json:"code"  example:"123"` //	"code": 0,
	Desc     string   `json:"desc"  example:"123"` //	"desc": "SUCCESS",
	SyncData SyncData `json:"data"  example:"123"` // "data": {

}
type SyncData struct {
	Count     int       `json:"count"  example:"123"` // "count": 5130,
	Sum       int64     `json:"sum"  example:"123"`   // "sum": 7007310,
	SyncParam SyncParam `json:"param"  example:"123"` //"param": {
}

type SyncParam struct {
	TradeTimeStart string `json:"tradeTimeStart"  example:"123"`
	SyncStatus     int    `json:"syncStatus"  example:"123"`
	ParkNo         int64  `json:"parkNo"  example:"123"`     //停车场编号
	ExitlaneNo     int    `json:"exitlaneNo"  example:"123"` //车道编号
	TradeTimeEnd   string `json:"tradeTimeEnd"  example:"123"`
}
