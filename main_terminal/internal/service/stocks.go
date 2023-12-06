package service

import (
	"context"
	"main_terminal/internal/infrastructure"
	"main_terminal/internal/models"
)

type StockLogic struct {
	Infra *infrastructure.Infrastructure
}

func NewStockLogic(infra *infrastructure.Infrastructure) *StockLogic {
	return &StockLogic{
		Infra: infra,
	}
}

func (s *StockLogic) GetStock(ctx context.Context, company *models.StockCompany, params map[string][]string) (*models.StockCompany, error) {
	comp, err := s.Infra.API.ExchangeAPI.GetStocks(ctx, company, params)

	if err != nil {
		return nil, err
	}

	err = s.Infra.Cache.StoreCompanyStocks(ctx, comp)

	if err != nil {
		return nil, err
	}

	return comp, nil
}
