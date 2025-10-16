package models

// 统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // 有数据时才显示
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func Error(code int, message string) *Response {
	return &Response{
		Code: code,
		Msg:  message,
	}
}

func NotFound(message string) *Response {
	return Error(404, message)
}

func BadRequest(message string) *Response {
	return Error(400, message)
}

func InternalError() *Response {
	return Error(500, "服务器内部错误")
}
