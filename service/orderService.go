package service

import (
	"context"
	"fmt"
	"restAPI/db/repository"
	"restAPI/models"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return s.repo.GetAllOrders(ctx)
}

func (s *OrderService) GetOrder(ctx context.Context, id int64) (*models.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *OrderService) CreateOrder(ctx context.Context, items []models.Item) (*models.Order, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("order must have at least one item")
	}

	order, err := models.NewOrder(items)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateTotalPrice(ctx context.Context, id int64, order *models.Order) error {
	return s.repo.UpdateTotalPrice(ctx, id, order)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id int64) error {
	return s.repo.DeleteOrder(ctx, id)
}

func (s *OrderService) ChangeOrderStatus(ctx context.Context, id int64, status string) error {
	if status != models.StatusPending && status != models.StatusCompleted && status != models.StatusCancelled {
		return fmt.Errorf("invalid status: %s", status)
	}
	return s.repo.UpdateOrderStatus(ctx, id, status)
}
