package repository

import "gorm.io/gorm"

type SqliteRepo struct {
	db *gorm.DB
}

func NewSqliteRepo(db *gorm.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}
