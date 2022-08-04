package authentication

import "ewallet/model"

type AuthUsecaseInterface interface {
	Login(credential *model.Credential) (*model.LoginResponse, error)
}
