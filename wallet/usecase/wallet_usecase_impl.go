package usecase

import (
	"ewallet/model"
	"ewallet/wallet"
)

type WalletUsecase struct {
	DB wallet.WalletRepo
}

func CreateWalletUsecase(walletRepo wallet.WalletRepo) WalletUsecase {
	return WalletUsecase{walletRepo}
}

func (e *WalletUsecase) GetByID(id string) (*model.Wallets, error) {
	return nil, nil
}
