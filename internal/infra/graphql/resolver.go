package graphql

import (
	"context"

	"desafio_clean_architecture/internal/application/usecase"
)

type Resolver struct {
	ListOrdersUC *usecase.ListOrdersUseCase
}

func (r *Resolver) Query_listOrders(ctx context.Context) ([]*Order, error) {
	orders, err := r.ListOrdersUC.Execute()
	if err != nil {
		return nil, err
	}

	var gqlOrders []*Order
	for _, o := range orders {
		gqlOrders = append(gqlOrders, &Order{
			ID:        o.ID,
			Product:   o.Product,
			Price:     o.Price,
			CreatedAt: o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return gqlOrders, nil
}

type Order struct {
	ID        string  `json:"id"`
	Product   string  `json:"product"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}
