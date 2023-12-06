package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
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

	st, _ := json.Marshal(stockCompany.Stocks)

	err := r.db.Set(ctx, stockCompany.Company, st, 1000*time.Second).Err()

	if err != nil {
		return err
	}

	log.Println("Stored to cache: ", stockCompany.Company)

	return nil
}
