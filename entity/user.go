package entity

import "time"

type User struct {
	ID int `gorm:"not null" json:"id"`
	Username string `gorm:"type:varchar(100);not null" json:"name"`
	Email string `gorm:"type:varchar(100);not null" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	IsActive bool `json:"is_active"`
}