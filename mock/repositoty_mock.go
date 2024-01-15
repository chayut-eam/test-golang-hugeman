package mock

import (
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func NewMockRepository() *mockRepository {
	return &mockRepository{}
}

func (m *mockRepository) GetAll() (*[]model.Data, error) {
	args := m.Called()
	return args.Get(0).(*[]model.Data), args.Error(1)
}
func (m *mockRepository) Create(req *model.Data) error {
	args := m.Called(req.ID)
	return args.Error(0)
}
func (m *mockRepository) SearchByTilteOrDes(searchKey string) (*[]model.Data, error) {
	args := m.Called(searchKey)
	return args.Get(0).(*[]model.Data), args.Error(1)
}
func (m *mockRepository) UpdateById(id string, req *model.Data) error {
	args := m.Called(id)
	return args.Error(0)
}
