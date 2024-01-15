package model_test

import (
	"testing"

	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	give := model.Response{
		TimeStamp: utils.Now(),
		Code:      200,
		Data:      "test",
		Message:   "test",
	}

	want := model.NewResponse(200, "test", "test")

	assert.Equal(t, give, want)
}
