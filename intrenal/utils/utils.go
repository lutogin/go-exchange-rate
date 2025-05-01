package utils

import (
	"bufio"
	"exchange-rate/intrenal/errors"
	"os"
)

func ParseUserInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", errors.InvalidInputError
	}
	return scanner.Text(), nil
}
