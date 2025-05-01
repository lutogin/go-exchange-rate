package exchange

import "exchange-rate/intrenal/currency"

type Repository interface {
	GetQuotes(source currency.Currency) (GetRatesResponse, error)
}
