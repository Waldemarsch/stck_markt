package infrastructure

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"main_terminal/internal/infrastructure/api"
	"main_terminal/internal/infrastructure/cache"
	"main_terminal/internal/models"
)

type Repository interface {
}

type Cache interface {
	StoreCompanyStocks(context.Context, *models.StockCompany) error
}

type Infrastructure struct {
	API *api.API
	Repository
	Cache
}

func NewInfrastructure(Rdb *gorm.DB, CConfig *redis.Options, stockAPI string, currencyAPI string) *Infrastructure {
	return &Infrastructure{
		API:        api.NewAPI(stockAPI, currencyAPI),
		Repository: nil,
		Cache:      cache.NewRedisRepo(CConfig),
	}
}
