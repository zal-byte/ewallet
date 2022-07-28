package user

import "ewallet/model"

type UserRepo interface {
	Create(user *model.Users) (*model.Users, error)
	GetAll() (*[]model.Users, error)
	GetById(id string) (*model.Users, error)
	// Update(id string, user *model.User) (*model.User, error)
	// Delete(id string) error
}
