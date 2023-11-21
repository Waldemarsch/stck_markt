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

func NewInfrastructure(Rdb *gorm.DB, Cdb *redis.Client, exAPI string) *Infrastructure {
	return &Infrastructure{
		API:        api.NewAPI(exAPI),
		Repository: nil,
		Cache:      cache.NewRedisRepo(Cdb),
	}
}
