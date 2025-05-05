package app

import (
	"errors"
	"exchange-rate/intrenal/_errors"
	"exchange-rate/intrenal/calculation"
	"exchange-rate/intrenal/input"
	"fmt"
	"os"
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
	for {
		fmt.Print("Enter data to convert (example: '100 USD to EUR'): ")
		// Parse user input
		if err := app.input.Parse(); err != nil {
			if errors.Is(err, _errors.EmptyInputError) {
				os.Exit(0)
			}
			return fmt.Errorf("context: %w", err)
		}

		// Call the calculation service
		resultChan := make(chan float64)
		errChan := make(chan error)
		go func() {
			defer close(resultChan)
			defer close(errChan)

			result, err := app.calc.Calculate(app.input.FromCurrency, app.input.ToCurrency, app.input.Amount)
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- result
		}()

		select {
		case result := <-resultChan:
			fmt.Println(result)
		case err := <-errChan:
			return fmt.Errorf("calculation error: %w", err)
		}
	}

	return nil
}
