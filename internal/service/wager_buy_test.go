package service

import (
	"net/http"
	"testing"
	"wagers/internal/entity"
	"wagers/internal/errors"

	"github.com/stretchr/testify/assert"
)

func TestBuyService(t *testing.T) {
	t.Run("Should return WagerGetError if failed to get wager by id", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.BuyRequest{
			WagerId:     1,
			BuyingPrice: float64(2.0),
		}
		s.WagerRepository.Mock.On("QueryById", request.WagerId).Return(nil, DatabaseError)
		_, err := s.WagerService.BuyWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.WagerGetError,
		}, err)
	})
	t.Run("Should return WagerNotFoundError if wager not found", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.BuyRequest{
			WagerId:     1,
			BuyingPrice: float64(2.0),
		}
		s.WagerRepository.Mock.On("QueryById", request.WagerId).Return(nil, nil)
		_, err := s.WagerService.BuyWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusNotFound,
			Message: errors.WagerNotFoundError,
		}, err)
	})
	t.Run("Should return BuyingPriceValueError if buying price > current selling price", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.BuyRequest{
			WagerId:     1,
			BuyingPrice: float64(200.0),
		}
		wager := CreateWagerTestData()
		s.WagerRepository.Mock.On("QueryById", request.WagerId).Return(wager, nil)
		_, err := s.WagerService.BuyWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.BuyingPriceValueError,
		}, err)
	})

	t.Run("Should return InsertDataError if failed to insert purchase", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.BuyRequest{
			WagerId:     1,
			BuyingPrice: float64(20.0),
		}
		wager := CreateWagerTestData()
		s.WagerRepository.Mock.On("QueryById", request.WagerId).Return(wager, nil)
		s.PurchaseRepository.Mock.On("Create", &entity.Purchase{
			WagerId:     request.WagerId,
			BuyingPrice: request.BuyingPrice,
		}).Return(nil, DatabaseError)

		_, err := s.WagerService.BuyWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.InsertDataError,
		}, err)
	})

	t.Run("Should return WagerUpdatedError if failed to update wager", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.BuyRequest{
			WagerId:     1,
			BuyingPrice: float64(20.0),
		}
		wager := CreateWagerTestData()
		s.WagerRepository.Mock.On("QueryById", request.WagerId).Return(wager, nil)
		pc := CreatePurchaseTestData()
		s.PurchaseRepository.Mock.On("Create", &entity.Purchase{
			WagerId:     request.WagerId,
			BuyingPrice: request.BuyingPrice,
		}).Return(pc, nil)
		amountSold := *wager.AmountSold + 1
		percentageSold := float64(amountSold * 100.0 / wager.TotalWagerValue)
		s.WagerRepository.Mock.On("Update", &entity.Wager{
			ID:                  request.WagerId,
			TotalWagerValue:     wager.TotalWagerValue,
			Odds:                wager.Odds,
			SellingPercentage:   wager.SellingPercentage,
			SellingPrice:        wager.SellingPrice,
			CurrentSellingPrice: wager.CurrentSellingPrice - request.BuyingPrice,
			PercentageSold:      &percentageSold,
			AmountSold:          &amountSold,
		}).Return(DatabaseError)
		_, err := s.WagerService.BuyWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.WagerUpdatedError,
		}, err)
	})

	t.Run("Should return map data if success", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.BuyRequest{
			WagerId:     1,
			BuyingPrice: float64(20.0),
		}
		wager := CreateWagerTestData()
		s.WagerRepository.Mock.On("QueryById", request.WagerId).Return(wager, nil)
		pc := CreatePurchaseTestData()
		s.PurchaseRepository.Mock.On("Create", &entity.Purchase{
			WagerId:     request.WagerId,
			BuyingPrice: request.BuyingPrice,
		}).Return(pc, nil)
		amountSold := *wager.AmountSold + 1
		percentageSold := float64(amountSold * 100.0 / wager.TotalWagerValue)

		s.WagerRepository.Mock.On("Update", &entity.Wager{
			ID:                  request.WagerId,
			TotalWagerValue:     wager.TotalWagerValue,
			Odds:                wager.Odds,
			SellingPercentage:   wager.SellingPercentage,
			SellingPrice:        wager.SellingPrice,
			CurrentSellingPrice: wager.CurrentSellingPrice - request.BuyingPrice,
			PercentageSold:      &percentageSold,
			AmountSold:          &amountSold,
		}).Return(nil)
		purchase, err := s.WagerService.BuyWager(request)
		assert.Nil(t, err)
		assert.Equal(t, pc, purchase)
	})
}
