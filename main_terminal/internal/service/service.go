package service

import (
	"context"
	"main_terminal/internal/infrastructure"
	"main_terminal/internal/models"
)

type Stocks interface {
	GetStock(context.Context, *models.StockCompany, map[string][]string) (*models.StockCompany, error)
}

type Currency interface {
}

type Service struct {
	Stocks
	Currency
}

func NewService(i *infrastructure.Infrastructure) *Service {
	return &Service{
		Stocks:   nil,
		Currency: nil,
	}
}
