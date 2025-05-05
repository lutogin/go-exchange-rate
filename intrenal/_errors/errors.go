package _errors

import "errors"

var InvalidInputError = errors.New("invalid user input")
var EmptyInputError = errors.New("empty user input")
var WrongNumberOfArgumentsError = errors.New("wrong number of arguments")
var InvalidCurrencyError = errors.New("invalid currency")
