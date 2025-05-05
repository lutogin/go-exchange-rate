package utils

import (
	"bufio"
	"exchange-rate/intrenal/_errors"
	"os"
)

func ParseUserInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", _errors.InvalidInputError
	}
	if scanner.Text() == "" {
		return "", _errors.EmptyInputError
	}
	return scanner.Text(), nil
}
