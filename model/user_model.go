package model

type Users struct {
	ID      int    `json:"id" xorm:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
