package response

import "github.com/luyasr/mpush/common/errs"

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func New(data any) *Response {
	return &Response{
		Code:    0,
		Data:    data,
		Message: "success",
	}
}

func NewWithError(err error) *Response {
	return &Response{
		Code:    errs.GetCode(err),
		Data:    nil,
		Message: errs.GetMessage(err),
	}
}
