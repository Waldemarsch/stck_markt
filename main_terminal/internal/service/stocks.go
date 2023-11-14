package service

import (
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

func (s *StockLogic) GetStockList(companies []*models.StockCompany) []*models.StockCompany {
	return nil
}
