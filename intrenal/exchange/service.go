package exchange

import (
	"exchange-rate/intrenal/cache"
	"exchange-rate/intrenal/currency"
	"fmt"
)

type Service struct {
	repo  Repository
	cache *cache.Cache
}

func NewService(repo Repository, cache *cache.Cache) *Service {
	return &Service{repo, cache}
}

func (s *Service) GetRates(cur currency.Currency) (currency.CurrentRates, error) {
	dataChan := make(chan currency.CurrentRates)
	errChan := make(chan error)
	defer close(dataChan)
	defer close(errChan)
	go func() {
		data, err := s.cache.GetCurrencyRates(cur)
		if err != nil {
			errChan <- err
			return
		}
		if len(data) > 0 {
			dataChan <- data
		} else {
			dataChan <- nil
		}
	}()

	select {
	case err := <-errChan:
		return nil, fmt.Errorf("failed to get exchange rate from cache: %w", err)
	case data := <-dataChan:
		if len(data) > 0 {
			return data, nil
		}
	}

	quotesChan := make(chan currency.CurrentRates)
	defer close(quotesChan)
	go func() {
		res, err := s.repo.GetQuotes(cur)
		if err != nil {
			errChan <- err
			return
		}
		if len(res.Quotes) > 0 {
			quotesChan <- res.Quotes
		} else {
			quotesChan <- nil
		}
	}()

	select {
	case err := <-errChan:
		return nil, fmt.Errorf("failed to get exchange rate from repository: %w", err)
	case quotes := <-quotesChan:
		go func() {
			errChan <- s.cache.SetCurrencyRates(cur, quotes, 0)
		}()

		select {
		case err := <-errChan:
			if err != nil {
				return nil, fmt.Errorf("failed to set currency rates in cache: %w", err)
			}
		}

		return quotes, nil
	}
}
