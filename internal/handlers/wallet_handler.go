package handlers

import (
	"sandbox/pockett-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type WalletHandler interface {
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
}

type walletHandler struct {
	walletRepository repositories.WalletRepository
}

func NewWalletHandler(walletRepository repositories.WalletRepository) *walletHandler {
	return &walletHandler{walletRepository}
}

func (h *walletHandler) Add(c *gin.Context) {
	// var body models.WalletReq
	// err := c.ShouldBindJSON(&body)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "invalid request body")
	// 	return
	// }
	// res, err := h.service.Add(body)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// 	return
	// }
	// c.JSON(http.StatusOK, res)
	// return
}

func (h *walletHandler) Delete(c *gin.Context) {}

func (h *walletHandler) Update(c *gin.Context) {}

func (h *walletHandler) Get(c *gin.Context) {}
