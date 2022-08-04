package usecase

import (
	"errors"
	"ewallet/authentication/repo"
	"ewallet/model"
	"ewallet/token/usecase"
	userusecase "ewallet/user/usecase"
	"strings"
)

type AuthUsecase struct {
	authRepo    repo.AuthRepo
	jwttoken    usecase.TokenUsecase
	userUsecase userusecase.UserUsecaseImpl
}

func CreateAuthUsecase(authUsecase repo.AuthRepo, token usecase.TokenUsecase, usrusecase userusecase.UserUsecaseImpl) AuthUsecase {
	return AuthUsecase{
		authUsecase,
		token,
		usrusecase,
	}
}

func (auth AuthUsecase) Login(credential *model.Credential) (*model.LoginResponse, error) {

	data, err := auth.authRepo.Login(credential)

	if err != nil {
		if strings.Contains("record not found", err.Error()) {
			return nil, errors.New("user not found")
		}
	}

	if data.Username == credential.Username {
		if data.Password == credential.Password {
			usrinfo, errs := auth.userUsecase.GetByUsername(credential.Username)

			if errs != nil {
				if strings.Contains("record not found", errs.Error()) {
					return nil, errors.New("user not found")
				}
			}


			generatedToken := auth.jwttoken.GenerateToken()
			generatedToken.Username = credential.Username

			var logres model.LoginResponse = model.LoginResponse{
				Status: "OK",
				Error:  "",
				Data:   usrinfo,
				Token:  generatedToken,
			}

			return &logres, nil

		} else {
			//Invalid password
			return nil, errors.New("invalid password")
		}
	} else {
		return nil, errors.New("invalid username")
		//Invalid username
	}
}
