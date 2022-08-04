package user

import "ewallet/model"

type UserUsecase interface {
	Create(user *model.Users) (*model.Users, error)
	GetAll() (*[]model.Users, error)
	GetById(id string) (*model.Users, error)
	GetByUsername(username string) (*model.Users, error)
	Update(id string, user *model.Users) (*model.Users, error)
	Delete(id string) error
}
