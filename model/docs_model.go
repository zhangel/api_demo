package model

type ListOutPut struct {
	Code    int32        `json:"code" example:"200"`
	Message string       `json:"message" example:"OK"`
	Data    []SampleInfo `json:"data" example:[]SampleInfo`
}

type HttpError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"服务器内部错误"`
}
