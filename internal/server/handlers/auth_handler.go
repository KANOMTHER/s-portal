package handlers

import (
	"github.com/gin-gonic/gin"
	"s-portal/internal/domain/service"
)

type AuthHandler struct {
	service *service.AuthService
}

type LoginRequest struct {
	UserId   uint   `json:"id"`
	Password string `json:"password"`
}

func NewAuthHandler(service *service.AuthService) AuthHandler {
	return AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Status(context *gin.Context) {
	user, err := h.service.GetContextUser(context)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if user == nil {
		context.JSON(404, gin.H{})
		return
	}

	context.JSON(200, user)
}

func (h *AuthHandler) Login(context *gin.Context) {
	var request LoginRequest
	bindError := context.ShouldBindJSON(&request)

	if bindError != nil {
		context.JSON(500, gin.H{
			"message": bindError.Error(),
		})
		return
	}

	user, err := h.service.ValidateUser(request.UserId, request.Password)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.service.SetContextUser(context, *user)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(200, user)
}

func (h *AuthHandler) Logout(context *gin.Context) {
	err := h.service.UnsetContextUser(context)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{})
}
