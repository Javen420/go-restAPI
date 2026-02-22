package models

import (
	"fmt"
	"time"
)

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusCancelled = "cancelled"
)

func getTotalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.MenuItem.GetPrice() * float64(item.Quantity)
	}
	return total
}

type Item struct {
	MenuItem MenuItem `json:"menu_item"`
	Quantity int      `json:"quantity"`
}

type Order struct {
	OrderNumber int64     `json:"order_number"`
	Items       []Item    `json:"items"`
	TotalPrice  float64   `json:"total_price"`
	DateTime    time.Time `json:"date_time"`
	Status      string    `json:"status"`
}

func getNextOrderID() int64 {
	return 0
}

func NewOrder(items []Item) (*Order, error) {
	order := &Order{
		OrderNumber: getNextOrderID(),
		Items:       items,
		TotalPrice:  getTotalPrice(items),
		DateTime:    time.Now(),
		Status:      "pending",
	}
	return order, nil
}

func (o *Order) changeOrderStatus(status string) error {
	if status != StatusPending && status != StatusCompleted && status != StatusCancelled {
		return fmt.Errorf("invalid status: %s", status)
	}
	o.Status = status
	return nil
}
