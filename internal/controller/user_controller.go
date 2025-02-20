package controller

import (
	"net/http"

	"AvitoTest/internal/model/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"AvitoTest/internal/repository"
	"AvitoTest/internal/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func InitUser(api *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	ctrl := NewUserController(userService)

	api.POST("/auth", ctrl.Login)
}

func (ctrl *UserController) Login(ctx *gin.Context) {
	var request = dto.AuthRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := ctrl.service.AuthUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
