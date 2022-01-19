package repository

import (
	"database/sql"
	"testing"
	"wagers/internal/entity"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type PurchaseRepositoryTextContext struct {
	PurchaseRepository PurchaseRepository
	db                 *sql.DB
}

func CreatePurchaseRepositoryTextContext() *PurchaseRepositoryTextContext {
	db, err := sql.Open(driverName, mockDatabaseUrl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connection database")
	}
	return &PurchaseRepositoryTextContext{
		PurchaseRepository: NewPurchaseRepository(db),
		db:                 db,
	}
}
func TestPurchaseRepository(t *testing.T) {
	initDb()
	t.Run("Test Create", func(t *testing.T) {
		s := CreatePurchaseRepositoryTextContext()
		purchase := &entity.Purchase{
			WagerId:     1,
			BuyingPrice: float64(20.0),
		}
		lastIndexId := getLastIndexId(s.db, purchaseTableName)
		actualPurchase, err := s.PurchaseRepository.Create(purchase)
		assert.Nil(t, err)
		assert.Equal(t, lastIndexId, actualPurchase.Id)
		assert.Equal(t, purchase.WagerId, actualPurchase.WagerId)
		assert.Equal(t, purchase.BuyingPrice, actualPurchase.BuyingPrice)
	})
	t.Run("Test Query by id", func(t *testing.T) {
		s := CreatePurchaseRepositoryTextContext()
		purchase := &entity.Purchase{
			Id: 1,
		}
		actualPurchase, err := s.PurchaseRepository.QueryById(purchase.Id)
		assert.Nil(t, err)
		assert.Equal(t, purchase.Id, actualPurchase.Id)
	})
	tearDown()
}
