package usecase

import (
	"desafio_clean_architecture/internal/entity"
)

type OrderLister interface {
	ListOrders() ([]entity.Order, error)
}

type ListOrdersUseCase struct {
	repo OrderLister
}

func NewListOrdersUseCase(repo OrderLister) *ListOrdersUseCase {
	return &ListOrdersUseCase{repo: repo}
}

func (u *ListOrdersUseCase) Execute() ([]entity.Order, error) {
	return u.repo.ListOrders()
}
