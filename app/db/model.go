package db

import "time"

type User struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"column:password_hash;not null"`
	Email        string    `gorm:"unique;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
