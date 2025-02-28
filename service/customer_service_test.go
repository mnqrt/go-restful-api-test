package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testCustomer        = domain.Customer{Name: "Bob Marley", Email: "bob.marley@example.com", Phone: "5551234", Address: "Reggae Street"}
	updatedCustomer     = domain.Customer{Name: "Bob Smith", Email: "bob.smith@example.com", Phone: "5555678", Address: "Downtown Avenue"}
	testCustomerRequest = web.CustomerCreateRequest{Name: "Bob Marley", Email: "bob.marley@example.com", Phone: "5551234", Address: "Reggae Street"}
	testCustomerUpdate  = web.CustomerUpdateRequest{Name: "Bob Smith", Email: "bob.smith@example.com", Phone: "5555678", Address: "Downtown Avenue"}
)

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	mockValidator := validator.New()
	customerService := NewCustomerService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.CustomerCreateRequest
		mock      func()
		expect    web.CustomerResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: testCustomerRequest,
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(testCustomer, nil)
			},
			expect:    helper.ToCustomerResponse(testCustomer),
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.CustomerCreateRequest{Name: "", Email: "", Phone: "", Address: ""},
			mock:      func() {},
			expect:    web.CustomerResponse{},
			expectErr: true,
		},
		{
			name:  "repository error",
			input: testCustomerRequest,
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Customer{}, errors.New("database error"))
			},
			expect:    web.CustomerResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := customerService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}
