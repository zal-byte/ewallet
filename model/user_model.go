package model

type Users struct {
	ID      int64  `json:"id" xorm:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
