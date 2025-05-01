package exchange

import "exchange-rate/intrenal/currency"

type GetRatesResponse struct {
	Success   bool                  `json:"success"`
	Terms     string                `json:"terms"`
	Privacy   string                `json:"privacy"`
	Timestamp int64                 `json:"timestamp"`
	Source    string                `json:"source"`
	Quotes    currency.CurrentRates `json:"quotes"`
}
