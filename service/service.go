package service

import (
	"fmt"

	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/repository"

	"github.com/sirupsen/logrus"
)

type Service interface {
	GetAll(logger *logrus.Entry) (*model.Response, error)
	Create(logger *logrus.Entry, req *model.Data) (*model.Response, error)
	SearchByTilteOrDes(logger *logrus.Entry, searchKey string) (*model.Response, error)
	UpdateById(logger *logrus.Entry, id string, req *model.Data) (*model.Response, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(logger *logrus.Entry) (*model.Response, error) {

	data, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	if len(*data) <= 0 {
		data = nil
	}
	logger.Infof("Get All Success")
	res := model.NewResponse(200, data, "")
	return &res, nil
}

func (s *service) Create(logger *logrus.Entry, req *model.Data) (*model.Response, error) {

	err := s.repository.Create(req)
	if err != nil {
		return nil, err
	}
	logger.Infof("Create Success")
	res := model.NewResponse(200, nil, fmt.Sprintf("'%s' create success!", req.Title))
	return &res, nil
}

func (s *service) SearchByTilteOrDes(logger *logrus.Entry, searchValue string) (*model.Response, error) {

	data, err := s.repository.SearchByTilteOrDes(searchValue)
	if err != nil {
		return nil, err
	}
	logger.Infof("Search data")
	res := model.NewResponse(200, data, "")
	return &res, nil
}


func (s *service) UpdateById(logger *logrus.Entry, id string, req *model.Data) (*model.Response, error) {

	err := s.repository.UpdateById(id, req)
	if err != nil {
		return nil, err
	}
	logger.Infof("Update data")
	res := model.NewResponse(200, fmt.Sprintf("'%s' update success!", id), "")
	return &res, nil
}
