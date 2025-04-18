package repository

import (
	"github.com/facelessEmptiness/order_service/internal/domain"
)

type OrderRepository interface {
	Create(o *domain.Order) (string, error)
	GetByID(id string) (*domain.Order, error)
}
