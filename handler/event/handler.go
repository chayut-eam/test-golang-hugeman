package event

import (
	"errors"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	er "github.com/chayut-eam/test-golang-hugeman/error"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/service"
	"github.com/chayut-eam/test-golang-hugeman/utils"
	v "github.com/chayut-eam/test-golang-hugeman/validation"

	"github.com/sirupsen/logrus"
)

// interface
type Handler interface {
	GetAllHandler(c *gin.Context)
	CreateHandler(c *gin.Context)
	SearchHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}

// concrete implementation
type HandlerImpl struct {
	logger  *logrus.Entry
	service service.Service
}

// constructor
func NewHelloHandlerImpl(service service.Service, logger *logrus.Entry) *HandlerImpl {
	return &HandlerImpl{
		logger:  logger,
		service: service,
	}
}

func (h *HandlerImpl) GetAllHandler(c *gin.Context) {

	logs := h.logger.WithFields(logrus.Fields{
		"route": c.FullPath(),
	})

	logs.Info("Start Get All")

	res, err := h.service.GetAll(logs)
	if err != nil {
		c.JSON(codeError(err), er.NewErrorResponse(err))
		return
	}

	c.JSON(200, res)
}

func (h *HandlerImpl) CreateHandler(c *gin.Context) {

	imgToBase64, err := convertImgToBase64(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, er.NewErrorResponse(err))
		return
	}

	req := model.Data{
		ID:          c.Request.FormValue("id"),
		Title:       c.Request.FormValue("title"),
		Description: c.Request.FormValue("description"),
		CreatedAt:   c.Request.FormValue("created_at"),
		Image:       *imgToBase64,
		Status:      c.Request.FormValue("status"),
	}

	err = validateRequest(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, er.NewErrorResponse(err))
		return
	}

	logs := h.logger.WithFields(logrus.Fields{
		"route": c.FullPath(),
	})

	logs.Info("Start Create")

	res, err := h.service.Create(logs, &req)
	if err != nil {
		c.JSON(codeError(err), er.NewErrorResponse(err))
		return
	}

	c.JSON(200, res)
}

func (h *HandlerImpl) SearchHandler(c *gin.Context) {

	searchKey := c.Param("search_value")
	logs := h.logger.WithFields(logrus.Fields{
		"route": c.FullPath(),
	})

	logs.Info("Search data")

	res, err := h.service.SearchByTilteOrDes(logs, searchKey)
	if err != nil {
		c.JSON(codeError(err), er.NewErrorResponse(err))
		return
	}

	c.JSON(200, res)
}

func (h *HandlerImpl) UpdateHandler(c *gin.Context) {

	id := c.Param("id")
	imgToBase64, err := convertImgToBase64(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, er.NewErrorResponse(err))
		return
	}

	req := model.Data{
		ID:          id,
		Title:       c.Request.FormValue("title"),
		Description: c.Request.FormValue("description"),
		CreatedAt:   c.Request.FormValue("created_at"),
		Image:       *imgToBase64,
		Status:      c.Request.FormValue("status"),
	}

	err = validateRequest(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	logs := h.logger.WithFields(logrus.Fields{
		"route": c.FullPath(),
	})

	logs.Info("Update data")

	res, err := h.service.UpdateById(logs, id, &req)
	if err != nil {
		c.JSON(codeError(err), er.NewErrorResponse(err))
		return
	}

	c.JSON(200, res)
}

func validateRequest(c *gin.Context, req model.Data) error {

	// validate
	if err := v.Validate(&req); err != nil {
		return err
	}

	return nil
}

func convertImgToBase64(c *gin.Context) (*string, error) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, errors.New("Unable to parse form data. Cause by limit of request size")
	}

	// Get the image file from the form-data
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		return nil, errors.New("Unable to get file from form data")
	}
	defer file.Close()

	imgToBase64, err := utils.ConvertImgToBase64(file)
	if err != nil {
		return nil, errors.New("Unable to convert image to base64")
	}

	return imgToBase64, nil
}

func codeError(err error) int {
	status := http.StatusInternalServerError
	if definedError, ok := err.(er.DefinedError); ok {
		status, _ = strconv.Atoi(definedError.Code)
	}

	return status
}
