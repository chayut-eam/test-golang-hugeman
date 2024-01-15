package error

import (
	"github.com/chayut-eam/test-golang-hugeman/utils"
)

type ErrorResponse struct {
	TimeStamp    string      `json:"timestamp"`
	ErrorCode    string      `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	ErrorDetail  interface{} `json:"error_details,omitempty"`
}

func NewErrorResponse(err error) ErrorResponse {
	if definedError, ok := err.(DefinedError); ok {
		baseError := definedError.BaseError
		return ErrorResponse{
			TimeStamp:    utils.Now(),
			ErrorCode:    baseError.Code,
			ErrorMessage: baseError.Error(),
			ErrorDetail:  baseError.Detail,
		}
	} else {
		return ErrorResponse{
			TimeStamp:    utils.Now(),
			ErrorCode:    "500",
			ErrorMessage: err.Error(),
		}
	}
}
