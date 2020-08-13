package dto

type QueryResponse struct {
	Code    int `json:"code"  example:"200"`
	CodeMsg string
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}

type RealTimeSettlementData struct {
	Shengnjssl   int   `json:"shengnjssl"  example:"123"`   //省内结算数量
	Shengnjsjine int64 `json:"shengnjsjine"  example:"123"` //省内结算金额
	Fassl        int   `json:"fassl"  example:"123"`        //已发送 数量
	Fasjine      int64 `json:"fasjine"  example:"123"`      //已发送 金额
	Jizsl        int   `json:"jizsl"  example:"123"`        //记账数量
	Jizjine      int64 `json:"jizjine"  example:"123"`      //记账金额
}

type SNClearandJiesuan struct {
	JiesuanMoney   int64 `json:"jiesuanmoney"  example:"123"`   //交易总金额
	ClearlingMoney int64 `json:"clearlingmoney"  example:"123"` //清分总金额
	DiffMoney      int64 `json:"diffmoney"  example:"123"`      //清分总金额
	JiesuanCount   int   `json:"jiesuancount"  example:"123"`   //交易结算条数
	ClearlingCount int   `json:"clearlingcount"  example:"123"` //清分总笔数
}

type DataSync struct {
	JiesuanCount int `json:"jiesuancount"  example:"123"` //交易结算条数
	HailCount    int `json:"hailcount"  example:"123"`    //清分总笔数
}

type ShengNDataClassification struct {
	Shengnzcount int `json:"shengnzcount"  example:"123"` //结算总数据
	Yiqkcount    int `json:"yiqkcount"  example:"123"`    //已清分总条数（不含坏账）
	Weifscount   int `json:"weifscount"  example:"123"`   //未打包
	Yifscount    int `json:"yifscount"  example:"123"`    //已发送
	Jufuzcount   int `json:"jufuzcount"  example:"123"`   //坏账
}
