package models

import "time"

type URL struct {
	ID        int    `gorm:"primaryKey"`
	Code      string `gorm:"not null"`
	LongURL   string `gorm:"not null"`
	Clicks    int
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
