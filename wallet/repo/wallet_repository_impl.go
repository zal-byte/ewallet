package repo

import (
	"ewallet/model"
	"ewallet/wallet"

	"gorm.io/gorm"
)

type WalletRepoImpl struct {
	DB *gorm.DB
}

func CreateWalletImpl(DB *gorm.DB) wallet.WalletRepo {
	return &WalletRepoImpl{DB}
}

func (e *WalletRepoImpl) GetByID(id string) (*model.Wallets, error) {
	return nil, nil
}
