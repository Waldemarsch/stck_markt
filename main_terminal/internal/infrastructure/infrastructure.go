package infrastructure

import (
	"gorm.io/gorm"
	"main_terminal/internal/infrastructure/repository"
)

type WebAPI interface {
}

type Repository interface {
}

type Infrastructure struct {
	WebAPI
	Repository
}

func NewInfrastructure(db *gorm.DB) *Infrastructure {
	return &Infrastructure{
		WebAPI:     nil,
		Repository: repository.NewSqliteRepo(db),
	}
}
