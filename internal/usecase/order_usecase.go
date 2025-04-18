package usecase

import (
	"github.com/facelessEmptiness/order_service/internal/domain"
	"github.com/facelessEmptiness/order_service/internal/repository"
)

type OrderUseCase struct {
	repo repository.OrderRepository
}

func NewOrderUseCase(r repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{repo: r}
}

func (uc *OrderUseCase) CreateOrder(o *domain.Order) (string, error) {
	o.Status = "pending"
	return uc.repo.Create(o)
}

func (uc *OrderUseCase) GetOrder(id string) (*domain.Order, error) {
	return uc.repo.GetByID(id)
}
