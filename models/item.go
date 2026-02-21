package models

import "time"

type MenuItem interface {
	GetName() string
	GetPrice() float64
}

type SideDish struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Calories int     `json:"calories"`
}

func (s SideDish) GetName() string   { return s.Name }
func (s SideDish) GetPrice() float64 { return s.Price }

type Drink struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Calories int     `json:"calories"`
	IsIced   bool    `json:"is_iced"`
}

func (d Drink) GetName() string   { return d.Name }
func (d Drink) GetPrice() float64 { return d.Price }

type Burger struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Calories int     `json:"calories"`
	IsMeal   bool    `json:"is_meal"`
}

func (b Burger) GetName() string   { return b.Name }
func (b Burger) GetPrice() float64 { return b.Price }

type Dessert struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Calories int     `json:"calories"`
}

func (d Dessert) GetName() string   { return d.Name }
func (d Dessert) GetPrice() float64 { return d.Price }

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

func getTotalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.MenuItem.GetPrice() * float64(item.Quantity)
	}
	return total
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
