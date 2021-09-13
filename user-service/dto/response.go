package dto

import "net/http"

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code"`
}

func Success(data interface{}) Response {
	return Response{
		Data:    data,
		Ok:      true,
		Message: "",
		Code:    http.StatusOK,
	}
}

func Error(err error, code int) Response {
	return Response{
		Data:    nil,
		Ok:      false,
		Message: err.Error(),
		Code:    code,
	}
}
