// Code generated by MockGen. DO NOT EDIT.
// Source: internal/informations/usecase.go
//
// Generated by this command:
//
//	mockgen -source=internal/informations/usecase.go -destination=internal/informations/mock/usecase_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/esgi-challenge/backend/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(user *models.User, informations *models.Informations) (*models.Informations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user, informations)
	ret0, _ := ret[0].(*models.Informations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(user, informations any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), user, informations)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(user *models.User, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", user, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), user, id)
}

// GetAll mocks base method.
func (m *MockUseCase) GetAll(user *models.User) (*[]models.Informations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", user)
	ret0, _ := ret[0].(*[]models.Informations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUseCaseMockRecorder) GetAll(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUseCase)(nil).GetAll), user)
}

// GetById mocks base method.
func (m *MockUseCase) GetById(user *models.User, id uint) (*models.Informations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", user, id)
	ret0, _ := ret[0].(*models.Informations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUseCaseMockRecorder) GetById(user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUseCase)(nil).GetById), user, id)
}
