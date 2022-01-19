package service

import (
	"net/http"
	"testing"
	"wagers/internal/entity"
	"wagers/internal/errors"

	"github.com/stretchr/testify/assert"
)

func TestPlaceService(t *testing.T) {
	t.Run("Should return TotalWagerValueRequiredError if total wager value <= 0", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.PlaceWager(&entity.PlaceRequest{
			TotalWagerValue: -1,
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.TotalWagerValueRequiredError,
		}, err)
	})
	t.Run("Should return OddsRequiredError if odds <= 0", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.PlaceWager(&entity.PlaceRequest{
			TotalWagerValue: 100,
			Odds:            -1,
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.OddsRequiredError,
		}, err)
	})
	t.Run("Should return SellingPercentageRangeError if selling percentage < 1", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.PlaceWager(&entity.PlaceRequest{
			TotalWagerValue:   100,
			Odds:              1,
			SellingPercentage: 0,
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.SellingPercentageRangeError,
		}, err)
	})
	t.Run("Should return SellingPercentageRangeError if selling percentage > 100", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.PlaceWager(&entity.PlaceRequest{
			TotalWagerValue:   100,
			Odds:              1,
			SellingPercentage: 101,
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.SellingPercentageRangeError,
		}, err)
	})
	t.Run("Should return SellingPriceRequiredError if selling price <= total wager value * (selling percentage / 100)", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.PlaceWager(&entity.PlaceRequest{
			TotalWagerValue:   100,
			Odds:              1,
			SellingPercentage: 10,
			SellingPrice:      float64(5.0),
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.SellingPriceRequiredError,
		}, err)
	})
	t.Run("Should return InsertDataError if failed to insert data", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.PlaceRequest{
			TotalWagerValue:   100,
			Odds:              1,
			SellingPercentage: 10,
			SellingPrice:      float64(20.0),
		}
		s.WagerRepository.Mock.On("Create", &entity.Wager{
			TotalWagerValue:     request.TotalWagerValue,
			Odds:                request.Odds,
			SellingPercentage:   request.SellingPercentage,
			SellingPrice:        request.SellingPrice,
			CurrentSellingPrice: request.SellingPrice,
		}).Return(nil, DatabaseError)
		_, err := s.WagerService.PlaceWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.InsertDataError,
		}, err)
	})
	t.Run("Should return inserted wager", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.PlaceRequest{
			TotalWagerValue:   100,
			Odds:              1,
			SellingPercentage: 10,
			SellingPrice:      float64(20.0),
		}
		wager := CreateWagerTestData()
		s.WagerRepository.Mock.On("Create", &entity.Wager{
			TotalWagerValue:     request.TotalWagerValue,
			Odds:                request.Odds,
			SellingPercentage:   request.SellingPercentage,
			SellingPrice:        request.SellingPrice,
			CurrentSellingPrice: request.SellingPrice,
		}).Return(wager, nil)
		actualWager, err := s.WagerService.PlaceWager(request)
		assert.Nil(t, err)
		assert.Equal(t, wager, actualWager)
	})
}
