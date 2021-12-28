package handlers

import (
	"net/http"
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
	var body models.TransactionCreateReq
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
	res, err := modules.NewTransaction(1).
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
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := modules.NewTransaction(1).
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
	sid := c.Query("id")
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
	modules := modules.NewTransaction(1)
	if sid == "" {
		id, _ := strconv.Atoi(sid)
		modules.Find(uint64(id))
	} else {
		modules.Bulk(page, size)
	}
	res, err := modules.Result()
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
