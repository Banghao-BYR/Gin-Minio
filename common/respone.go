package common

type BaseResponse struct {
	StatusCode int32       `json:"code"`
	StatusMsg  string      `json:"msg"`
	Data       interface{} `json:"data"`
}
