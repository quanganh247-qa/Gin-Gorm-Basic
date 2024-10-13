package db

import (
	"context"

	"gorm.io/gorm"
)

var StoreDB Store

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return Store{db: db}
}

func InitStore(db *gorm.DB) Store {
	StoreDB = NewStore(db)
	return StoreDB
}

func (s *Store) ExecTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := s.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
