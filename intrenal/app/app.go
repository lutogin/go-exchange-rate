package app

import (
	"exchange-rate/intrenal/calculation"
	"exchange-rate/intrenal/input"
	"fmt"
)

type App struct {
	calc  *calculation.Service
	input *input.ParsingService
}

func New(calcService *calculation.Service, inputService *input.ParsingService) *App {
	return &App{
		calc:  calcService,
		input: inputService,
	}
}

func (app *App) Run() error {
	fmt.Print("Enter data to convert (example: '100 USD to EUR'): ")
	// Parse user input
	if err := app.input.Parse(); err != nil {
		return fmt.Errorf("context: %w", err)
	}

	// Call the calculation service
	result, err := app.calc.Calculate(app.input.FromCurrency, app.input.ToCurrency, app.input.Amount)
	if err != nil {
		return fmt.Errorf("context: %w", err)
	}

	fmt.Println(result)
	return nil
}
