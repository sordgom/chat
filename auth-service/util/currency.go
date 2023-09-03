package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	JPY = "JPY"
	GBP = "GBP"
	AUD = "AUD"
	CHF = "CHF"
	CNY = "CNY"
	SEK = "SEK"
	NZD = "NZD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, JPY, GBP, AUD, CHF, CNY, SEK, NZD:
		return true
	}
	return false
}
