package model

import (
	"github.com/chayut-eam/test-golang-hugeman/utils"
)

type Response struct {
	TimeStamp string      `json:"timestamp"`
	Code      int         `json:"error_code"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"error_message,omitempty"`
}

type Data struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required"`
	CreatedAt   string `json:"created_at" validate:"required"`
	Image       string `json:"image"`
	Status      string `json:"status" validate:"required,oneof=IN_PROGRESS COMPLETED"`
}

func NewResponse(code int, data interface{}, msg string) Response {
	return Response{
		TimeStamp: utils.Now(),
		Code:      code,
		Data:      data,
		Message:   msg,
	}
}
