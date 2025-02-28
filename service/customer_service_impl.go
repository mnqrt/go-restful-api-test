package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/exception"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomerServiceImpl struct {
	CustomerRepo repository.CustomerRepository
	Validator    *validator.Validate
}

func NewCustomerService(customerRepo repository.CustomerRepository, validator *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepo: customerRepo,
		Validator:    validator,
	}
}

// Create - Adds a new customer to the database
func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) (web.CustomerResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.CustomerResponse{}, err
	}

	customer := domain.Customer{Name: request.Name}
	savedCustomer, err := service.CustomerRepo.Save(ctx, customer)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(savedCustomer), nil
}

// Update - Modifies an existing customer's details
func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) (web.CustomerResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.CustomerResponse{}, err
	}

	customer, err := service.CustomerRepo.FindById(ctx, request.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CustomerResponse{}, exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return web.CustomerResponse{}, err
	}

	customer.Name = request.Name
	updatedCustomer, err := service.CustomerRepo.Update(ctx, customer)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(updatedCustomer), nil
}

// Delete - Removes a customer from the database
func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId uint64) error {
	customer, err := service.CustomerRepo.FindById(ctx, customerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return err
	}

	return service.CustomerRepo.Delete(ctx, customer)
}

// FindById - Retrieves a customer using their ID
func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId uint64) (web.CustomerResponse, error) {
	customer, err := service.CustomerRepo.FindById(ctx, customerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CustomerResponse{}, exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(customer), nil
}

// FindAll - Retrieves all customers from the database
func (service *CustomerServiceImpl) FindAll(ctx context.Context) ([]web.CustomerResponse, error) {
	customers, err := service.CustomerRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToCustomerResponses(customers), nil
}
