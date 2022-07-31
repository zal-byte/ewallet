package wallet

import "ewallet/model"

type WalletRepo interface {
	GetByID(id string) (*model.Wallets, error)
}
