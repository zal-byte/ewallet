package model

type Wallets struct {
	ID      int    `json:"id" gorm:"id"`
	Balance string `json:"balance"`
}
