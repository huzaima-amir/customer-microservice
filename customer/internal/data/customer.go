package data

import (
	"context"
	"customer/internal/biz"
	"gorm.io/gorm"
)

//  GORM models 
type Customer struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string
	DateOfBirth string
	Emails      []Email
	PhoneNumbers []PhoneNumber
	Addresses      []Address
}


type Email   struct {
	ID         int64  `gorm:"primaryKey"`
	CustomerID int64  `gorm:"index"`
	Email      string `gorm:"uniqueIndex"`
}

type PhoneNumber struct {
	ID         int64  `gorm:"primaryKey"`
	CustomerID int64  `gorm:"index"`
	PhoneNumber      string `gorm:"uniqueIndex"`
}

type Address  struct {
	ID         int64  `gorm:"primaryKey"`
	CustomerID int64  `gorm:"index"`
	Address    string
}

//  Repo 

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(data *Data) biz.CustomerRepo {
	return &customerRepo{db: data.db}
}


// helpers to map db slices to biz types
func toBizEmails(in []Email) []biz.Email {
    out := make([]biz.Email, len(in))
    for i, e := range in {
        out[i] = biz.Email{
            ID:         e.ID,
            CustomerID: e.CustomerID,
            Email:      e.Email,
        }
    }
    return out
}

func toBizPhones(in []PhoneNumber) []biz.PhoneNumber {
    out := make([]biz.PhoneNumber, len(in))
    for i, p := range in {
        out[i] = biz.PhoneNumber{
            ID:          p.ID,
            CustomerID:  p.CustomerID,
            PhoneNumber: p.PhoneNumber,
        }
    }
    return out
}

func toBizAddresses(in []Address) []biz.Address {
    out := make([]biz.Address, len(in))
    for i, a := range in {
        out[i] = biz.Address{
            ID:         a.ID,
            CustomerID: a.CustomerID,
            Address:    a.Address,
        }
    }
    return out
}


//  customer 
func (r *customerRepo) CreateCustomer(ctx context.Context, c *biz.Customer) error {
	model := Customer{
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	c.ID = model.ID
	return nil
}

func (r *customerRepo) UpdateCustomer(ctx context.Context, c *biz.Customer) error {
	return r.db.WithContext(ctx).
		Model(&Customer{}).
		Where("id = ?", c.ID).
		Updates(map[string]interface{}{
			"name":          c.Name,
			"date_of_birth": c.DateOfBirth,
		}).Error
}

func (r *customerRepo) DeleteCustomer(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&Customer{}, id).Error
}

func (r *customerRepo) GetCustomer(ctx context.Context, id int64) (*biz.Customer, error) { // TODO preload emails, phonenumbers, and addresses
	var m Customer
	err := r.db.WithContext(ctx).
			Preload("Emails").
			Preload("PhoneNumbers").				
			Preload("Addresses").
			First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &biz.Customer{
		ID:          m.ID,
		Name:        m.Name,
		DateOfBirth: m.DateOfBirth,
		Emails:      toBizEmails(m.Emails),
		PhoneNumbers: toBizPhones(m.PhoneNumbers),
		Addresses:    toBizAddresses(m.Addresses),
	}, nil
}

func (r *customerRepo) ListCustomer(ctx context.Context) ([]*biz.Customer, error) {
    var models []Customer
    err := r.db.WithContext(ctx).
        Preload("Emails").
        Preload("PhoneNumbers").
        Preload("Addresses").
        Find(&models).Error
    if err != nil {
        return nil, err
    }

    out := make([]*biz.Customer, 0, len(models))
    for _, m := range models {
        out = append(out, &biz.Customer{
            ID:           m.ID,
            Name:         m.Name,
            DateOfBirth:  m.DateOfBirth,
            Emails:       toBizEmails(m.Emails),
            PhoneNumbers: toBizPhones(m.PhoneNumbers),
            Addresses:    toBizAddresses(m.Addresses),
        })
    }
    return out, nil
}



// email 

func (r *customerRepo) AddEmail(ctx context.Context, e *biz.Email) error {
	return r.db.WithContext(ctx).Create(&Email{
		CustomerID: e.CustomerID,
		Email:      e.Email,
	}).Error
}

func (r *customerRepo) DeleteEmail(ctx context.Context, customerID int64, email string) error {
	return r.db.WithContext(ctx).
		Where("customer_id = ? AND email = ?", customerID, email).
		Delete(&Email{}).Error
}

func (r *customerRepo) ListEmails(ctx context.Context, customerID int64) ([]string, error) {
	var emails []string
	err := r.db.WithContext(ctx).
		Model(&Email{}).
		Where("customer_id = ?", customerID).
		Pluck("email", &emails).Error //to select a single column and get in a slice!!!
	return emails, err
}

func (r *customerRepo) GetCustomerByEmail(ctx context.Context, email string) (*biz.Customer, error) {
    var c Customer
    err := r.db.WithContext(ctx).
        Preload("Emails").
        Preload("PhoneNumbers").
        Preload("Addresses").
        Joins("JOIN emails ON emails.customer_id = customers.id").
        Where("emails.email = ?", email).
        First(&c).Error
    if err != nil {
        return nil, err
    }

    return &biz.Customer{
        ID:           c.ID,
        Name:         c.Name,
        DateOfBirth:  c.DateOfBirth,
        Emails:       toBizEmails(c.Emails),
        PhoneNumbers: toBizPhones(c.PhoneNumbers),
        Addresses:    toBizAddresses(c.Addresses),
    }, nil
}


// phone 
func (r *customerRepo) AddPhoneNumber(ctx context.Context, p *biz.PhoneNumber) error {
	return r.db.WithContext(ctx).Create(&PhoneNumber{
		CustomerID: p.CustomerID,
		PhoneNumber:      p.PhoneNumber,
	}).Error
}

func (r *customerRepo) DeletePhoneNumber(ctx context.Context, customerID int64, phone string) error {
	return r.db.WithContext(ctx).
		Where("customer_id = ? AND phone_number = ?", customerID, phone).
		Delete(&PhoneNumber{}).Error
}

func (r *customerRepo) ListPhoneNumbers(ctx context.Context, customerID int64) ([]string, error) {
	var phones []string
	err := r.db.WithContext(ctx).
		Model(&PhoneNumber{}).
		Where("customer_id = ?", customerID).
		Pluck("phone_number", &phones).Error
	return phones, err
}

func (r *customerRepo) GetCustomerByPhoneNumber(ctx context.Context, phone string) (*biz.Customer, error) {
    var c Customer
    err := r.db.WithContext(ctx).
        Preload("Emails").
        Preload("PhoneNumbers").
        Preload("Addresses").
        Joins("JOIN phone_numbers ON phone_numbers.customer_id = customers.id").
        Where("phone_numbers.phone_number = ?", phone).
        First(&c).Error
    if err != nil {
        return nil, err
    }

    return &biz.Customer{
        ID:           c.ID,
        Name:         c.Name,
        DateOfBirth:  c.DateOfBirth,
        Emails:       toBizEmails(c.Emails),
        PhoneNumbers: toBizPhones(c.PhoneNumbers),
        Addresses:    toBizAddresses(c.Addresses),
    }, nil
}


// address 
func (r *customerRepo) AddAddress(ctx context.Context, a *biz.Address) error {
	return r.db.WithContext(ctx).Create(&Address{
		CustomerID: a.CustomerID,
		Address:    a.Address,
	}).Error
}

func (r *customerRepo) DeleteAddress(ctx context.Context, customerID int64, address string) error {
	return r.db.WithContext(ctx).
		Where("customer_id = ? AND address = ?", customerID, address).
		Delete(&Address{}).Error
}

func (r *customerRepo) ListAddresses(ctx context.Context, customerID int64) ([]string, error) {
	var addresses []string
	err := r.db.WithContext(ctx).
		Model(&Address{}).
		Where("customer_id = ?", customerID).
		Pluck("address", &addresses).Error
	return addresses, err
}


// transaction helper:
func (r *customerRepo) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // temporarily replace the repo's DB with the transaction DB
        old := r.db
        r.db = tx
        defer func() { r.db = old }()
        return fn(ctx)
    })
}
