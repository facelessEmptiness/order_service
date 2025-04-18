package grpc

import (
	"context"
	"github.com/facelessEmptiness/order_service/internal/domain"
	"github.com/facelessEmptiness/order_service/internal/usecase"
	pb "github.com/facelessEmptiness/order_service/proto"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	uc *usecase.OrderUseCase
}

func NewOrderHandler(uc *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{uc: uc}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	items := make([]domain.OrderItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = domain.OrderItem{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
	o := &domain.Order{
		UserID:        req.UserId,
		Items:         items,
		PaymentMethod: req.PaymentMethod,
	}
	id, err := h.uc.CreateOrder(o)
	if err != nil {
		return nil, err
	}
	o.ID = id
	return &pb.OrderResponse{
		Id:            o.ID,
		UserId:        o.UserID,
		Items:         req.Items,
		PaymentMethod: o.PaymentMethod,
		Status:        o.Status,
	}, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.OrderID) (*pb.OrderResponse, error) {
	o, err := h.uc.GetOrder(req.Id)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.OrderItem, len(o.Items))
	for i, item := range o.Items {
		items[i] = &pb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		}
	}
	return &pb.OrderResponse{
		Id:            o.ID,
		UserId:        o.UserID,
		Items:         items,
		PaymentMethod: o.PaymentMethod,
		Status:        o.Status,
	}, nil
}
