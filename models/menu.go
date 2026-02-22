package models

import (
	"fmt"
	"time"
)

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
