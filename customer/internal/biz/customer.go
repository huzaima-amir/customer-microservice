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

    // email
    AddEmail(ctx context.Context, e *Email) error
    DeleteEmail(ctx context.Context, customerID int64, email string) error
    ListEmails(ctx context.Context, customerID int64) ([]string, error)

    // phone
    AddPhoneNumber(ctx context.Context, p *PhoneNumber) error
    DeletePhoneNumber(ctx context.Context, customerID int64, phone string) error
    ListPhoneNumbers(ctx context.Context, customerID int64) ([]string, error)

    // address
    AddAddress(ctx context.Context, a *Address) error
    DeleteAddress(ctx context.Context, customerID int64, address string) error
    ListAddresses(ctx context.Context, customerID int64) ([]string, error)

    // transactions
    Tx(ctx context.Context, fn func(ctx context.Context) error) error
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

func (uc *CustomerUsecase) AddEmail(ctx context.Context, id int64, e string) (*Email, error) {
	if e == "" {
		return nil, errors.New("email cannot be empty")
	}

	//ensure customer exists
	if _, err := uc.repo.GetCustomer(ctx, id); err != nil {
		return nil, err
	}

	email := &Email{
		CustomerID: id,
		Email:      e,
	}

	if err := uc.repo.AddEmail(ctx, email); err != nil {
		return nil, err
	}

	return email, nil
}


func (uc *CustomerUsecase) DeleteEmail(ctx context.Context, id int64, e string) error {
	if _, err := uc.repo.GetCustomer(ctx, id); err != nil {
		return err
	}

	return uc.repo.DeleteEmail(ctx, id, e)
}


func (uc *CustomerUsecase) AddPhoneNumber(ctx context.Context, id int64, p string) (*PhoneNumber, error) {
	if p == "" {
		return nil, errors.New("phone number cannot be empty")
	}

	if _, err := uc.repo.GetCustomer(ctx, id); err != nil {
		return nil, err
	}

	phone := &PhoneNumber{
		CustomerID:  id,
		PhoneNumber: p,
	}

	if err := uc.repo.AddPhoneNumber(ctx, phone); err != nil {
		return nil, err
	}

	return phone, nil
}


func (uc *CustomerUsecase) DeletePhoneNumber(ctx context.Context, id int64, p string) error {
	if _, err := uc.repo.GetCustomer(ctx, id); err != nil {
		return err
	}
	return uc.repo.DeletePhoneNumber(ctx, id, p)
}


func (uc *CustomerUsecase) AddAddress(ctx context.Context, id int64, addr string) (*Address, error) {
	if addr == "" {
		return nil, errors.New("address cannot be empty")
	}

	if _, err := uc.repo.GetCustomer(ctx, id); err != nil {
		return nil, err
	}

	address := &Address{
		CustomerID: id,
		Address:    addr,
	}

	if err := uc.repo.AddAddress(ctx, address); err != nil {
		return nil, err
	}

	return address, nil
}



func (uc *CustomerUsecase) DeleteAddress(ctx context.Context, id int64, address string) error {
	if _, err := uc.repo.GetCustomer(ctx, id); err != nil {
		return err
	}
	return uc.repo.DeleteAddress(ctx, id, address)
}

func (uc *CustomerUsecase) ListEmail(ctx context.Context, id int64) ([]string, error) {
	return uc.repo.ListEmails(ctx, id)
}

func (uc *CustomerUsecase) ListPhoneNumber(ctx context.Context, id int64) ([]string, error) {
	return uc.repo.ListPhoneNumbers(ctx, id)
}


func (uc *CustomerUsecase) ListAddress(ctx context.Context, id int64) ([]string, error) {
    return uc.repo.ListAddresses(ctx, id) 
}

func (uc *CustomerUsecase) CreateCustomerWithDetails(
    ctx context.Context,
    c *Customer,
    e *Email,
    p *PhoneNumber,
    a *Address,
) error {

    return uc.repo.Tx(ctx, func(ctx context.Context) error {

        //Create customer
        if err := uc.repo.CreateCustomer(ctx, c); err != nil {
            return err
        }

        // Add email (if provided)
        if e != nil {
            e.CustomerID = c.ID
            if err := uc.repo.AddEmail(ctx, e); err != nil {
                return err
            }
        }

        // Add phone number (if provided)
        if p != nil {
            p.CustomerID = c.ID
            if err := uc.repo.AddPhoneNumber(ctx, p); err != nil {
                return err
            }
        }

        //Add address (if provided)
        if a != nil {
            a.CustomerID = c.ID
            if err := uc.repo.AddAddress(ctx, a); err != nil {
                return err
            }
        }

        return nil
    })
}



