package entity

import (
	"time"
)

type Wager struct {
	ID                  int       `json:"id"`
	TotalWagerValue     int       `json:"total_wager_value"`
	Odds                int       `json:"odds"`
	SellingPercentage   int       `json:"selling_percentage"`
	SellingPrice        float64   `json:"selling_price"`
	CurrentSellingPrice float64   `json:"current_selling_price"`
	PercentageSold      *float64  `json:"percentage_sold"`
	AmountSold          *int      `json:"amount_sold"`
	PlacedAt            time.Time `json:"placed_at"`
}
