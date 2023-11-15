package api

import (
	"context"
	"main_terminal/internal/models"
)

type WebAPI interface {
}

type ExchangeAPI interface {
	GetStocks(context.Context, []*models.StockCompany, map[string][]string) ([]*models.StockCompany, error)
}

type API struct {
	WebAPI
	ExchangeAPI
}

func NewAPI(exAPI string) *API {
	return &API{
		WebAPI:      nil,
		ExchangeAPI: NewMoexAPI(exAPI, ""),
	}
}
