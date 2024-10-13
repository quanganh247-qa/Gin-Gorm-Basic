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

type Notes struct {
	Note_id   int64     `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
