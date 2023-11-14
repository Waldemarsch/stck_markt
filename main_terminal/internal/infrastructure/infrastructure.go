package infrastructure

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"main_terminal/internal/infrastructure/cache"
	"main_terminal/internal/infrastructure/repository"
	"main_terminal/internal/models"
)

type WebAPI interface {
}

type Repository interface {
}

type Cache interface {
	StoreCompanyStocks(context.Context, []*models.StockCompany) error
}

type Infrastructure struct {
	WebAPI
	Repository
	Cache
}

func NewInfrastructure(Rdb *gorm.DB, Cdb *redis.Client) *Infrastructure {
	return &Infrastructure{
		WebAPI:     nil,
		Repository: repository.NewSqliteRepo(Rdb),
		Cache:      cache.NewRedisRepo(Cdb),
	}
}
