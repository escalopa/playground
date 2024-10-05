package handler

import (
	"net/http"

	"github.com/escalopa/SQLite/model"
	"github.com/gin-gonic/gin"
)

type userService interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(name string) (model.User, error)
}

type UserHandler struct {
	us userService
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.us.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := h.us.CreateUser(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}
