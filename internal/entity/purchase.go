package entity

import "time"

type Purchase struct {
	Id          int       `json:"id"`
	WagerId     int       `json:"wager_id"`
	BuyingPrice float64   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}
