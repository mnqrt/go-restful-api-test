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

func TestCategoryRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCategoryRepository(ctrl)
	ctx := context.Background()

	testCases := []struct {
		testName  string
		mockSetup func()
		method    func() (interface{}, error)
		expected  interface{}
		expectErr bool
	}{
		{
			testName: "Successful Save",
			mockSetup: func() {
				category := domain.Category{Id: 1, Name: "Electronics"}
				repo.EXPECT().Save(ctx, category).Return(category, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Category{Id: 1, Name: "Electronics"})
			},
			expected:  domain.Category{Id: 1, Name: "Electronics"},
			expectErr: false,
		},
		{
			testName: "Failed Save",
			mockSetup: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Category{}, errors.New("failed to save"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Category{Name: "Invalid"})
			},
			expected:  domain.Category{},
			expectErr: true,
		},
		{
			testName: "Successful Update",
			mockSetup: func() {
				category := domain.Category{Id: 1, Name: "Updated Name"}
				repo.EXPECT().Update(ctx, category).Return(category, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Category{Id: 1, Name: "Updated Name"})
			},
			expected:  domain.Category{Id: 1, Name: "Updated Name"},
			expectErr: false,
		},
		{
			testName: "Retrieve By Id Successfully",
			mockSetup: func() {
				repo.EXPECT().FindById(ctx, 1).Return(domain.Category{Id: 1, Name: "Electronics"}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expected:  domain.Category{Id: 1, Name: "Electronics"},
			expectErr: false,
		},
		{
			testName: "Retrieve By Id Not Found",
			mockSetup: func() {
				repo.EXPECT().FindById(ctx, gomock.Any()).Return(domain.Category{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expected:  domain.Category{},
			expectErr: true,
		},
		{
			testName: "Retrieve All Categories Successfully",
			mockSetup: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Category{{Id: 1, Name: "Electronics"}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expected:  []domain.Category{{Id: 1, Name: "Electronics"}},
			expectErr: false,
		},
		{
			testName: "Successful Deletion",
			mockSetup: func() {
				repo.EXPECT().Delete(ctx, domain.Category{Id: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Category{Id: 1})
			},
			expected:  nil,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.mockSetup()
			result, err := tc.method()

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
