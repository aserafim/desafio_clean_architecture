package repository

import (
	"database/sql"
	"desafio_clean_architecture/internal/domain/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) ListOrders() ([]entity.Order, error) {
	rows, err := r.db.Query("SELECT id, product, price, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var o entity.Order
		err := rows.Scan(&o.ID, &o.Product, &o.Price, &o.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
