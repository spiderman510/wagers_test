package api

import (
	"database/sql"
	"wagers/internal/controller"
	"wagers/internal/repository"
	"wagers/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateAPIEngine(db *sql.DB) (*gin.Engine, error) {
	r := gin.New()
	v1 := r.Group("/")
	wagerRepository := repository.NewWagerRepository(db)
	purchaseRepository := repository.NewPurchaseRepository(db)
	wagerService := service.NewWagerService(wagerRepository, purchaseRepository)
	wagerController := controller.NewWagerController(wagerService)
	registerPath(v1, wagerController)

	return r, nil
}

func registerPath(r *gin.RouterGroup, wagerController *controller.WagerController) {
	r.POST("wagers", wagerController.Place)
	r.POST("buy/:wager_id", wagerController.Buy)
	r.GET("/wagers", wagerController.List)
}
