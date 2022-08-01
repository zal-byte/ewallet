package main

import (
	"ewallet/database"
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

	//repository
	uRepo := userRepo.CreateUserRepo(db)
	wRepo := walletRepo.CreateWalletRepo(db)
	//usecase
	uUsecase := userUsecases.CreateUserUsecase(uRepo)
	wUsecase := walletUsecases.CreateWalletUsecase(wRepo, uRepo)
	//Handler
	userHandler.CreateUserHandler(router, uUsecase)
	walletHandler.CreateWalletHandler(router, &wUsecase)

	router.Run("localhost:8080")
}
