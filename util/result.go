package util

import "encoding/json"

// Result 通用返回格式
type Result struct {
	Code int         `json:"code" form:"code"`
	Msg  string      `json:"msg" form:"msg"`
	Data interface{} `json:"data" form:"data"`
}

type tmp struct {
	Item  interface{} `json:"item" form:"item"`
	Count int         `json:"count" form:"count"`
}

const (
	Success int = 200
	Fail    int = 300
)

// NewResult 返回通用结果
func NewResult(code int, msg string, data ...interface{}) []byte {
	if len(data) > 0 {
		res := Result{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}

		md, _ := json.Marshal(res)
		return md
	} else {
		res := Result{
			Code: code,
			Msg:  msg,
		}

		md, _ := json.Marshal(res)
		return md

	}
}

func NewPage(code int, msg string, data interface{}, count int) []byte {
	res := Result{
		Code: code,
		Msg:  msg,
		Data: tmp{
			Item:  data,
			Count: count,
		},
	}

	md, _ := json.Marshal(res)
	return md
}
