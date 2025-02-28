// Code generated by MockGen. DO NOT EDIT.
// Source: service/customer_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	web "github.com/aronipurwanto/go-restful-api/model/web"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerService) Create(ctx context.Context, request web.CustomerCreateRequest) (web.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(web.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomerServiceMockRecorder) Create(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockCustomerService) Delete(ctx context.Context, categoryId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, categoryId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerServiceMockRecorder) Delete(ctx, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerService)(nil).Delete), ctx, categoryId)
}

// FindAll mocks base method.
func (m *MockCustomerService) FindAll(ctx context.Context) ([]web.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]web.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCustomerServiceMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCustomerService)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockCustomerService) FindById(ctx context.Context, categoryId uint64) (web.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, categoryId)
	ret0, _ := ret[0].(web.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCustomerServiceMockRecorder) FindById(ctx, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCustomerService)(nil).FindById), ctx, categoryId)
}

// Update mocks base method.
func (m *MockCustomerService) Update(ctx context.Context, request web.CustomerUpdateRequest) (web.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(web.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomerServiceMockRecorder) Update(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerService)(nil).Update), ctx, request)
}
