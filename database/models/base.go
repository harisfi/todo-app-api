package models

import "time"

type IDModel struct {
    ID uint `gorm:"primaryKey" json:"id"`
}

type TimestampModel struct {
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   *time.Time  `gorm:"index" json:"deleted_at"`
}