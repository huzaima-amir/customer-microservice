package biz

import (
	"context"
	"errors"
)

//  entities

type Customer struct {
	ID          int64
	Name        string
	DateOfBirth string
}

type Email struct {
	ID         int64
	CustomerID int64
	Email      string
}

type PhoneNumber struct {
	ID         int64
	CustomerID int64
	Phone      string
}

type Address struct {
	ID         int64
	CustomerID int64
	Address    string
}

//  repository interface 

type CustomerRepo interface {
	// customer
	CreateCustomer(ctx context.Context, c *Customer) error
	UpdateCustomer(ctx context.Context, c *Customer) error
	DeleteCustomer(ctx context.Context, id int64) error
	GetCustomer(ctx context.Context, id int64) (*Customer, error)
	ListCustomers(ctx context.Context) ([]*Customer, error)

	// email
	CreateEmail(ctx context.Context, e *Email) error
	DeleteEmail(ctx context.Context, customerID int64, email string) error
	ListEmails(ctx context.Context, customerID int64) ([]string, error)
	GetCustomerByEmail(ctx context.Context, email string) (*Customer, error)

	// phone
	CreatePhone(ctx context.Context, p *PhoneNumber) error
	DeletePhone(ctx context.Context, customerID int64, phone string) error
	ListPhones(ctx context.Context, customerID int64) ([]string, error)
	GetCustomerByPhone(ctx context.Context, phone string) (*Customer, error)

	// address
	CreateAddress(ctx context.Context, a *Address) error
	DeleteAddress(ctx context.Context, customerID int64, address string) error
	ListAddresses(ctx context.Context, customerID int64) ([]string, error)
}

// usecase 

type CustomerUsecase struct {
	repo CustomerRepo
}

func NewCustomerUsecase(repo CustomerRepo) *CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

// business Logic 

func (uc *CustomerUsecase) CreateCustomer(ctx context.Context, c *Customer) error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	return uc.repo.CreateCustomer(ctx, c)
}

func (uc *CustomerUsecase) GetCustomer(ctx context.Context, id int64) (*Customer, error) {
	return uc.repo.GetCustomer(ctx, id)
}

func (uc *CustomerUsecase) ListCustomers(ctx context.Context) ([]*Customer, error) {
	return uc.repo.ListCustomers(ctx)
}
