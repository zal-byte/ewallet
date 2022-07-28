package handler

import (
	"ewallet/model"
	"ewallet/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func CreateUserHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{
		userUsecase: userUsecase,
	}
	r.GET("/users", userHandler.getAll)
	r.GET("/users/:id", userHandler.getById)
	r.POST("/users", userHandler.addUser)

}

func (e *UserHandler) getById(c *gin.Context) {
	id := c.Param("id")

	user, err := e.userUsecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (e *UserHandler) getAll(c *gin.Context) {

	users, err := e.userUsecase.GetAll()

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, users)
}

func (e *UserHandler) addUser(c *gin.Context) {
	var user *model.Users

	err := c.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
	}

	newUser, err := e.userUsecase.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
