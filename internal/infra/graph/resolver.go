package graph

import (
	"desafio_clean_architecture/internal/usecase"

	"desafio_clean_architecture/internal/entity"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	OrderRepo          entity.OrderRepositoryInterface
}
