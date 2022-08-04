package repo

import (
	"ewallet/model"

	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func CreateAuthRepo(db *gorm.DB) AuthRepo {
	return AuthRepo{
		db,
	}
}

func (auth *AuthRepo) Login(credential *model.Credential) (*model.Users, error) {
	var cred *model.Users
	res := auth.DB.Model(&model.Users{}).Where("username", credential.Username).First(&cred)
	return cred, res.Error
}
