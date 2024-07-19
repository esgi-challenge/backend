// Code generated by MockGen. DO NOT EDIT.
// Source: internal/chat/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/chat/repository.go -destination=internal/chat/mock/repository_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/esgi-challenge/backend/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(chat *models.Channel) (*models.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", chat)
	ret0, _ := ret[0].(*models.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(chat any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), chat)
}

// GetAllByUser mocks base method.
func (m *MockRepository) GetAllByUser(userId uint) (*[]models.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUser", userId)
	ret0, _ := ret[0].(*[]models.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUser indicates an expected call of GetAllByUser.
func (mr *MockRepositoryMockRecorder) GetAllByUser(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUser", reflect.TypeOf((*MockRepository)(nil).GetAllByUser), userId)
}

// GetById mocks base method.
func (m *MockRepository) GetById(id uint) (*models.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRepositoryMockRecorder) GetById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), id)
}

// SaveMessage mocks base method.
func (m *MockRepository) SaveMessage(msg *models.Message) (*models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveMessage", msg)
	ret0, _ := ret[0].(*models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveMessage indicates an expected call of SaveMessage.
func (mr *MockRepositoryMockRecorder) SaveMessage(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveMessage", reflect.TypeOf((*MockRepository)(nil).SaveMessage), msg)
}