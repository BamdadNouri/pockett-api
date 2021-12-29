package handlers

import (
	"net/http"
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/modules"
	"sandbox/pockett-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo repositories.UserRepository
}

func NewUserHandler(userRepo repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo}
}

func (h *UserHandler) Add(c *gin.Context) {
	var body models.UserCreateReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}

	res, err := modules.NewUser(h.userRepo).
		Register(body).
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
		map[string]interface{}{
			"token": res.Token(),
			"user":  res.ToRes()},
	)
	return
}

func (h *UserHandler) Login(c *gin.Context) {
	var body models.UserLogin
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}
	res, err := modules.NewUser(h.userRepo).
		Login(body).
		Result()
	if err != nil {
		// if token == "" unauthorized...
		c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"token": res.Token(),
			"user":  res.ToRes()},
	)
	return
}

func (h *UserHandler) Delete(c *gin.Context) {}

func (h *UserHandler) Update(c *gin.Context) {}

func (h *UserHandler) Me(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "invalid request body")
	// 	return
	// }
	// modules.NewUser().Find(id)
}
