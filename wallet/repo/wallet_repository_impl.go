package repo

import (
	"ewallet/model"
	"ewallet/wallet"

	"gorm.io/gorm"
)

type WalletRepoImpl struct {
	DB *gorm.DB
}

func CreateWalletRepo(DB *gorm.DB) wallet.WalletRepo {
	return &WalletRepoImpl{DB}
}

func (e *WalletRepoImpl) Create(wallet *model.Wallets) (*model.Wallets, error) {
	res := e.DB.Create(&wallet)

	return wallet, res.Error
}

func (e *WalletRepoImpl) GetById(id string) (*model.Wallets, error) {
	var wallet model.Wallets

	res := e.DB.Model(&model.Wallets{}).Where("id = ?", id).Preload("Users").First(&wallet)

	return &wallet, res.Error
}

func (e *WalletRepoImpl) Delete(id string) error {
	var wall model.Wallets

	res := e.DB.Model(wall).Where("id = ?", id).First(&wall).Delete(&wall)

	return res.Error
}

func (e *WalletRepoImpl) Update(id string, wall *model.Wallets) (*model.Wallets, error) {

	res := e.DB.Model(&model.Wallets{}).Where("id = ?", id).First(&model.Wallets{}).Updates(&wall).Preload("Users").First(&wall)

	return wall, res.Error
}

func (e *WalletRepoImpl) GetByUserId(id string) (*model.Wallets, error) {
	var wallet model.Wallets

	res := e.DB.Model(&model.Wallets{}).Where("user_id = ?", id).First(&wallet)

	return &wallet, res.Error
}
