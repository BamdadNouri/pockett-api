package handlers

import (
	"net/http"
	"sandbox/pockett-api/internal/errs"
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/modules"
	"sandbox/pockett-api/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionHandler(transactionRepo repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{transactionRepo}
}

func (h *TransactionHandler) Add(c *gin.Context) {
	uid := getUserID(c)

	var body models.TransactionCreateReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": errs.ErrInvalidRequest,
			},
		)
		return
	}
	res, err := modules.NewTransaction(uid, h.transactionRepo).
		Create(body).
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

func (h *TransactionHandler) Delete(c *gin.Context) {
	uid := getUserID(c)
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := modules.NewTransaction(uid, h.transactionRepo).
		SoftDelete(uint64(id)).
		Result()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		res.ToRes(),
	)
	return
}

func (h *TransactionHandler) Update(c *gin.Context) {}

func (h *TransactionHandler) Get(c *gin.Context) {
	uid := getUserID(c)
	sid := c.Query("id")
	sWallet := c.Param("walletID")
	wallet, err := strconv.Atoi(sWallet)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			map[string]string{"message": err.Error()},
		)
		return
	}
	spage := c.Query("page")
	ssize := c.Query("size")

	page, size := 0, 10
	// var err error
	if spage != "" {
		page, _ = strconv.Atoi(spage)
	}
	if ssize != "" {
		size, _ = strconv.Atoi(ssize)
	}
	var res []models.TransactionRes
	tr := modules.NewTransaction(uid, h.transactionRepo)
	if sid != "" {
		id, _ := strconv.Atoi(sid)
		tr.Find(uint64(id), uint64(wallet))
	} else {
		res = tr.Bulk(uint64(wallet), page, size)
	}
	balance := tr.GetBalance()
	_, err = tr.Result()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
		return
	}
	modules.NewWallet(uid)
	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"transactions": res,
			"wallet": map[string]interface{}{
				"balance": balance,
				"curr":    "IRT",
			},
		},
	)
	return
}

func getUserID(c *gin.Context) uint64 {
	uid, _ := c.Get("userID")
	return uid.(uint64)
}
