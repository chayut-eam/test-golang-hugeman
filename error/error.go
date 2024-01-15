package error

type (
	BaseError struct {
		Code    string
		Message string
		Detail  interface{}
	}

	DefinedError struct {
		BaseError
	}
)

func (e BaseError) Error() string {
	return e.Message
}

func NewDefinedError(errmsg string, detail interface{}, code string) DefinedError {
	return DefinedError{
		BaseError{
			Code:    code,
			Message: errmsg,
			Detail:  detail,
		},
	}
}

func NewFieldValidationError(detail interface{}) DefinedError {
	return DefinedError{
		BaseError{
			Code:    "400",
			Message: "Field Validation Error",
			Detail:  detail,
		},
	}
}
