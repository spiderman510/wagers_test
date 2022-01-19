package entity

type ListRequest struct {
	Page  int `form:"page" binding:"required"`
	Limit int `form:"limit" binding:"required"`
}

type PlaceRequest struct {
	TotalWagerValue   int     `json:"total_wager_value"`
	Odds              int     `json:"odds"`
	SellingPercentage int     `json:"selling_percentage"`
	SellingPrice      float64 `json:"selling_price"`
}
type BuyRequest struct {
	BuyingPrice float64 `json:"buying_price"`
	WagerId     int
}
