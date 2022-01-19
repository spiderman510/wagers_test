package repository

import (
	"database/sql"
	"testing"
	"wagers/internal/entity"

	"github.com/stretchr/testify/assert"
)

type WagerRepositoryTextContext struct {
	WagerRepository WagerRepository
	db              *sql.DB
}

func TestWagerRepository(t *testing.T) {
	initDb()
	t.Run("Test Create", func(t *testing.T) {
		s := CreateWagerRepositoryTextContext(db)
		wager := &entity.Wager{
			TotalWagerValue:     100,
			Odds:                1,
			SellingPercentage:   10,
			SellingPrice:        float64(5.0),
			CurrentSellingPrice: float64(5.0),
		}
		lastIndexId := getLastIndexId(s.db, wagerTableName)
		actualWager, err := s.WagerRepository.Create(wager)
		assert.Nil(t, err)
		assert.Equal(t, lastIndexId, actualWager.ID)
		assert.Equal(t, wager.TotalWagerValue, actualWager.TotalWagerValue)
		assert.Equal(t, wager.Odds, actualWager.Odds)
		assert.Equal(t, wager.SellingPercentage, actualWager.SellingPercentage)
		assert.Equal(t, wager.SellingPrice, actualWager.SellingPrice)
		assert.Equal(t, wager.CurrentSellingPrice, actualWager.CurrentSellingPrice)
	})
	t.Run("Test Update", func(t *testing.T) {
		s := CreateWagerRepositoryTextContext(db)
		percentageSold := float64(10.0)
		amountSold := 1
		wager := &entity.Wager{
			TotalWagerValue:     100,
			Odds:                1,
			SellingPercentage:   10,
			SellingPrice:        float64(5.0),
			CurrentSellingPrice: float64(5.0),
			ID:                  1,
			PercentageSold:      &percentageSold,
			AmountSold:          &amountSold,
		}
		err := s.WagerRepository.Update(wager)
		assert.Nil(t, err)
		updatedWager, err := s.WagerRepository.QueryById(wager.ID)
		assert.Nil(t, err)
		assert.Equal(t, wager.ID, updatedWager.ID)
		assert.Equal(t, wager.TotalWagerValue, updatedWager.TotalWagerValue)
		assert.Equal(t, wager.Odds, updatedWager.Odds)
		assert.Equal(t, wager.SellingPercentage, updatedWager.SellingPercentage)
		assert.Equal(t, wager.SellingPrice, updatedWager.SellingPrice)
		assert.Equal(t, wager.CurrentSellingPrice, updatedWager.CurrentSellingPrice)
		assert.Equal(t, wager.PercentageSold, updatedWager.PercentageSold)
		assert.Equal(t, wager.AmountSold, updatedWager.AmountSold)
	})
	t.Run("Test Query by id", func(t *testing.T) {
		s := CreateWagerRepositoryTextContext(db)
		wager := &entity.Wager{
			ID: 1,
		}
		actualWager, err := s.WagerRepository.QueryById(wager.ID)
		assert.Nil(t, err)
		assert.Equal(t, wager.ID, actualWager.ID)
	})
	t.Run("Test Query", func(t *testing.T) {
		s := CreateWagerRepositoryTextContext(db)
		actualWagers, err := s.WagerRepository.Query(0, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(actualWagers))
		assert.Equal(t, 1, actualWagers[0].ID)
	})
	tearDown()
}
func CreateWagerRepositoryTextContext(db *sql.DB) *WagerRepositoryTextContext {
	wagerRepository := NewWagerRepository(db)
	return &WagerRepositoryTextContext{
		WagerRepository: wagerRepository,
		db:              db,
	}
}
