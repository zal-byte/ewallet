package model

type Users struct {
	ID      int    `json:"id" gorm:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
