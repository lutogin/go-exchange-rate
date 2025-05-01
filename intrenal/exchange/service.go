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

func (s *Service) GetRates(currency currency.Currency) (currency.CurrentRates, error) {
	data, err := s.cache.GetCurrencyRates(currency)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	if len(data) > 0 {
		return data, nil
	}

	res, err := s.repo.GetQuotes(currency)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	err = s.cache.SetCurrencyRates(currency, res.Quotes, 0)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return res.Quotes, nil
}
