package repository

import (
	"context"
	"restAPI/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}

func (repo *OrderRepository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order

	rows, err := repo.pool.Query(ctx, "SELECT order_number, total_price, date_time, status FROM orders")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var o models.Order
		err := rows.Scan(&o.OrderNumber, &o.TotalPrice, &o.DateTime, &o.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	rows.Close()

	return orders, nil
}

func (repo *OrderRepository) GetOrderByID(ctx context.Context, id int64) (*models.Order, error) {
	var o models.Order
	err := repo.pool.QueryRow(ctx,
		"SELECT order_number, total_price, date_time, status FROM orders WHERE order_number = $1", id).
		Scan(&o.OrderNumber, &o.TotalPrice, &o.DateTime, &o.Status)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (repo *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	_, err := repo.pool.Exec(ctx, "INSERT INTO orders (order_number, total_price, date_time, status) "+
		"VALUES ($1, $2, $3, $4)", order.OrderNumber, order.TotalPrice, order.DateTime, order.Status)
	if err != nil {
		return err
	}
	return nil
}

func (repo *OrderRepository) UpdateTotalPrice(ctx context.Context, id int64, order *models.Order) error {
	_, err := repo.pool.Exec(ctx, "UPDATE orders SET total_price = $1 WHERE order_number = $2",
		order.TotalPrice, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *OrderRepository) DeleteOrder(ctx context.Context, id int64) error {
	_, err := repo.pool.Exec(ctx, "DELETE FROM orders WHERE order_number = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *OrderRepository) UpdateOrderStatus(ctx context.Context, id int64, status string) error {
	_, err := repo.pool.Exec(ctx, "UPDATE orders SET status = $1 WHERE order_number = $2",
		status, id)
	if err != nil {
		return err
	}
	return nil
}
