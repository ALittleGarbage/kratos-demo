package order

import (
	"context"
	orderv1 "kratos-demo/api/order/v1"
	userv1 "kratos-demo/api/user/v1"
	"kratos-demo/internal/client/order"

	pb "kratos-demo/api/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	client *order.GRPCClient
}

func NewOrderService(client *order.GRPCClient) *OrderService {
	return &OrderService{client: client}
}

func (s *OrderService) GetUserByOrderId(ctx context.Context, req *pb.GetUserByOrderIdRequest) (*pb.GetUserByOrderIdReply, error) {
	resp, err := s.client.Uc.GetUser(ctx, &userv1.GetUserRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByOrderIdReply{
		User: &orderv1.UserDto{
			Id:       resp.GetUser().GetId(),
			Username: resp.GetUser().GetUsername(),
		},
	}, nil
}
