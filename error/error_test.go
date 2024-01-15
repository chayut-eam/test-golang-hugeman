package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Detail struct {
	Field       string
	Description string
}

func TestNewserverError(t *testing.T) {
	errmsg := "Invalid data"
	detail := Detail{
		Field:       "ID",
		Description: "is empty",
	}
	want := DefinedError{
		BaseError{
			Code:    "500",
			Message: errmsg,
			Detail:  detail,
		},
	}
	t.Run("test success", func(t *testing.T) {
		result := NewDefinedError(errmsg, detail, "500")
		assert.Equal(t, want, result)
	})
}

func TestNewFieldValidationError(t *testing.T) {
	errmsg := "Invalid data"
	want := DefinedError{
		BaseError{
			Code:    "400",
			Message: "Field Validation Error",
			Detail:  errmsg,
		},
	}
	t.Run("test success", func(t *testing.T) {
		result := NewFieldValidationError(errmsg)
		assert.Equal(t, want, result)
		assert.Equal(t, want.BaseError.Error(), result.BaseError.Error())
	})
}

func TestNewErrorResponse(t *testing.T) {
	errmsg := "Invalid data"
	testError := DefinedError{
		BaseError{
			Code:    "400",
			Message: errmsg,
			Detail:  nil,
		},
	}

	testError2 := BaseError{
		Code:    "400",
		Message: errmsg,
		Detail:  nil,
	}
	resultError1 := NewErrorResponse(testError)

	assert.NotEmpty(t, resultError1)

	resultError2 := NewErrorResponse(testError2)

	assert.NotEmpty(t, resultError2)

}
