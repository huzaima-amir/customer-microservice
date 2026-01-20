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

func (r *customerRepo) GetCustomer(ctx context.Context, id int64) (*biz.Customer, error) {
	var m Customer
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &biz.Customer{
		ID:          m.ID,
		Name:        m.Name,
		DateOfBirth: m.DateOfBirth,
	}, nil
}

func (r *customerRepo) ListCustomers(ctx context.Context) ([]*biz.Customer, error) {
	var models []Customer
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	out := make([]*biz.Customer, 0, len(models))
	for _, m := range models {
		out = append(out, &biz.Customer{
			ID:          m.ID,
			Name:        m.Name,
			DateOfBirth: m.DateOfBirth,
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
		Joins("JOIN emails ON emails.customer_id = customers.id").
		Where("emails.email = ?", email).
		First(&c).Error
	if err != nil {
		return nil, err
	}
	return &biz.Customer{ID: c.ID, Name: c.Name, DateOfBirth: c.DateOfBirth}, nil
}

// phone 
func (r *customerRepo) AddPhone(ctx context.Context, p *biz.PhoneNumber) error {
	return r.db.WithContext(ctx).Create(&Phone{
		CustomerID: p.CustomerID,
		Phone:      p.PhoneNumber,
	}).Error
}

func (r *customerRepo) DeletePhone(ctx context.Context, customerID int64, phone string) error {
	return r.db.WithContext(ctx).
		Where("customer_id = ? AND phone = ?", customerID, phone).
		Delete(&Phone{}).Error
}

func (r *customerRepo) ListPhones(ctx context.Context, customerID int64) ([]string, error) {
	var phones []string
	err := r.db.WithContext(ctx).
		Model(&Phone{}).
		Where("customer_id = ?", customerID).
		Pluck("phone", &phones).Error
	return phones, err
}

func (r *customerRepo) GetCustomerByPhone(ctx context.Context, phone string) (*biz.Customer, error) {
	var c Customer
	err := r.db.WithContext(ctx).
		Joins("JOIN phones ON phones.customer_id = customers.id").
		Where("phones.phone = ?", phone).
		First(&c).Error
	if err != nil {
		return nil, err
	}
	return &biz.Customer{ID: c.ID, Name: c.Name, DateOfBirth: c.DateOfBirth}, nil
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
