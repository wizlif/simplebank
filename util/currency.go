package util

// Constants for all supported currencies
const (
	UGX = "UGX"
	USD = "USD"
	EUR = "EUR"
)

// Is supported currency return true if supported currency elese returns false
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, UGX, EUR:
		return true
	}
	return false
}
