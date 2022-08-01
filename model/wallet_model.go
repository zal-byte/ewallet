package model

import (
	"time"

	"gorm.io/gorm"
)

type Wallets struct {
	gorm.Model `json:"-"`
	ID         int        `json:"id" gorm:"primary_key"`
	Balance    string     `json:"balance"`
	User_id    string     `json:"user_id" gorm:"column:user_id"`
	Users      Users      `json:"user" gorm:"foreignKey:ID"`
	CreatedAt  time.Time  `json:"-" gorm:"column:created_at"`
	UpdatedAt  time.Time  `json:"-" gorm:"column:updated_at"`
	DeletedAt  *time.Time `json:"-" gorm:"column:deleted_at"`
}
