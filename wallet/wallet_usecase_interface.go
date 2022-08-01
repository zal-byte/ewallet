package wallet

import "ewallet/model"

type WalletUsecase interface {
	Create(wallet *model.Wallets) (*model.Wallets, error)
	GetById(id string) (*model.Wallets, error)
	Delete(id string) error
	Update(id string, wall *model.Wallets) (*model.Wallets, error)
	GetByUserId(id string) (*model.Wallets, error)
}
