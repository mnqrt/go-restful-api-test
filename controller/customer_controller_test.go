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

type CustomerControllerTestSuite struct {
	ctrl            *gomock.Controller
	mockService     *mocks.MockCustomerService
	app             *fiber.App
}

func setupTest() *CustomerControllerTestSuite {
	ctrl := gomock.NewController(nil)
	mockService := mocks.NewMockCustomerService(ctrl)
	controller := NewCustomerController(mockService)

	app := fiber.New()
	api := app.Group("/api")
	customers := api.Group("/customers")
	customers.Post("/", controller.Create)
	customers.Put("/:customerId", controller.Update)
	customers.Delete("/:customerId", controller.Delete)
	customers.Get("/:customerId", controller.FindById)
	customers.Get("/", controller.FindAll)

	return &CustomerControllerTestSuite{ctrl, mockService, app}
}

func TestCreateCustomer(t *testing.T) {
	suite := setupTest()
	defer suite.ctrl.Finish()

	requestBody := web.CustomerCreateRequest{Name: "John Doe"}
	requestJSON, _ := json.Marshal(requestBody)
	suite.mockService.EXPECT().Create(gomock.Any(), requestBody).Return(web.CustomerResponse{Id: 1, Name: "John Doe"}, nil)

	req := httptest.NewRequest(http.MethodPost, "/api/customers/", bytes.NewReader(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := suite.app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateCustomer(t *testing.T) {
	suite := setupTest()
	defer suite.ctrl.Finish()

	requestBody := web.CustomerUpdateRequest{Id: 1, Name: "Jane Doe"}
	requestJSON, _ := json.Marshal(requestBody)
	suite.mockService.EXPECT().Update(gomock.Any(), requestBody).Return(web.CustomerResponse{Id: 1, Name: "Jane Doe"}, nil)

	req := httptest.NewRequest(http.MethodPut, "/api/customers/1", bytes.NewReader(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := suite.app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteCustomer(t *testing.T) {
	suite := setupTest()
	defer suite.ctrl.Finish()

	suite.mockService.EXPECT().Delete(gomock.Any(), uint64(1)).Return(nil)
	req := httptest.NewRequest(http.MethodDelete, "/api/customers/1", nil)
	resp, _ := suite.app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestFindCustomerById(t *testing.T) {
	suite := setupTest()
	defer suite.ctrl.Finish()

	suite.mockService.EXPECT().FindById(gomock.Any(), uint64(1)).Return(web.CustomerResponse{Id: 1, Name: "John Doe"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/customers/1", nil)
	resp, _ := suite.app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestFindAllCustomers(t *testing.T) {
	suite := setupTest()
	defer suite.ctrl.Finish()

	suite.mockService.EXPECT().FindAll(gomock.Any()).Return([]web.CustomerResponse{{Id: 1, Name: "John Doe"}}, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/customers/", nil)
	resp, _ := suite.app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}