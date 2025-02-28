package controller

import (
	"bytes"
	"encoding/json"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initializeTestCategoryApp(mockService *mocks.MockCategoryService) *fiber.App {
	app := fiber.New()
	categoryCtrl := NewCategoryController(mockService)

	api := app.Group("/api")
	categories := api.Group("/categories")
	categories.Post("/", categoryCtrl.Create)
	categories.Put("/:categoryId", categoryCtrl.Update)
	categories.Delete("/:categoryId", categoryCtrl.Delete)
	categories.Get("/:categoryId", categoryCtrl.FindById)
	categories.Get("/", categoryCtrl.FindAll)

	return app
}

func TestCategoryHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockCategoryService(ctrl)
	app := initializeTestCategoryApp(mockService)

	testCases := []struct {
		testName       string
		requestMethod  string
		requestURL     string
		requestBody    interface{}
		mockSetup      func()
		expectedStatus int
		expectedResult web.WebResponse
	}{
		{
			testName:      "Successful category update",
			requestMethod: "PUT",
			requestURL:    "/api/categories/1",
			requestBody:   web.CategoryUpdateRequest{Id: 1, Name: "Updated"},
			mockSetup: func() {
				mockService.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(web.CategoryResponse{Id: 1, Name: "Updated"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.CategoryResponse{Id: 1, Name: "Updated"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.mockSetup()

			var body []byte
			if tc.requestBody != nil {
				body, _ = json.Marshal(tc.requestBody)
			}

			req := httptest.NewRequest(tc.requestMethod, tc.requestURL, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			var respData web.WebResponse
			json.NewDecoder(resp.Body).Decode(&respData)

			if parsedData, ok := respData.Data.(map[string]interface{}); ok {
				respData.Data = web.CategoryResponse{
					Id:   uint64(parsedData["id"].(float64)),
					Name: parsedData["name"].(string),
				}
			}

			assert.Equal(t, tc.expectedResult, respData)
		})
	}
}
