package wallet

import "ewallet/model"

type WalletRepo interface {
	Create(*model.Wallets) (*model.Wallets, error)
	GetById(id string) (*model.Wallets, error)
	Delete(id string) error
	Update(id string, wall *model.Wallets) (*model.Wallets, error)
	GetByUserId(id string) (*model.Wallets, error)
}
