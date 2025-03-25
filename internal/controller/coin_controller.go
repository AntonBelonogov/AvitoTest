package controller

import (
	"net/http"

	"AvitoTest/internal/model/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"AvitoTest/internal/repository"
	"AvitoTest/internal/service"
)

func InitCoin(api *gin.RouterGroup, db *gorm.DB) {
	txRepo := repository.NewTransactionRepository(db)
	hisRepo := repository.NewHistoryRepository(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	coinService := service.NewCoinService(txRepo, hisRepo, userRepo, productRepo)
	ctrl := NewCoinController(coinService)

	api.POST("/sendCoin", ctrl.HandleSendCoin)
	api.GET("/buy/:item", ctrl.HandleBuyItem)
}

type CoinController struct {
	service *service.CoinService
}

func NewCoinController(service *service.CoinService) *CoinController {
	return &CoinController{service: service}
}

func (ctrl *CoinController) HandleSendCoin(ctx *gin.Context) {
	var request = dto.SendCoinRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.SendCoin(request, ctx.GetHeader("user_id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (ctrl *CoinController) HandleBuyItem(ctx *gin.Context) {
	item := ctx.Param("item")

	if err := ctrl.service.BuyItem(item, ctx.GetHeader("user_id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
