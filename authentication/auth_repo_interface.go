package authentication

import "ewallet/model"

type AuthRepoInterface interface {
	Login(credential *model.Credential) (*model.Users, error)
}
