package di

import (
	"exchange-rate/intrenal/app"
	"exchange-rate/intrenal/cache"
	"exchange-rate/intrenal/calculation"
	"exchange-rate/intrenal/config"
	"exchange-rate/intrenal/exchange"
	"exchange-rate/intrenal/input"
	"exchange-rate/intrenal/storage"
	"fmt"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Provide dependencies
	container.Provide(func() *config.Config {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Errorf("failed to load config: %w", err)
		}
		return cfg
	})
	container.Provide(func(cfg *config.Config) *exchange.Client {
		return exchange.New(cfg.ApiKey)
	})
	container.Provide(func() *storage.Storage {
		return storage.New()
	})
	container.Provide(func(store *storage.Storage) *cache.Cache {
		return cache.NewCache(store.Db)
	})
	container.Provide(func(client *exchange.Client, cache *cache.Cache) *exchange.Service {
		return exchange.NewService(client, cache)
	})
	container.Provide(func(exService *exchange.Service) *calculation.Service {
		return calculation.NewService(exService)
	})
	container.Provide(func() *input.ParsingService {
		return input.NewParsingService()
	})
	container.Provide(func(calcService *calculation.Service, inputService *input.ParsingService) *app.App {
		return app.New(calcService, inputService)
	})

	return container
}
