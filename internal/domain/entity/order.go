package entity

import (
	"time"
)

type Order struct {
	ID        string
	Product   string
	Price     float64
	CreatedAt time.Time
}
