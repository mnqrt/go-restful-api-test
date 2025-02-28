package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCustomerRepository(ctrl)
	ctx := context.Background()
	sampleCustomer := domain.Customer{CustomerID: 1, Name: "John Doe", Email: "john.doe@example.com", Phone: "987654", Address: "123 Street"}
	modifiedCustomer := domain.Customer{CustomerID: 1, Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "112233", Address: "456 Avenue"}
	anotherCustomer := domain.Customer{CustomerID: 2, Name: "Alice Smith", Email: "alice.smith@example.com", Phone: "334455", Address: "789 Road"}

	testCases := []struct {
		testName   string
		mockSetup func()
		operation func() (interface{}, error)
		expected  interface{}
		expectErr bool
	}{
		{
			testName: "Successful Save",
			mockSetup: func() {
				repo.EXPECT().Save(ctx, sampleCustomer).Return(sampleCustomer, nil)
			},
			operation: func() (interface{}, error) {
				return repo.Save(ctx, sampleCustomer)
			},
			expected:  sampleCustomer,
			expectErr: false,
		},
		{
			testName: "Failed Save",
			mockSetup: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Customer{}, errors.New("error saving"))
			},
			operation: func() (interface{}, error) {
				return repo.Save(ctx, anotherCustomer)
			},
			expected:  domain.Customer{},
			expectErr: true,
		},
		{
			testName: "Successful Update",
			mockSetup: func() {
				repo.EXPECT().Update(ctx, sampleCustomer).Return(modifiedCustomer, nil)
			},
			operation: func() (interface{}, error) {
				return repo.Update(ctx, sampleCustomer)
			},
			expected:  modifiedCustomer,
			expectErr: false,
		},
		{
			testName: "Retrieve Customer by ID",
			mockSetup: func() {
				repo.EXPECT().FindById(ctx, uint64(1)).Return(sampleCustomer, nil)
			},
			operation: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expected:  sampleCustomer,
			expectErr: false,
		},
		{
			testName: "Customer Not Found",
			mockSetup: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Customer{}, errors.New("not found"))
			},
			operation: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expected:  domain.Customer{},
			expectErr: true,
		},
		{
			testName: "Retrieve All Customers",
			mockSetup: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Customer{sampleCustomer, anotherCustomer}, nil)
			},
			operation: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expected:  []domain.Customer{sampleCustomer, anotherCustomer},
			expectErr: false,
		},
		{
			testName: "Successful Deletion",
			mockSetup: func() {
				repo.EXPECT().Delete(ctx, sampleCustomer).Return(nil)
			},
			operation: func() (interface{}, error) {
				return nil, repo.Delete(ctx, sampleCustomer)
			},
			expected:  nil,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.mockSetup()
			result, err := tc.operation()

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
