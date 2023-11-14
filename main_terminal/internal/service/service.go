package service

import "main_terminal/internal/infrastructure"

type Stocks interface {
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
