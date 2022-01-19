package service

import (
	"net/http"
	"wagers/internal/entity"
	"wagers/internal/errors"
	"wagers/internal/repository"

	"github.com/rs/zerolog/log"
)

type WagerService interface {
	ListWager(request *entity.ListRequest) ([]*entity.Wager, *errors.Errors)
	PlaceWager(request *entity.PlaceRequest) (*entity.Wager, *errors.Errors)
	BuyWager(request *entity.BuyRequest) (*entity.Purchase, *errors.Errors)
}

type service struct {
	wagerRepository    repository.WagerRepository
	purchaseRepository repository.PurchaseRepository
}

func (s *service) ListWager(request *entity.ListRequest) ([]*entity.Wager, *errors.Errors) {
	if request.Page <= 0 {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.PageRequiredError,
		}
	}
	if request.Limit <= 0 {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.LimitRequiredError,
		}
	}
	wagers, err := s.wagerRepository.Query((request.Page-1)*request.Limit, request.Limit)

	if err != nil {
		log.Error().Err(err).Msg("Fail to get wager list")
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.ListWagerError,
		}
	}

	return wagers, nil
}

func (s *service) PlaceWager(request *entity.PlaceRequest) (*entity.Wager, *errors.Errors) {
	if request.TotalWagerValue <= 0 {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.TotalWagerValueRequiredError,
		}
	}
	if request.Odds <= 0 {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.OddsRequiredError,
		}
	}
	if request.SellingPercentage < 1 || request.SellingPercentage > 100 {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.SellingPercentageRangeError,
		}
	}
	if request.SellingPrice <= float64(request.TotalWagerValue*request.SellingPercentage/100) {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.SellingPriceRequiredError,
		}
	}
	wager := &entity.Wager{
		TotalWagerValue:     request.TotalWagerValue,
		Odds:                request.Odds,
		SellingPercentage:   request.SellingPercentage,
		SellingPrice:        request.SellingPrice,
		CurrentSellingPrice: request.SellingPrice,
	}

	wager, err := s.wagerRepository.Create(wager)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create wager")
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.InsertDataError,
		}
	}
	return wager, nil
}

func (s *service) BuyWager(request *entity.BuyRequest) (*entity.Purchase, *errors.Errors) {
	wg, err := s.wagerRepository.QueryById(request.WagerId)
	if err != nil {
		log.Error().Err(err).Msg("Fail to get wager by id")
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.WagerGetError,
		}
	}

	if wg == nil {
		return nil, &errors.Errors{
			Code:    http.StatusNotFound,
			Message: errors.WagerNotFoundError,
		}
	}

	if request.BuyingPrice > wg.CurrentSellingPrice {
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.BuyingPriceValueError,
		}
	}

	insertedPurchase, err := s.purchaseRepository.Create(&entity.Purchase{
		WagerId:     request.WagerId,
		BuyingPrice: request.BuyingPrice,
	})
	if err != nil {
		log.Error().Err(err).Msg("Fail to create a purchase")
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.InsertDataError,
		}
	}
	amountSold := 0
	if wg.AmountSold == nil {
		amountSold = 1
	} else {
		amountSold = *wg.AmountSold + 1
	}
	wg.AmountSold = &amountSold

	percentageSold := float64(*wg.AmountSold * 100.0 / wg.TotalWagerValue)
	wg.PercentageSold = &percentageSold
	err = s.wagerRepository.Update(&entity.Wager{
		ID:                  request.WagerId,
		TotalWagerValue:     wg.TotalWagerValue,
		Odds:                wg.Odds,
		SellingPercentage:   wg.SellingPercentage,
		SellingPrice:        wg.SellingPrice,
		CurrentSellingPrice: wg.CurrentSellingPrice - request.BuyingPrice,
		PercentageSold:      wg.PercentageSold,
		AmountSold:          wg.AmountSold,
	})
	if err != nil {
		log.Error().Err(err).Msg("Fail to update wager by id")
		return nil, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.WagerUpdatedError,
		}
	}

	return insertedPurchase, nil
}
func NewWagerService(wagerRepository repository.WagerRepository, purchaseRepository repository.PurchaseRepository) WagerService {
	return &service{
		wagerRepository:    wagerRepository,
		purchaseRepository: purchaseRepository,
	}
}
