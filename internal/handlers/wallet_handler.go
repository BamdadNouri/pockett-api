package handlers

import (
	"net/http"
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/modules"
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
	uid := getUserID(c)

	var body models.WalletCreateReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request",
			},
		)
		return
	}
	res, err := modules.NewWallet(uid).
		AddWallet(body).
		Result()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		res.ToRes(),
	)
}

func (h *walletHandler) Delete(c *gin.Context) {}

func (h *walletHandler) Update(c *gin.Context) {}

func (h *walletHandler) Get(c *gin.Context) {}
