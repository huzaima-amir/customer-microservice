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
	Emails      []Email
	PhoneNumbers []PhoneNumber
	Addresses    []Address
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

func (uc *CustomerUsecase) AddEmail(ctx context.Context, id int64, email string) error {
    customer, err := uc.repo.GetCustomer(ctx, id)
    if err != nil {
        return err
    }

    newEmail := Email{
        CustomerID: id,
        Email:      email,
    }

    customer.Emails = append(customer.Emails, newEmail)

    return uc.repo.UpdateCustomer(ctx, customer)
}


// func (uc *CustomerUsecase) DeleteEmail (ctx context.Context, e *Email) error {

// }