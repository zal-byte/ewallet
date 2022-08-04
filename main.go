package main

import (
	authHandler "ewallet/authentication/handler"
	authRepo "ewallet/authentication/repo"
	authUsecase "ewallet/authentication/usecase"
	"ewallet/database"
	tokenUsecase "ewallet/token/usecase"
	userHandler "ewallet/user/handler"
	userRepo "ewallet/user/repo"
	userUsecases "ewallet/user/usecase"
	walletHandler "ewallet/wallet/handler"
	walletRepo "ewallet/wallet/repo"
	walletUsecases "ewallet/wallet/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.DbConn()

	con, _ := db.DB()
	defer con.Close()

	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})
	//middleware

	//repository
	uRepo := userRepo.CreateUserRepo(db)
	wRepo := walletRepo.CreateWalletRepo(db)
	authRepos := authRepo.CreateAuthRepo(db)
	//usecase
	tUsecase := tokenUsecase.CreateTokenUsecase()
	uUsecase := userUsecases.CreateUserUsecase(uRepo)
	wUsecase := walletUsecases.CreateWalletUsecase(wRepo, uRepo)

	aUsecase := authUsecase.CreateAuthUsecase(authRepos, tUsecase, uUsecase)
	//Handler
	authHandler.CreateAuthHandler(router, aUsecase, tUsecase)
	userHandler.CreateUserHandler(router, uUsecase, tUsecase)
	walletHandler.CreateWalletHandler(router, &wUsecase)

	router.Run("localhost:8080")
}
