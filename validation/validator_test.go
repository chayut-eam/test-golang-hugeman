package validation_test

import (
	"encoding/json"
	"testing"

	"github.com/chayut-eam/test-golang-hugeman/validation"

	"github.com/stretchr/testify/assert"
)

type Request struct {
	Name string `json:"name,omitempty" validate:"required,notEmpty,max=10"`
	Test string `json:"-"`
}

func TestValidator(t *testing.T) {
	validation.Init()
	req_test := `{
		"name": ""
	}`
	req := Request{}
	json.Unmarshal([]byte(req_test), &req)
	err := validation.Validate(req)
	assert.Error(t, err)

	req_test = `{
		"name": "564565qsdsadwasdwasdwa"
	}`
	req = Request{}
	json.Unmarshal([]byte(req_test), &req)
	err = validation.Validate(req)
	assert.Error(t, err)

	req_test = `{
		"name": "564565"
	}`
	req = Request{}
	json.Unmarshal([]byte(req_test), &req)
	err = validation.Validate(req)
	assert.NoError(t, err)

	unbufferedChannel := make(chan int)
	err = validation.Validate(unbufferedChannel)
	assert.Error(t, err)
}
