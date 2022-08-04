package handler

import (
	"ewallet/authentication/usecase"
	"ewallet/model"
	tokenUsecase "ewallet/token/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	jwttoken    tokenUsecase.TokenUsecase
}

func CreateAuthHandler(r *gin.Engine, authUsecase usecase.AuthUsecase, token tokenUsecase.TokenUsecase) {

	handler := AuthHandler{
		authUsecase: authUsecase,
		jwttoken:    token,
	}

	r.POST("/login", handler.login)

}

func (auth *AuthHandler) login(c *gin.Context) {
	var credential model.Credential

	err := c.ShouldBindJSON(&credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, errs := auth.authUsecase.Login(&credential)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	c.Abort()
}
