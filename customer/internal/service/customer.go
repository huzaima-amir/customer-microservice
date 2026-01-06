package service

import (
	"context"

	pb "customer/api/customer/v1"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerReq) (*pb.CreateCustomerReply, error) {
    return &pb.CreateCustomerReply{}, nil
}
func (s *CustomerService) CreateEmail(ctx context.Context, req *pb.CreateEmailReq) (*pb.CreateEmailReply, error) {
    return &pb.CreateEmailReply{}, nil
}
func (s *CustomerService) CreatePhoneNumber(ctx context.Context, req *pb.CreatePhoneNumberReq) (*pb.CreatePhoneNumberReply, error) {
    return &pb.CreatePhoneNumberReply{}, nil
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
func (s *CustomerService) CreateAddress(ctx context.Context, req *pb.CreateAddressReq) (*pb.CreateAddressReply, error) {
    return &pb.CreateAddressReply{}, nil
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
