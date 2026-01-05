package service

import (
	"context"

	pb "customer/api/customer/v1"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerReply, error) {
    return &pb.CreateCustomerReply{}, nil
}
