package usecase

import (
	"errors"
	"ewallet/model"
	uRepo "ewallet/user"
	"ewallet/wallet"
	"strings"
)

type WalletUsecase struct {
	WalletRepo wallet.WalletRepo
	UserRepo   uRepo.UserRepo
}

func CreateWalletUsecase(walletRepo wallet.WalletRepo, userRepo uRepo.UserRepo) WalletUsecase {
	return WalletUsecase{walletRepo, userRepo}
}

func (e *WalletUsecase) Create(wallet *model.Wallets) (*model.Wallets, error) {

	ures, err := e.UserRepo.GetById(wallet.User_id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	re, _ := e.WalletRepo.GetByUserId(wallet.User_id)

	if re != nil {
		if re.User_id == wallet.User_id {
			return nil, errors.New("this user already has a wallet")
		}
	}

	res, err := e.WalletRepo.Create(wallet)

	res.Users = *ures

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *WalletUsecase) GetById(id string) (*model.Wallets, error) {
	return e.WalletRepo.GetById(id)
}

func (e *WalletUsecase) Delete(id string) error {

	err := e.WalletRepo.Delete(id)

	if err != nil {
		if strings.Contains("record not found", err.Error()) {
			return errors.New("wallet not found")
		}
	}
	return nil
}

func (e *WalletUsecase) Update(id string, wall *model.Wallets) (*model.Wallets, error) {
	return e.WalletRepo.Update(id, wall)
}

func (e *WalletUsecase) GetByUserId(id string) (*model.Wallets, error) {
	return e.WalletRepo.GetByUserId(id)
}
