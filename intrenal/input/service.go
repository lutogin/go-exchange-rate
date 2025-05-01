package input

import (
	"exchange-rate/intrenal/currency"
	"exchange-rate/intrenal/utils"
	"fmt"
	"strconv"
	"strings"
)

type ParsingService struct {
	Amount       float64
	FromCurrency currency.Currency
	ToCurrency   currency.Currency
}

func NewParsingService() *ParsingService {
	return &ParsingService{}
}

func (p *ParsingService) Parse() error {
	input, err := utils.ParseUserInput()
	if err != nil {
		return fmt.Errorf("context: %w", err)
	}

	parts := strings.Split(input, " ")

	if len(parts) != 4 {
		return fmt.Errorf("invalid input format, expected 'amount currency TO currency'")
	}

	p.Amount, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %w", err)
	}

	fromCurrency := strings.ToUpper(parts[1])
	if err = currency.CheckCurrency(fromCurrency); err != nil {
		// todo: add a link to the supported currencies. remove duplications
		return fmt.Errorf("invalid currency: %s. supported curencies https://exchangerate.host/currencies context: %s", parts[1], err.Error())
	}
	p.FromCurrency = currency.Currency(fromCurrency)

	if strings.ToLower(parts[2]) != "to" {
		return fmt.Errorf("invalid to: %s", parts[2])
	}

	toCurrency := strings.ToUpper(parts[3])
	if err = currency.CheckCurrency(toCurrency); err != nil {
		return fmt.Errorf("invalid currency: %s. supported curencies https://exchangerate.host/currencies context: %s", parts[3], err.Error())
	}
	p.ToCurrency = currency.Currency(toCurrency)

	return nil
}
