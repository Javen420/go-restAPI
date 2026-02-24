package service

import (
	"context"
	"restAPI/db/repository"
	"restAPI/models"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{repo}
}

func (s *MenuService) GetFullMenu(ctx context.Context) (*models.FullMenu, error) {
	return s.repo.GetFullMenu(ctx)
}

func (s *MenuService) GetAllBurgers(ctx context.Context) ([]models.Burger, error) {
	return s.repo.GetAllBurgers(ctx)
}

func (s *MenuService) GetAllDrinks(ctx context.Context) ([]models.Drink, error) {
	return s.repo.GetAllDrinks(ctx)
}

func (s *MenuService) GetAllSides(ctx context.Context) ([]models.Sides, error) {
	return s.repo.GetAllSides(ctx)
}

func (s *MenuService) GetAllDesserts(ctx context.Context) ([]models.Dessert, error) {
	return s.repo.GetAllDesserts(ctx)
}
