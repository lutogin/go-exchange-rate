package calculation

import (
	"exchange-rate/intrenal/currency"
	"exchange-rate/intrenal/exchange"
	"fmt"
)

type Service struct {
	exService *exchange.Service
}

func NewService(exService *exchange.Service) *Service {
	return &Service{exService}
}

func (calcService *Service) Calculate(from, to currency.Currency, amount float64) (float64, error) {
	// Get the exchange rate from the service
	rateChan := make(chan currency.CurrentRates)
	errChan := make(chan error)

	go func() {
		rate, err := calcService.exService.GetRates(from)
		if err != nil {
			errChan <- err
		}
		rateChan <- rate
	}()

	select {
	case err := <-errChan:
		return 0, fmt.Errorf("failed to get exchange rate: %w", err)
	case rate := <-rateChan:
		{
			if toCurrencyRate, ok := rate[currency.PairCurrency(fmt.Sprintf("%s%s", from, to))]; !ok {
				return 0, fmt.Errorf("exchange rate not found for %s to %s", from, to)
			} else {
				// Calculate the converted amount
				convertedAmount := toCurrencyRate * amount
				return convertedAmount, nil
			}
		}
	}
}
