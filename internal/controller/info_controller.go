package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"AvitoTest/internal/repository"
	"AvitoTest/internal/service"
)

func InitInfo(api *gin.RouterGroup, db *gorm.DB) {
	infoRepository := repository.NewHistoryRepository(db)
	userRepository := repository.NewUserRepository(db)
	infoService := service.NewInfoService(infoRepository, userRepository)
	ctrl := NewInfoController(infoService)

	api.GET("/info", ctrl.HandleGetInfo)
}

type InfoController struct {
	service *service.InfoService
}

func NewInfoController(service *service.InfoService) *InfoController {
	return &InfoController{service: service}
}

func (ctrl *InfoController) HandleGetInfo(ctx *gin.Context) {
	if response, err := ctrl.service.GetUserInfo(ctx.GetHeader("user_id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, response)
		return
	}
}
