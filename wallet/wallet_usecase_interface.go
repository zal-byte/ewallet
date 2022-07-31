package wallet

import "ewallet/model"

type WalletUsecase interface {
	GetByID(id string) (*model.Wallets, error)
}
