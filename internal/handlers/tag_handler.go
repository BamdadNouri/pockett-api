package handlers

import (
	"sandbox/pockett-api/internal/services"

	"github.com/gin-gonic/gin"
)

type TagHandler interface {
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	GetBulk(c *gin.Context)
}

type tagHandler struct {
	service services.TagService
}

func NewTagHandler(service services.TagService) *tagHandler {
	return &tagHandler{service}
}

func (h *tagHandler) Add(c *gin.Context) {
	// var body models.TagReq
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

func (h *tagHandler) Delete(c *gin.Context) {}

func (h *tagHandler) Update(c *gin.Context) {}

func (h *tagHandler) GetBulk(c *gin.Context) {}
