package grpc

import (
	"context"

	grpcpb "github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/usecase"
)

type OrderService struct {
	grpcpb.UnimplementedOrderServiceServer
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewOrderService(
	createOrderUseCase *usecase.CreateOrderUseCase,
	listOrdersUseCase *usecase.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *grpcpb.CreateOrderRequest) (*grpcpb.CreateOrderResponse, error) {
	output, err := s.CreateOrderUseCase.Execute(usecase.CreateOrderInputDTO{
		ID:    req.GetId(),
		Price: float64(req.GetPrice()),
		Tax:   float64(req.GetTax()),
	})
	if err != nil {
		return nil, err
	}

	return &grpcpb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *grpcpb.Blank) (*grpcpb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := make([]*grpcpb.CreateOrderResponse, 0, len(output))
	for _, order := range output {
		orders = append(orders, &grpcpb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &grpcpb.ListOrdersResponse{Orders: orders}, nil
}
