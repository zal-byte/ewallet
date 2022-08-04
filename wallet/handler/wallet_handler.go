package handler

import (
	"ewallet/middleware"
	"ewallet/model"
	"ewallet/wallet"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletUsecase wallet.WalletUsecase
}

func CreateWalletHandler(r *gin.Engine, walletUsecase wallet.WalletUsecase) {
	walletHandler := WalletHandler{
		walletUsecase: walletUsecase,
	}

	group := r.Group("/api")

	group.Use(middleware.Middleware())
	group.POST("/wallets", walletHandler.create)
	group.GET("/wallets/:id", walletHandler.getById)
	group.DELETE("/wallets/:id", walletHandler.delete)
	group.PUT("/wallets/:id", walletHandler.update)
}

func (e *WalletHandler) update(c *gin.Context) {
	id := c.Param("id")

	var wallet model.Wallets

	c.ShouldBindJSON(&wallet)

	res, err := e.walletUsecase.Update(id, &wallet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (e *WalletHandler) delete(c *gin.Context) {
	id := c.Param("id")

	res := e.walletUsecase.Delete(id)

	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (e *WalletHandler) getById(c *gin.Context) {

	id := c.Param("id")

	wallets, err := e.walletUsecase.GetById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, wallets)

}
func (e *WalletHandler) create(c *gin.Context) {
	var wallets model.Wallets

	errs := c.ShouldBindJSON(&wallets)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	res, err := e.walletUsecase.Create(&wallets)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
