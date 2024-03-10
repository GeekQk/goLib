package main

import (
	"errors"
	"fmt"
)

// Response 定义了一个通用的接口返回结构
type Response struct {
	Data interface{} `json:"data"` // 数据字段，可以是任意类型
	Code int         `json:"code"` // 状态码字段
	Msg  string      `json:"msg"`  // 消息字段
}

// NewSuccessResponse 创建一个成功的响应
func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Data: data,
		Code: 0,
		Msg:  "Success",
	}
}

// NewErrorResponse 创建一个错误响应
func NewErrorResponse(err error, code int) *Response {
	return &Response{
		Data: nil,
		Code: code,
		Msg:  err.Error(),
	}
}

// 示例接口函数
func SomeFunction(param string) *Response {
	// 假设这是你的业务逻辑
	if param == "error" {
		// 返回一个错误响应
		return NewErrorResponse(errors.New("something went wrong"), 400)
	}

	// 假设业务逻辑成功
	// 返回一个成功的响应
	return NewSuccessResponse("Operation successful")
}

func main() {
	// 调用接口
	resp := SomeFunction("errors")

	// 处理响应
	if resp.Code == 0 {
		// 处理成功的情况
		fmt.Println("Success:", resp.Code, resp.Msg, resp.Data)
	} else {
		// 处理错误的情况
		fmt.Println("Error:", resp.Code, resp.Msg, resp.Data)
	}
}
