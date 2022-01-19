package repository

import (
	"database/sql"
	"wagers/internal/entity"
)

func NewPurchaseRepository(db *sql.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

type PurchaseRepository interface {
	Create(pc *entity.Purchase) (*entity.Purchase, error)
	QueryById(purchaseId int) (*entity.Purchase, error)
}
type purchaseRepository struct {
	db *sql.DB
}

func (r *purchaseRepository) Create(pc *entity.Purchase) (*entity.Purchase, error) {
	rows, err := r.db.Exec(
		"INSERT INTO purchases (wager_id, buying_price) VALUES(?,?)",
		pc.WagerId, pc.BuyingPrice,
	)
	if err != nil {
		return nil, err
	}

	lastIndex, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.QueryById(int(lastIndex))
}

func (r *purchaseRepository) QueryById(wagerId int) (*entity.Purchase, error) {
	rows, err := r.db.Query(
		"SELECT * FROM purchases WHERE id = ?",
		wagerId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	purchase := &entity.Purchase{}

	err = rows.Scan(&purchase.Id, &purchase.WagerId, &purchase.BuyingPrice, &purchase.BoughtAt)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}
