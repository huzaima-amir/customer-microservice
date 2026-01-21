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
	PhoneNumber      string
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
	ListCustomer(ctx context.Context) ([]*Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (*Customer, error)
	GetCustomerByPhoneNumber(ctx context.Context, phone string) (*Customer, error)
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

func (uc *CustomerUsecase) DeleteCustomer(ctx context.Context, id int64) error { // TODO check if customer exissts before deleting
    _, err := uc.repo.GetCustomer(ctx, id)
    if err != nil {
        return err
    }

    return uc.repo.DeleteCustomer(ctx, id)
}

func (uc *CustomerUsecase) UpdateCustomer(ctx context.Context, c *Customer) error {
    return uc.repo.UpdateCustomer(ctx, c)
}

func (uc *CustomerUsecase) GetCustomer(ctx context.Context, id int64) (*Customer, error) {
	return uc.repo.GetCustomer(ctx, id)
}

func (uc *CustomerUsecase) GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	return uc.repo.GetCustomerByEmail(ctx, email)
}

func (uc *CustomerUsecase) GetCustomerByPhoneNumber(ctx context.Context, phone string) (*Customer, error) {
	return uc.repo.GetCustomerByPhoneNumber(ctx, phone)
}

func (uc *CustomerUsecase) ListCustomer(ctx context.Context) ([]*Customer, error) {
	return uc.repo.ListCustomer(ctx)
}

func (uc *CustomerUsecase) AddEmail(ctx context.Context, id int64, e string) error {
	// TODO add check for existing customer with this email
	customer, err := uc.repo.GetCustomer(ctx, id)
	if err != nil {			
		return err
	}

	newEmail := Email{			
		CustomerID: id,
		Email:      e,
	}

	customer.Emails = append(customer.Emails, newEmail)

	return uc.repo.UpdateCustomer(ctx, customer)

}



func (uc *CustomerUsecase) DeleteEmail(ctx context.Context, id int64, e string) error {
	customer, err := uc.repo.GetCustomer(ctx, id)
	if err != nil {
		return err
	}

	for i, x := range customer.Emails {
		if x.Email == e {
			customer.Emails = append(customer.Emails[:i], customer.Emails[i+1:]...)
			break
		}
	}

	return uc.repo.UpdateCustomer(ctx, customer)

}

func (uc *CustomerUsecase) AddPhoneNumber(ctx context.Context, id int64, p string) error {
	// TODO add check for existing customer with this phone numer
	customer, err := uc.repo.GetCustomer(ctx, id)
	if err != nil {
		return err
	}

	newPhoneNumber := PhoneNumber {
		CustomerID: id,
		PhoneNumber: p,
	}

	customer.PhoneNumbers = append(customer.PhoneNumbers, newPhoneNumber)

	return uc.repo.UpdateCustomer(ctx, customer)
}

func (uc *CustomerUsecase) DeletePhoneNumber(ctx context.Context, id int64, p string) error {
	customer, err := uc.repo.GetCustomer(ctx, id)

	if err != nil {
		return err
	}

	for i, x := range customer.PhoneNumbers {
		if x.PhoneNumber == p {
			customer.PhoneNumbers = append(customer.PhoneNumbers[:i], customer.PhoneNumbers[i + 1:]...)
			break
		}
	}

	return uc.repo.UpdateCustomer(ctx, customer)
}

func (uc *CustomerUsecase) AddAddress(ctx context.Context, id int64, address string) error { // addresses can be shared between different customers
	customer, err := uc.repo.GetCustomer(ctx, id)
	if err != nil {
		return err
	}

	newAddress := Address {
		CustomerID: id,
		Address: address,
	}

	customer.Addresses = append(customer.Addresses, newAddress)

	return uc.repo.UpdateCustomer(ctx, customer)
}

func (uc *CustomerUsecase) DeleteAddress(ctx context.Context, id int64, address string) error {
	customer, err := uc.repo.GetCustomer(ctx, id)
	if err != nil {
		return err
	}

	for i, x := range customer.Addresses {
		if x.Address == address {
			customer.Addresses = append(customer.Addresses[:i], customer.Addresses[i + 1:]...)
			break
		}
	}

	return uc.repo.UpdateCustomer(ctx, customer)
}

func (uc *CustomerUsecase) ListEmail(ctx context.Context, id int64) ([]Email, error) {
    customer, err := uc.repo.GetCustomer(ctx, id)
    if err != nil {
        return nil, err
    }
    return customer.Emails, nil
}


func (uc *CustomerUsecase) ListPhoneNumber(ctx context.Context, id int64) ([]PhoneNumber, error) {
	    customer, err := uc.repo.GetCustomer(ctx, id)
    if err != nil {
        return nil, err
    }
    return customer.PhoneNumbers, nil
}


func (uc *CustomerUsecase) ListAddress(ctx context.Context, id int64) ([]Address, error) {
    customer, err := uc.repo.GetCustomer(ctx, id)
    if err != nil {
        return nil, err
    }
    return customer.Addresses, nil
}


// TODO - add error inside delete methods if item doesnt exist (currently ignored)