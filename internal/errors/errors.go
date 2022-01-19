package errors

import "github.com/gin-gonic/gin"

type Errors struct {
	Code    int
	Message interface{}
}

var TotalWagerValueRequiredError = gin.H{"error": "Total wager value must be greater than 0"}
var OddsRequiredError = gin.H{"error": "Odds must be greater than 0"}
var SellingPercentageRangeError = gin.H{"error": "Selling percentage must be between 1 and 100"}
var SellingPriceRequiredError = gin.H{"error": "Selling price must be greater than total wager value * (selling percentage / 100)"}

var BuyingPriceRequiredError = gin.H{"error": "Buying price must be greater than 0"}
var BuyingPriceValueError = gin.H{"error": "Buying price must be lesser or equal to current selling price"}

var PageRequiredError = gin.H{"error": "Page should be greater than 0"}
var LimitRequiredError = gin.H{"error": "Limit should be greater than 0"}

var WagerNotFoundError = gin.H{"error": "Wager is not found"}

var ListWagerError = gin.H{"error": "Failed to get wager list"}
var WagerGetError = gin.H{"error": "Failed to get wager"}
var InsertDataError = gin.H{"error": "Failed to create data"}
var WagerUpdatedError = gin.H{"error": "Failed to update wager"}

var InternalServerError = gin.H{"error": "Interal server error"}
