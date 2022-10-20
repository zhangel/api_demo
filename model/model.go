package model

type JsonOut struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"OK"`
	Data    interface{} `json:"data"`
}

type ServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"服务器内部错误"`
	Data    bool   `json:"data" example:"false"`
}

type ExecSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
	Data    bool   `json:"data" example:"true"`
}
