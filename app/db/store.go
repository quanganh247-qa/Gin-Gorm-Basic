package db

import (
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
