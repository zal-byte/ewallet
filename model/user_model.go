package model

import (
	"time"
)

type Users struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Username  string     `json:"-" gorm:"column:username"`
	Password  string     `json:"-" gorm:"column:password"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"-" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"-" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
