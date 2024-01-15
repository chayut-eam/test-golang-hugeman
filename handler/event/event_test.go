package event_test

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	ers "github.com/chayut-eam/test-golang-hugeman/error"
	"github.com/chayut-eam/test-golang-hugeman/handler/event"
	logs "github.com/chayut-eam/test-golang-hugeman/logger"
	"github.com/chayut-eam/test-golang-hugeman/mock"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/utils"
	"github.com/chayut-eam/test-golang-hugeman/validation"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

type InMemoryFile struct {
	*bytes.Buffer
}

// Implement the Close method for *os.File compatibility
func (f *InMemoryFile) Close() error {
	return nil
}
func TestGetAllHandler(t *testing.T) {
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

	testCases := []struct {
		name string
		data *[]model.Data
		code int
		er   error
	}{
		{
			name: "success",
			data: &[]model.Data{
				{
					ID:          "1223344545-qwe2231312432",
					Title:       "test",
					Description: "test",
					CreatedAt:   "2024-01-15T09:30:00.123Z",
					Image:       "YXNkd2FzZHdhc2R3YXM=",
					Status:      "IN_PROGRESS",
				},
			},
			code: 200,
			er:   nil,
		},
		{
			name: "error",
			data: nil,
			code: 500,
			er:   errors.New("error"),
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			res := model.NewResponse(testcase.code, testcase.data, "")
			mockService := mock.NewMockService()
			mockService.On("GetAll").Return(&res, testcase.er)
			handler := event.NewHelloHandlerImpl(mockService, logs)

			gin.SetMode(gin.TestMode)

			router := gin.New()
			router.GET("/data", handler.GetAllHandler)

			req, err := http.NewRequest("GET", "/data", nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, recorder.Code, testcase.code)

		})
	}
}
func TestCreateHandler(t *testing.T) {
	validation.Init()
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

	testCases := []struct {
		name string
		data *model.Data
		code int
		file string
		er   error
	}{
		{
			name: "success",
			data: &model.Data{
				ID:          "1223344545-qwe2231312432",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS",
			},
			code: 200,
			file: "test.jpg",
			er:   nil,
		},
		{
			name: "error",
			data: &model.Data{
				ID:          "error",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS",
			},
			code: 500,
			file: "test.jpg",
			er:   errors.New("error"),
		},
		{
			name: "field validation",
			data: &model.Data{
				ID:          "",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS22",
			},
			code: 400,
			file: "test.jpg",
			er:   nil,
		},
		{
			name: "read file error",
			data: &model.Data{
				ID:          "t21312321321aprowep_test",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS",
			},
			code: 400,
			file: "xxxx.jpg",
			er:   nil,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			res := model.NewResponse(testcase.code, &testcase.data, "")
			mockService := mock.NewMockService()
			mockService.On("Create").Return(&res, testcase.er)
			handler := event.NewHelloHandlerImpl(mockService, logs)

			gin.SetMode(gin.TestMode)

			router := gin.New()
			router.POST("/create", handler.CreateHandler)

			bodyReader := &bytes.Buffer{}

			reqData := testcase.data
			writer := multipart.NewWriter(bodyReader)
			writer.WriteField("id", reqData.ID)
			writer.WriteField("title", reqData.Title)
			writer.WriteField("description", reqData.Description)
			writer.WriteField("created_at", reqData.CreatedAt)
			writer.WriteField("status", reqData.Status)
			currentPath, _ := utils.GetRelativePath()
			imgPath := filepath.Join(currentPath, "utils", testcase.file)
			if testcase.name != "read file error" {
				part, err := writer.CreateFormFile("image", imgPath)
				if err != nil {
					panic(err)
				}

				var buf bytes.Buffer
				data := make([]byte, 1)
				_, err = buf.Write(data)
				if err != nil {
					panic(err)
				}

				file := &InMemoryFile{Buffer: &buf}
				_, err = io.Copy(part, file)
				if err != nil {
					panic(err)
				}

				writer.Close()
			}

			req, err := http.NewRequest("POST", "/create", bodyReader)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, recorder.Code, testcase.code)

		})
	}
}

func TestSearchHandler(t *testing.T) {
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

	testCases := []struct {
		name      string
		searchKey string
		code      int
		er        error
	}{
		{
			name:      "success",
			searchKey: "test",
			code:      200,
			er:        nil,
		},
		{
			name:      "error",
			searchKey: "error",
			code:      500,
			er:        errors.New("error"),
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			res := model.NewResponse(testcase.code, testcase.searchKey, "")
			mockService := mock.NewMockService()
			mockService.On("SearchByTilteOrDes", testcase.searchKey).Return(&res, testcase.er)
			handler := event.NewHelloHandlerImpl(mockService, logs)

			gin.SetMode(gin.TestMode)

			router := gin.New()
			router.GET("/data/:search_value", handler.SearchHandler)

			req, err := http.NewRequest("GET", "/data/"+testcase.searchKey, nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, recorder.Code, testcase.code)

		})
	}
}

func TestUpdateHandler(t *testing.T) {
	validation.Init()
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

	testCases := []struct {
		name string
		data *model.Data
		code int
		er   error
	}{
		{
			name: "success",
			data: &model.Data{
				ID:          "1223344545-qwe2231312432",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS",
			},
			code: 200,
			er:   nil,
		},
		{
			name: "error",
			data: &model.Data{
				ID:          "error",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS",
			},
			code: 500,
			er:   errors.New("error"),
		},
		{
			name: "defined error",
			data: &model.Data{
				ID:          "error",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS",
			},
			code: 400,
			er:   ers.NewDefinedError("test", nil, "400"),
		},
		{
			name: "field validation",
			data: &model.Data{
				ID:          "test",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS22",
			},
			code: 400,
			er:   nil,
		},
		{
			name: "read file error",
			data: &model.Data{
				ID:          "test",
				Title:       "test",
				Description: "test",
				CreatedAt:   "2024-01-15T09:30:00.123Z",
				Image:       "YXNkd2FzZHdhc2R3YXM=",
				Status:      "IN_PROGRESS22",
			},
			code: 400,
			er:   nil,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			res := model.NewResponse(testcase.code, &testcase.data, "")
			mockService := mock.NewMockService()
			mockService.On("UpdateById", testcase.data.ID).Return(&res, testcase.er)
			handler := event.NewHelloHandlerImpl(mockService, logs)

			gin.SetMode(gin.TestMode)

			router := gin.New()
			router.PATCH("/update/:id", handler.UpdateHandler)

			bodyReader := &bytes.Buffer{}

			reqData := testcase.data
			writer := multipart.NewWriter(bodyReader)
			writer.WriteField("id", reqData.ID)
			writer.WriteField("title", reqData.Title)
			writer.WriteField("description", reqData.Description)
			writer.WriteField("created_at", reqData.CreatedAt)
			writer.WriteField("status", reqData.Status)
			currentPath, _ := utils.GetRelativePath()
			imgPath := filepath.Join(currentPath, "utils", "test.jpg")
			if testcase.name != "read file error" {
				part, err := writer.CreateFormFile("image", imgPath)
				if err != nil {
					panic(err)
				}

				var buf bytes.Buffer
				data := make([]byte, 1)
				_, err = buf.Write(data)
				if err != nil {
					panic(err)
				}

				file := &InMemoryFile{Buffer: &buf}
				_, err = io.Copy(part, file)
				if err != nil {
					panic(err)
				}

				writer.Close()
			}

			req, err := http.NewRequest("PATCH", "/update/"+reqData.ID, bodyReader)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, recorder.Code, testcase.code)

		})
	}
}
