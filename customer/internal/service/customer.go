package service

import (
	"context"
	pb "customer/api/customer/v1"
	"customer/internal/biz"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	uc *biz.CustomerUsecase
}

func NewCustomerService(uc *biz.CustomerUsecase) *CustomerService {
	return &CustomerService{uc: uc}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerReq) (*pb.CreateCustomerReply, error) {
	customer := &biz.Customer{
        Name: req.Name,
        DateOfBirth: req.DateOfBirth,
    }
    
    if err := s.uc.CreateCustomer(ctx, customer); err != nil {
        return nil, err
    }

	return &pb.CreateCustomerReply{Id: customer.ID}, nil
}

func (s *CustomerService) AddEmail(ctx context.Context, req *pb.AddEmailReq) (*pb.AddEmailReply, error) {
    if err := s.uc.AddEmail(ctx, req.CustomerId, req.Email); err != nil {
        return nil, err
    }
    return &pb.AddEmailReply{Id: req.CustomerId}, nil
}

func (s *CustomerService) AddPhoneNumber(ctx context.Context, req *pb.AddPhoneNumberReq) (*pb.AddPhoneNumberReply, error) {
    if err := s.uc.AddPhoneNumber(ctx, req.CustomerId, req.PhoneNumber); err != nil {
        return nil, err
    }
    return &pb.AddPhoneNumberReply{Id: req.CustomerId,}, nil
}
func (s *CustomerService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerReq) (*pb.UpdateCustomerReply, error) {
    customer, err := s.uc.GetCustomer(ctx, req.Id)
    if err != nil {
        return nil, err
    }

    if req.Name != "" {
        customer.Name = req.Name
    }
    if req.DateOfBirth != "" {
        customer.DateOfBirth = req.DateOfBirth
    }

    if err := s.uc.UpdateCustomer(ctx, customer); err != nil {
        return nil, err
    }

    phoneNumbers := make([]string, len(customer.PhoneNumbers))
    for i, p := range customer.PhoneNumbers {
        phoneNumbers[i] = p.PhoneNumber
    }

    emails := make([]string, len(customer.Emails))
    for i, e := range customer.Emails {
        emails[i] = e.Email
    }

    addresses := make([]string, len(customer.Addresses))
    for i, a := range customer.Addresses {
        addresses[i] = a.Address
    }

    return &pb.UpdateCustomerReply{
        Id:           customer.ID,
        Name:         customer.Name,
        PhoneNumbers: phoneNumbers,
        Emails:       emails,
        Addresses:    addresses,
        DateOfBirth:  customer.DateOfBirth,
    }, nil
}


func (s *CustomerService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerReq) (*pb.DeleteCustomerReply, error) {
    err := s.uc.DeleteCustomer(ctx, req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.DeleteCustomerReply{
        Success: true,
    }, nil
}


func (s *CustomerService) ListCustomer(ctx context.Context, req *pb.ListCustomerReq) (*pb.ListCustomerReply, error) {
    customers, err := s.uc.ListCustomer(ctx)
    if err != nil {
        return nil, err
    }

    pbCustomers := make([]*pb.GetCustomerReply, 0, len(customers))

    for _, c := range customers {

        phoneNumbers := make([]string, len(c.PhoneNumbers))
        for i, p := range c.PhoneNumbers {
            phoneNumbers[i] = p.PhoneNumber
        }

        emails := make([]string, len(c.Emails))
        for i, e := range c.Emails {
            emails[i] = e.Email
        }

        addresses := make([]string, len(c.Addresses))
        for i, a := range c.Addresses {
            addresses[i] = a.Address
        }

        pbCustomers = append(pbCustomers, &pb.GetCustomerReply{
            Id:           c.ID,
            Name:         c.Name,
            PhoneNumbers: phoneNumbers,
            Emails:       emails,
            Addresses:    addresses,
            DateOfBirth:  c.DateOfBirth,
        })
    }

    return &pb.ListCustomerReply{
        Customers: pbCustomers,
    }, nil
}

func (s *CustomerService) AddAddress(ctx context.Context, req *pb.AddAddressReq) (*pb.AddAddressReply, error) {
    if err := s.uc.AddAddress(ctx, req.CustomerId, req.Address); err != nil {
        return nil, err
    }
    return &pb.AddAddressReply{}, nil
}
func (s *CustomerService) ListAddress(ctx context.Context, req *pb.ListAddressReq) (*pb.ListAddressReply, error) {
    addresses, err := s.uc.ListAddress(ctx, req.CustomerId)
    if err != nil {
        return nil, err
    }
    addressStrings := make([]string, len(addresses))
    for i, e := range addresses {
        addressStrings[i] = e.Address
    }
    return &pb.ListAddressReply{
        Addresses: addressStrings,
    }, nil
}

func (s *CustomerService) ListPhoneNumber(ctx context.Context, req *pb.ListPhoneNumberReq) (*pb.ListPhoneNumberReply, error) {
    phoneNumbers, err := s.uc.ListPhoneNumber(ctx, req.CustomerId)
    if err != nil {
        return nil, err
    }
    phoneNumberStrings := make([]string, len(phoneNumbers))
    for i, e := range phoneNumbers {
        phoneNumberStrings[i] = e.PhoneNumber
    }
    return &pb.ListPhoneNumberReply{
        PhoneNumbers: phoneNumberStrings,
    }, nil
}

func (s *CustomerService) ListEmail(ctx context.Context, req *pb.ListEmailReq) (*pb.ListEmailReply, error) {
    emails, err := s.uc.ListEmail(ctx, req.CustomerId)
    if err != nil {
        return nil, err
    }

    emailStrings := make([]string, len(emails))
    for i, e := range emails {
        emailStrings[i] = e.Email
    }

    return &pb.ListEmailReply{
        Emails: emailStrings,
    }, nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerReq) (*pb.GetCustomerReply, error) {
    customer, err := s.uc.GetCustomer(ctx, req.Id)
    if err != nil {
        return nil, err
    }

    phoneNumbers := make([]string, len(customer.PhoneNumbers))
    for i, p := range customer.PhoneNumbers {
        phoneNumbers[i] = p.PhoneNumber
    }

    emails := make([]string, len(customer.Emails))
    for i, e := range customer.Emails {
        emails[i] = e.Email
    }

    addresses := make([]string, len(customer.Addresses))
    for i, a := range customer.Addresses {
        addresses[i] = a.Address
    }

    return &pb.GetCustomerReply{
        Id:           customer.ID,
        Name:         customer.Name,
        PhoneNumbers: phoneNumbers,
        Emails:       emails,
        Addresses:    addresses,
        DateOfBirth:  customer.DateOfBirth,
    }, nil
}

func (s *CustomerService) GetCustomerByEmail(ctx context.Context, req *pb.GetCustomerByEmailReq) (*pb.GetCustomerByEmailReply, error) {
    customer, err := s.uc.GetCustomerByEmail(ctx, req.Email)
    if err != nil {
        return nil, err
    }
        phoneNumbers := make([]string, len(customer.PhoneNumbers))
    for i, p := range customer.PhoneNumbers {
        phoneNumbers[i] = p.PhoneNumber
    }

    emails := make([]string, len(customer.Emails))
    for i, e := range customer.Emails {
        emails[i] = e.Email
    }

    addresses := make([]string, len(customer.Addresses))
    for i, a := range customer.Addresses {
        addresses[i] = a.Address
    }

    return &pb.GetCustomerByEmailReply{
        Id:           customer.ID,
        Name:         customer.Name,
        PhoneNumbers: phoneNumbers,
        Emails:       emails,
        Addresses:    addresses,
        DateOfBirth:  customer.DateOfBirth,
    }, nil
}

func (s *CustomerService) GetCustomerByPhoneNumber(ctx context.Context, req *pb.GetCustomerByPhoneNumberReq) (*pb.GetCustomerByPhoneNumberReply, error) {
    customer, err := s.uc.GetCustomerByPhoneNumber(ctx, req.PhoneNumber)
    if err != nil {
        return nil, err
    }
    phoneNumbers := make([]string, len(customer.PhoneNumbers))
    for i, p := range customer.PhoneNumbers {
        phoneNumbers[i] = p.PhoneNumber
    }

    emails := make([]string, len(customer.Emails))
    for i, e := range customer.Emails {
        emails[i] = e.Email
    }

    addresses := make([]string, len(customer.Addresses))
    for i, a := range customer.Addresses {
        addresses[i] = a.Address
    }

    return &pb.GetCustomerByPhoneNumberReply{
        Id:           customer.ID,
        Name:         customer.Name,
        PhoneNumbers: phoneNumbers,
        Emails:       emails, 
        Addresses:    addresses,
        DateOfBirth:  customer.DateOfBirth,
    }, nil
}

func (s *CustomerService) DeletePhoneNumber(ctx context.Context, req *pb.DeletePhoneNumberReq) (*pb.DeletePhoneNumberReply, error) {
    err := s.uc.DeletePhoneNumber(ctx, req.CustomerId, req.PhoneNumber)
    if err != nil {
        return nil, err
    }

    return &pb.DeletePhoneNumberReply{
        Success: true,
    }, nil
}

func (s *CustomerService) DeleteAddress(ctx context.Context, req *pb.DeleteAddressReq) (*pb.DeleteAddressReply, error) {
    err := s.uc.DeleteAddress(ctx, req.CustomerId, req.Address)
    if err != nil {
        return nil, err
    }

    return &pb.DeleteAddressReply{
        Success: true,
    }, nil
}
func (s *CustomerService) DeleteEmail(ctx context.Context, req *pb.DeleteEmailReq) (*pb.DeleteEmailReply, error) {
    err := s.uc.DeleteEmail(ctx, req.CustomerId, req.Email)
    if err != nil {
        return nil, err
    }

    return &pb.DeleteEmailReply{
        Success: true,
    }, nil
}
