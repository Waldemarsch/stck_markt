package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"main_terminal/internal/models"
	"time"
)

type RedisRepo struct {
	db *redis.Client
}

func NewRedisRepo(db *redis.Client) *RedisRepo {
	return &RedisRepo{
		db: db,
	}
}

func (r *RedisRepo) StoreStocks(ctx context.Context, stockCompanies []*models.StockCompany) error {
	for _, stockCompany := range stockCompanies {
		err := r.db.Set(ctx, stockCompany.Company, stockCompany.Stocks, 100*time.Second).Err()

		if err != nil {
			return err
		}
	}
	return nil
}
