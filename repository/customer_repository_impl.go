package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db: db}
}

// Save - Insert a new customer into the database
func (repository *CustomerRepositoryImpl) Save(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	if err := repository.db.WithContext(ctx).Create(&customer).Error; err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

// Update - Modify an existing customer record
func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	if err := repository.db.WithContext(ctx).Save(&customer).Error; err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

// Delete - Remove a customer entry from the database
func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, customer domain.Customer) error {
	if err := repository.db.WithContext(ctx).Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Retrieve a customer by ID
func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, customerId uint64) (domain.Customer, error) {
	var customer domain.Customer
	err := repository.db.WithContext(ctx).First(&customer, customerId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customer, errors.New("customer not found")
	}
	return customer, err
}

// FindAll - Fetch all customer records
func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) ([]domain.Customer, error) {
	var customers []domain.Customer
	err := repository.db.WithContext(ctx).Find(&customers).Error
	return customers, err
}
