package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"restAPI/models"
)

type MenuRepository struct {
	pool *pgxpool.Pool
}

func NewMenuRepository(pool *pgxpool.Pool) *MenuRepository {
	return &MenuRepository{pool: pool}
}

func (repo *MenuRepository) GetFullMenu(ctx context.Context) (*models.FullMenu, error) {
	burgers, err := repo.GetAllBurgers(ctx)
	if err != nil {
		return nil, err
	}

	drinks, err := repo.GetAllDrinks(ctx)
	if err != nil {
		return nil, err
	}

	sides, err := repo.GetAllSides(ctx)
	if err != nil {
		return nil, err
	}

	desserts, err := repo.GetAllDesserts(ctx)
	if err != nil {
		return nil, err
	}

	return &models.FullMenu{
		Burgers:  burgers,
		Drinks:   drinks,
		Sides:    sides,
		Desserts: desserts,
	}, nil
}

func (repo *MenuRepository) GetAllBurgers(ctx context.Context) ([]models.Burger, error) {
	var burgers []models.Burger

	rows, err := repo.pool.Query(ctx, "SELECT name, price, calories FROM burgers")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b models.Burger
		err := rows.Scan(&b.Name, &b.Price, &b.Calories)
		if err != nil {
			return nil, err
		}
		burgers = append(burgers, b)
	}
	rows.Close()

	return burgers, nil
}

func (repo *MenuRepository) GetAllDrinks(ctx context.Context) ([]models.Drink, error) {
	var drinks []models.Drink

	rows, err := repo.pool.Query(ctx, "SELECT name, price, calories, is_iced FROM drinks")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d models.Drink
		err := rows.Scan(&d.Name, &d.Price, &d.Calories, &d.IsIced)
		if err != nil {
			return nil, err
		}
		drinks = append(drinks, d)
	}
	rows.Close()

	return drinks, nil
}

func (repo *MenuRepository) GetAllSides(ctx context.Context) ([]models.Sides, error) {
	var sides []models.Sides

	rows, err := repo.pool.Query(ctx, "SELECT name, price, calories FROM sides")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s models.Sides
		err := rows.Scan(&s.Name, &s.Price, &s.Calories)
		if err != nil {
			return nil, err
		}
		sides = append(sides, s)
	}
	rows.Close()

	return sides, nil
}

func (repo *MenuRepository) GetAllDesserts(ctx context.Context) ([]models.Dessert, error) {
	var desserts []models.Dessert

	rows, err := repo.pool.Query(ctx, "SELECT name, price, calories FROM desserts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var de models.Dessert
		err := rows.Scan(&de.Name, &de.Price, &de.Calories)
		if err != nil {
			return nil, err
		}
		desserts = append(desserts, de)
	}
	rows.Close()

	return desserts, nil
}
