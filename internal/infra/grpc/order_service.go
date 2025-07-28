package grpc

import (
	"context"

	"desafio_clean_architecture/internal/application/usecase"
	"desafio_clean_architecture/internal/infra/grpc/pb/orderpb"
)

type OrderServiceServer struct {
	orderpb.UnimplementedOrderServiceServer
	ListOrdersUC *usecase.ListOrdersUseCase
}

func NewOrderServiceServer(listUC *usecase.ListOrdersUseCase) *OrderServiceServer {
	return &OrderServiceServer{ListOrdersUC: listUC}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	orders, err := s.ListOrdersUC.Execute()
	if err != nil {
		return nil, err
	}

	var protoOrders []*orderpb.Order
	for _, o := range orders {
		protoOrders = append(protoOrders, &orderpb.Order{
			Id:        o.ID,
			Product:   o.Product,
			Price:     float32(o.Price),
			CreatedAt: o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &orderpb.ListOrdersResponse{Orders: protoOrders}, nil
}
