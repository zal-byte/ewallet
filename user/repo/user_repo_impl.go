package repo

import (
	"ewallet/model"
	"ewallet/user"
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

type UserRepoXorm struct {
	DB *xorm.Engine
}

func CreateUserRepoAlt(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

//Xorm
func (e *UserRepoXorm) Create(Users *model.Users) (*model.Users, error) {
	affected, err := e.DB.Insert(Users)
	if err != nil {
		return nil, err
	}

	Users.ID = int(affected)

	return Users, nil
}

func (e *UserRepoXorm) GetAll() (*[]model.Users, error) {
	var user []model.Users

	err := e.DB.Find(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (e *UserRepoXorm) GetById(id string) (*model.Users, error) {
	var user model.Users

	err := e.DB.Find(&user, id)

	return &user, err
}

func (e *UserRepoXorm) Update(id string, user *model.Users) (*model.Users, error) {
	_, err := e.DB.ID(&model.Users{}).Update(&user)
	return user, err
}

//end xorm

//gorm
func CreateUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

//endgorm
func (e *UserRepoImpl) GetAll() (*[]model.Users, error) {
	var user []model.Users

	err := e.DB.Find(&user)

	return &user, err.Error
}

func (e *UserRepoImpl) Create(Users *model.Users) (*model.Users, error) {
	err := e.DB.Save(Users).Error
	return Users, err
}

func (e *UserRepoImpl) GetByUsername(username string) (*model.Users, error) {
	var usr *model.Users
	res := e.DB.Model(&model.Users{}).Where("username", username).First(&usr)

	return usr, res.Error
}

func (e *UserRepoImpl) GetById(id string) (*model.Users, error) {
	var user model.Users

	usrId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	res := e.DB.Where(&model.Users{ID: usrId}).First(&user)

	return &user, res.Error
}

func (e *UserRepoImpl) Update(id string, user *model.Users) (*model.Users, error) {
	var usr model.Users
	res := e.DB.Model(&model.Users{}).Where("id", id).Updates(&user).First(&usr)

	return &usr, res.Error
}

func (e *UserRepoImpl) Delete(id string) error {
	// var usr model.Users

	usrId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	res := e.DB.Where(&model.Users{ID: usrId}).Delete(&model.Users{})
	fmt.Println(res.Error)
	return res.Error
}
