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

func (r *RedisRepo) StoreCompanyStocks(ctx context.Context, stockCompany *models.StockCompany) error {

	err := r.db.Set(ctx, stockCompany.Company, stockCompany.Stocks, 1000*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}
