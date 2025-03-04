// Code generated by MockGen. DO NOT EDIT.
// Source: repository/customer_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/aronipurwanto/go-restful-api/model/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCustomerRepository) Delete(ctx context.Context, customer domain.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerRepositoryMockRecorder) Delete(ctx, customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerRepository)(nil).Delete), ctx, customer)
}

// FindAll mocks base method.
func (m *MockCustomerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCustomerRepositoryMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCustomerRepository)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockCustomerRepository) FindById(ctx context.Context, customerId uint64) (domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, customerId)
	ret0, _ := ret[0].(domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCustomerRepositoryMockRecorder) FindById(ctx, customerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCustomerRepository)(nil).FindById), ctx, customerId)
}

// Save mocks base method.
func (m *MockCustomerRepository) Save(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, customer)
	ret0, _ := ret[0].(domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockCustomerRepositoryMockRecorder) Save(ctx, customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCustomerRepository)(nil).Save), ctx, customer)
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, customer)
	ret0, _ := ret[0].(domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(ctx, customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), ctx, customer)
}
