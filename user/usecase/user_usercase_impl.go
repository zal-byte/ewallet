package usecase

import (
	"errors"
	"ewallet/model"
	"ewallet/user"
	"fmt"
	"strings"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) UserUsecaseImpl {
	return UserUsecaseImpl{userRepo}
}

func (e UserUsecaseImpl) Create(user *model.Users) (*model.Users, error) {
	if user.ID != 0 {
		return nil, errors.New("set id not permitted")
	}
	return e.userRepo.Create(user)
}

func (e UserUsecaseImpl) GetAll() (*[]model.Users, error) {
	return e.userRepo.GetAll()
}

func (e UserUsecaseImpl) GetById(id string) (*model.Users, error) {

	res, err := e.userRepo.GetById(id)

	if err != nil {
		if strings.Contains("record not found", err.Error()) {
			return nil, errors.New("user not found")
		}
	}

	return res, nil
}

func (e UserUsecaseImpl) Update(id string, user *model.Users) (*model.Users, error) {
	res, err := e.userRepo.Update(id, user)

	if err != nil {
		fmt.Println(err)
	}

	return res, err
}

func (e UserUsecaseImpl) Delete(id string) error {
	err := e.userRepo.Delete(id)

	if err != nil {
		return err
	}
	return nil
}
