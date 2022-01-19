package repository

import (
	"database/sql"
	"wagers/internal/entity"
)

func NewWagerRepository(db *sql.DB) WagerRepository {
	return &wagerRepository{db: db}
}

type WagerRepository interface {
	Create(wg *entity.Wager) (*entity.Wager, error)
	Update(wager *entity.Wager) error
	QueryById(wagerId int) (*entity.Wager, error)
	Query(offset int, limit int) ([]*entity.Wager, error)
}
type wagerRepository struct {
	db *sql.DB
}

func (r *wagerRepository) Create(wg *entity.Wager) (*entity.Wager, error) {
	rows, err := r.db.Exec(
		"INSERT INTO wagers (total_wager_value, odds, selling_percentage, selling_price, current_selling_price, percentage_sold, amount_sold) VALUES(?,?,?,?,?,?,?)",
		wg.TotalWagerValue, wg.Odds, wg.SellingPercentage, wg.SellingPrice, wg.CurrentSellingPrice, wg.PercentageSold, wg.AmountSold,
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

func (r *wagerRepository) Update(wager *entity.Wager) error {
	rows, err := r.db.Query(
		"UPDATE wagers SET total_wager_value = ?, odds = ?, selling_percentage = ?, selling_price = ?, current_selling_price = ?, percentage_sold = ?, amount_sold = ? WHERE id = ?",
		wager.TotalWagerValue, wager.Odds, wager.SellingPercentage, wager.SellingPrice, wager.CurrentSellingPrice, wager.PercentageSold, wager.AmountSold, wager.ID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
func (r *wagerRepository) QueryById(wagerId int) (*entity.Wager, error) {
	rows, err := r.db.Query(
		"SELECT * FROM wagers WHERE id = ?",
		wagerId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	wager := &entity.Wager{}

	err = rows.Scan(&wager.ID, &wager.TotalWagerValue, &wager.Odds,
		&wager.SellingPercentage, &wager.SellingPrice, &wager.CurrentSellingPrice,
		&wager.PercentageSold, &wager.AmountSold, &wager.PlacedAt,
	)
	if err != nil {
		return nil, err
	}

	return wager, nil
}
func (r *wagerRepository) Query(offset int, limit int) ([]*entity.Wager, error) {
	rows, err := r.db.Query(
		"SELECT * FROM wagers LIMIT ? OFFSET ?",
		limit, offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wagers := []*entity.Wager{}

	for rows.Next() {
		wager := &entity.Wager{}

		err = rows.Scan(&wager.ID, &wager.TotalWagerValue, &wager.Odds,
			&wager.SellingPercentage, &wager.SellingPrice, &wager.CurrentSellingPrice,
			&wager.PercentageSold, &wager.AmountSold, &wager.PlacedAt,
		)
		if err != nil {
			return nil, err
		}

		wagers = append(wagers, wager)
	}

	return wagers, nil
}
