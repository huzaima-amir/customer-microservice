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
    return &pb.AddEmailReply{}, nil
}
func (s *CustomerService) AddPhoneNumber(ctx context.Context, req *pb.AddPhoneNumberReq) (*pb.AddPhoneNumberReply, error) {
    return &pb.AddPhoneNumberReply{}, nil
}
func (s *CustomerService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerReq) (*pb.UpdateCustomerReply, error) {
    return &pb.UpdateCustomerReply{}, nil
}
func (s *CustomerService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerReq) (*pb.DeleteCustomerReply, error) {
    return &pb.DeleteCustomerReply{}, nil
}
func (s *CustomerService) ListCustomer(ctx context.Context, req *pb.ListCustomerReq) (*pb.ListCustomerReply, error) {
    return &pb.ListCustomerReply{}, nil
}
func (s *CustomerService) AddAddress(ctx context.Context, req *pb.AddAddressReq) (*pb.AddAddressReply, error) {
    return &pb.AddAddressReply{}, nil
}
func (s *CustomerService) ListAddress(ctx context.Context, req *pb.ListAddressReq) (*pb.ListAddressReply, error) {
    return &pb.ListAddressReply{}, nil
}
func (s *CustomerService) ListPhoneNumber(ctx context.Context, req *pb.ListPhoneNumberReq) (*pb.ListPhoneNumberReply, error) {
    return &pb.ListPhoneNumberReply{}, nil
}
func (s *CustomerService) ListEmail(ctx context.Context, req *pb.ListEmailReq) (*pb.ListEmailReply, error) {
    return &pb.ListEmailReply{}, nil
}
func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerReq) (*pb.GetCustomerReply, error) {
    return &pb.GetCustomerReply{}, nil
}
func (s *CustomerService) GetCustomerByEmail(ctx context.Context, req *pb.GetCustomerByEmailReq) (*pb.GetCustomerByEmailReply, error) {
    return &pb.GetCustomerByEmailReply{}, nil
}
func (s *CustomerService) GetCustomerByPhoneNumber(ctx context.Context, req *pb.GetCustomerByPhoneNumberReq) (*pb.GetCustomerByPhoneNumberReply, error) {
    return &pb.GetCustomerByPhoneNumberReply{}, nil
}
func (s *CustomerService) DeletePhoneNumber(ctx context.Context, req *pb.DeletePhoneNumberReq) (*pb.DeletePhoneNumberReply, error) {
    return &pb.DeletePhoneNumberReply{}, nil
}
func (s *CustomerService) DeleteAddress(ctx context.Context, req *pb.DeleteAddressReq) (*pb.DeleteAddressReply, error) {
    return &pb.DeleteAddressReply{}, nil
}
func (s *CustomerService) DeleteEmail(ctx context.Context, req *pb.DeleteEmailReq) (*pb.DeleteEmailReply, error) {
    return &pb.DeleteEmailReply{}, nil
}
