package handler

import (
	"ewallet/middleware"
	"ewallet/model"
	"ewallet/token/usecase"
	"ewallet/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase user.UserUsecase
	apiToken    usecase.TokenUsecase
}

func CreateUserHandler(r *gin.Engine, userUsecase user.UserUsecase, apiToken usecase.TokenUsecase) {
	userHandler := UserHandler{
		userUsecase: userUsecase,
		apiToken:    apiToken,
	}
	group := r.Group("/api")

	group.Use(middleware.Middleware())
	group.GET("/users", userHandler.getAll)
	group.GET("/users/:id", userHandler.getById)
	group.DELETE("/users/:id", userHandler.delete)
	group.PUT("/users/:id", userHandler.update)
	group.POST("/users", userHandler.addUser)

}

func (e *UserHandler) delete(c *gin.Context) {
	id := c.Param("id")

	err := e.userUsecase.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Delete successfully."})
}

func (e *UserHandler) update(c *gin.Context) {
	var user model.Users

	id := c.Param("id")
	c.ShouldBindJSON(&user)

	res, err := e.userUsecase.Update(id, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
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
