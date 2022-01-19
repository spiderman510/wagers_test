package service

import (
	"time"
	"wagers/internal/entity"
	"wagers/test/mocks"
)

type WagerServiceTestContext struct {
	WagerRepository    *mocks.WagerRepository
	PurchaseRepository *mocks.PurchaseRepository
	WagerService       WagerService
}

func CreateWagerServiceTestContext() *WagerServiceTestContext {
	wagerRepository := &mocks.WagerRepository{}
	purchaseRepository := &mocks.PurchaseRepository{}
	wagerService := NewWagerService(wagerRepository, purchaseRepository)
	return &WagerServiceTestContext{
		WagerRepository:    wagerRepository,
		PurchaseRepository: purchaseRepository,
		WagerService:       wagerService,
	}
}

func CreateWagerTestData() *entity.Wager {
	var percentageSold float64 = 12.0
	var amountSold int = 12
	return &entity.Wager{
		ID:                  1,
		TotalWagerValue:     100,
		Odds:                1,
		SellingPercentage:   10,
		SellingPrice:        float64(100.0),
		CurrentSellingPrice: float64(76.0),
		PercentageSold:      &percentageSold,
		AmountSold:          &amountSold,
		PlacedAt:            time.Date(2022, 1, 1, 10, 10, 10, 651387237, time.UTC),
	}
}

func CreatePurchaseTestData() *entity.Purchase {
	return &entity.Purchase{
		Id:          1,
		WagerId:     1,
		BuyingPrice: float64(20.0),
		BoughtAt:    time.Date(2022, 1, 1, 10, 10, 10, 651387237, time.UTC),
	}
}
