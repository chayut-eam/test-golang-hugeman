package mock

import (
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func NewMockService() *mockService {
	return &mockService{}
}

func (m *mockService) GetAll(logger *logrus.Entry) (*model.Response, error) {
	args := m.Called()
	return args.Get(0).(*model.Response), args.Error(1)
}
func (m *mockService) Create(logger *logrus.Entry, req *model.Data) (*model.Response, error) {
	args := m.Called()
	return args.Get(0).(*model.Response), args.Error(1)
}
func (m *mockService) SearchByTilteOrDes(logger *logrus.Entry, id string) (*model.Response, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Response), args.Error(1)
}
func (m *mockService) UpdateById(logger *logrus.Entry, id string, req *model.Data) (*model.Response, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Response), args.Error(1)
}
