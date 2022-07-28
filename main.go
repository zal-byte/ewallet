package main

import (
	"ewallet/database"
	"ewallet/user/handler"
	"ewallet/user/repo"
	"ewallet/user/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.DbConn()
	con, _ := db.DB()
	defer con.Close()

	router := gin.Default()

	userRepo := repo.CreateUserRepo(db)
	// userRepo := repo.CreateUserRepoAlt(db)
	userUsecase := usecase.CreateUserUsecase(userRepo)
	handler.CreateUserHandler(router, userUsecase)

	router.Run("localhost:8080")
}
