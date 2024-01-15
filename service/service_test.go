package service_test

import (
	"errors"
	"testing"

	"github.com/chayut-eam/test-golang-hugeman/mock"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/service"

	logs "github.com/chayut-eam/test-golang-hugeman/logger"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	config := model.AppConfig{
		AppInfo: model.AppInfo{
			Name: "test",
		},
		LoggerConfig: model.LoggerConfig{
			LogLevel: "info",
		},
	}
	logs.Init(config.AppInfo, config.LoggerConfig)
	logs := logs.LoggerSystem()

	testcases := []struct {
		name         string
		expectResult struct {
			resultData *[]model.Data
			Error      error
		}
	}{
		{
			name: "Exist data",
			expectResult: struct {
				resultData *[]model.Data
				Error      error
			}{
				resultData: &[]model.Data{
					{
						ID:          "1223344545-qwe2231312432",
						Title:       "test",
						Description: "test",
						CreatedAt:   "2024-01-15T09:30:00.123Z",
						Image:       "YXNkd2FzZHdhc2R3YXM=",
						Status:      "IN_PROGRESS",
					},
				},
				Error: nil,
			},
		},
		{
			name: "Empty data",
			expectResult: struct {
				resultData *[]model.Data
				Error      error
			}{
				resultData: &[]model.Data{},
				Error:      nil,
			},
		},
		{
			name: "Error data",
			expectResult: struct {
				resultData *[]model.Data
				Error      error
			}{
				resultData: nil,
				Error:      errors.New("Error Data"),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result := testcase.expectResult.resultData
			newMockRepository := mock.NewMockRepository()
			newMockRepository.On("GetAll").Return(testcase.expectResult.resultData, testcase.expectResult.Error)
			mockService := service.NewService(newMockRepository)
			res, err := mockService.GetAll(logs)
			if err != nil {
				assert.Error(t, err)
			} else {
				if len(*result) == 0 {
					result = nil
				}
				assert.Equal(t, res.Data, result)
			}
		})
	}
}

func TestCreate(t *testing.T) {

	config := model.AppConfig{
		AppInfo: model.AppInfo{
			Name: "test",
		},
		LoggerConfig: model.LoggerConfig{
			LogLevel: "info",
		},
	}
	logs.Init(config.AppInfo, config.LoggerConfig)
	logs := logs.LoggerSystem()

	testcases := []struct {
		name         string
		req          *model.Data
		expectResult struct {
			Error error
		}
	}{
		{
			name: "Create Success",
			req: &model.Data{
				ID:        "123312321312-31287sadsa",
				Title:     "test",
				CreatedAt: "2024-01-15T09:30:00.123Z",
				Image:     "YXNkd2FzZHdhc2R3YXM=",
				Status:    "IN_PROGRESS",
			},
			expectResult: struct {
				Error error
			}{
				Error: nil,
			},
		},
		{
			name: "Empty Failed. cause by exist data",
			req: &model.Data{
				ID:        "999999-31287sadsa",
				Title:     "exist data",
				CreatedAt: "2024-01-15T09:30:00.123Z",
				Image:     "YXNkd2FzZHdhc2R3YXM=",
				Status:    "IN_PROGRESS",
			},
			expectResult: struct {
				Error error
			}{
				Error: errors.New("exist data"),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			newMockRepository := mock.NewMockRepository()
			newMockRepository.On("Create", testcase.req.ID).Return(testcase.expectResult.Error)
			mockService := service.NewService(newMockRepository)
			res, err := mockService.Create(logs, testcase.req)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, res.Code, 200)
			}
		})
	}
}

func TestSearchById(t *testing.T) {

	config := model.AppConfig{
		AppInfo: model.AppInfo{
			Name: "test",
		},
		LoggerConfig: model.LoggerConfig{
			LogLevel: "info",
		},
	}
	logs.Init(config.AppInfo, config.LoggerConfig)
	logs := logs.LoggerSystem()

	testcases := []struct {
		name         string
		id           string
		expectResult struct {
			Data  *[]model.Data
			Error error
		}
	}{
		{
			name: "Found",
			id:   "123312321312-31287sadsa",
			expectResult: struct {
				Data  *[]model.Data
				Error error
			}{
				Data: &[]model.Data{
					{
						ID:        "123312321312-31287sadsa",
						Title:     "exist data",
						CreatedAt: "2024-01-15T09:30:00.123Z",
						Image:     "YXNkd2FzZHdhc2R3YXM=",
						Status:    "IN_PROGRESS",
					},
				},
				Error: nil,
			},
		},
		{
			name: "not found",
			id:   "11111-notfound",
			expectResult: struct {
				Data  *[]model.Data
				Error error
			}{
				Data:  nil,
				Error: errors.New("not found"),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			newMockRepository := mock.NewMockRepository()
			newMockRepository.On("SearchByTilteOrDes", testcase.id).Return(testcase.expectResult.Data, testcase.expectResult.Error)
			mockService := service.NewService(newMockRepository)
			res, err := mockService.SearchByTilteOrDes(logs, testcase.id)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, res.Code, 200)
			}
		})
	}
}

func TestUpdateById(t *testing.T) {

	config := model.AppConfig{
		AppInfo: model.AppInfo{
			Name: "test",
		},
		LoggerConfig: model.LoggerConfig{
			LogLevel: "info",
		},
	}
	logs.Init(config.AppInfo, config.LoggerConfig)
	logs := logs.LoggerSystem()

	testcases := []struct {
		name         string
		id           string
		Data         *model.Data
		expectResult struct {
			Error error
		}
	}{
		{
			name: "Update Sucess",
			id:   "123312321312-31287sadsa",
			Data: &model.Data{
				Title:     "Sucess",
				CreatedAt: "2024-01-15T09:30:00.123Z",
				Image:     "YXNkd2FzZHdhc2R3YXM=",
				Status:    "IN_PROGRESS",
			},
			expectResult: struct {
				Error error
			}{
				Error: nil,
			},
		},
		{
			name: "Update Failed",
			id:   "11111-failed",
			Data: &model.Data{
				Title:     "failed",
				CreatedAt: "2024-01-15T09:30:00.123Z",
				Image:     "YXNkd2FzZHdhc2R3YXM=",
				Status:    "IN_PROGRESS",
			},
			expectResult: struct {
				Error error
			}{
				Error: errors.New("not found"),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			newMockRepository := mock.NewMockRepository()
			newMockRepository.On("UpdateById", testcase.id).Return(testcase.expectResult.Error)
			mockService := service.NewService(newMockRepository)
			res, err := mockService.UpdateById(logs, testcase.id, testcase.Data)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, res.Code, 200)
			}
		})
	}
}
