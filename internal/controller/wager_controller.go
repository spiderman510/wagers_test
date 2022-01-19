package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wagers/internal/entity"
	"wagers/internal/errors"
	"wagers/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type WagerController struct {
	wagerService service.WagerService
}

const (
	WagerIdField = "wager_id"
)

func (w *WagerController) List(c *gin.Context) {
	request := &entity.ListRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wagers, e := w.wagerService.ListWager(request)
	if e != nil {
		c.JSON(e.Code, e.Message)
		return
	}
	wagersJson, err := json.Marshal(wagers)
	if err != nil {
		log.Error().Err(err).Msg("Fail to parse data")
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", wagersJson)
}
func (w *WagerController) Place(c *gin.Context) {
	request := &entity.PlaceRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wager := &entity.Wager{
		TotalWagerValue:     request.TotalWagerValue,
		Odds:                request.Odds,
		SellingPercentage:   request.SellingPercentage,
		SellingPrice:        request.SellingPrice,
		CurrentSellingPrice: request.SellingPrice,
	}

	wager, e := w.wagerService.PlaceWager(request)
	if e != nil {
		c.JSON(e.Code, e.Message)
		return
	}

	jsonData, err := json.Marshal(wager)

	if err != nil {
		log.Error().Err(err).Msg("Fail to marshal wager")
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}
func (w *WagerController) Buy(c *gin.Context) {
	request := &entity.BuyRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value := c.Param(WagerIdField)
	wagerId, _ := strconv.Atoi(value)
	if request.BuyingPrice <= 0 {
		c.JSON(http.StatusInternalServerError, errors.BuyingPriceRequiredError)
		return
	}
	request.WagerId = wagerId

	wg, e := w.wagerService.BuyWager(request)
	if e != nil {
		c.JSON(e.Code, e.Message)
		return
	}

	jsonData, err := json.Marshal(&wg)
	if err != nil {
		log.Error().Err(err).Msg("Fail to marshal wager")
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}
func NewWagerController(wagerService service.WagerService) *WagerController {
	return &WagerController{
		wagerService: wagerService,
	}
}
