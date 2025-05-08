package v1

import (
	"encoding/json"
	"shopping/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Data:   nil,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
