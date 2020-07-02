package dto

type Response struct {
	Code    int         `json:"code"  example:"200"`
	Message string      `json:"message" example:"响应成功信息"`
	Data    interface{} `json:"data"`
}
type ResponseFailure struct {
	Code    int         `json:"code"  example:"404"`
	Message string      `json:"message" example:"响应失败信息"`
	Data    interface{} `json:"data"`
}

type Reqlogin struct {
	UserName string `json:"name" example:"abc"`
}
