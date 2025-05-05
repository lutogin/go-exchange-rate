package cache

import (
	"database/sql"
	"encoding/json"
	"exchange-rate/intrenal/currency"
	"time"
)

type Cache struct {
	db *sql.DB
}

func NewCache(db *sql.DB) *Cache {
	return &Cache{db: db}
}

func (c *Cache) SetCurrencyRates(currency currency.Currency, quotes currency.CurrentRates, ttl time.Duration) error {
	if ttl == 0 {
		ttl = 60 * time.Minute // Default TTL
	}

	// Serialize quotes to JSON
	quotesJSON, err := json.Marshal(quotes)
	if err != nil {
		return err
	}

	expiration := time.Now().Add(ttl).Unix()
	_, err = c.db.Exec(`
		INSERT OR REPLACE INTO cache (currency, quotes, expiration)
		VALUES (?, ?, ?)
	`, currency, quotesJSON, expiration)
	return err
}

func (c *Cache) GetCurrencyRates(cur currency.Currency) (currency.CurrentRates, error) {
	var quotesJSON string
	var expiration time.Time

	err := c.db.QueryRow(`
		SELECT quotes, expiration FROM cache WHERE currency = ?
	`, cur).Scan(&quotesJSON, &expiration)
	if err == sql.ErrNoRows {
		return currency.CurrentRates{}, nil
	} else if err != nil {
		return currency.CurrentRates{}, err
	}

	// check if the cached rate is expired
	if time.Now().Unix() > expiration.Unix() {
		go func() {
			_ = c.DeleteCurrencyRate(cur)
		}()
		return currency.CurrentRates{}, nil
	}

	// Parse the JSON string into CurrentRates
	quotes := currency.CurrentRates{}
	err = json.Unmarshal([]byte(quotesJSON), &quotes)
	if err != nil {
		return currency.CurrentRates{}, err
	}

	return quotes, nil
}

func (c *Cache) DeleteCurrencyRate(currency currency.Currency) error {
	_, err := c.db.Exec(`
		DELETE FROM cache WHERE currency = ?
	`, currency)
	return err
}
