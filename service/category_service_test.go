package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepository(ctrl)
	validator := validator.New()
	service := NewCategoryService(mockRepo, validator)

	testCases := []struct {
		testName  string
		input     web.CategoryCreateRequest
		mockSetup func()
		expected  web.CategoryResponse
		expectErr bool
	}{
		{
			testName: "Successful Creation",
			input:    web.CategoryCreateRequest{Name: "Electronics"},
			mockSetup: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Category{Id: 1, Name: "Electronics"}, nil)
			},
			expected:  web.CategoryResponse{Id: 1, Name: "Electronics"},
			expectErr: false,
		},
		{
			testName: "Validation Failure",
			input:    web.CategoryCreateRequest{Name: ""},
			mockSetup: func() {},
			expected:  web.CategoryResponse{},
			expectErr: true,
		},
		{
			testName: "Repository Error",
			input:    web.CategoryCreateRequest{Name: "Toys"},
			mockSetup: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Category{}, errors.New("database error"))
			},
			expected:  web.CategoryResponse{},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.mockSetup()
			resp, err := service.Create(context.Background(), tc.input)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, resp)
			}
		})
	}
}
