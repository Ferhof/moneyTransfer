package main

import (
	"fmt"
)

const USDinEUR float64 = 1.18
const USDinRUB float64 = 76.75
const EURinRUB float64 = USDinEUR * USDinRUB

const EURO string = "EURO"
const USD string = "USD"
const RUB string = "RUB"

var userInput float64
var originalCurrency string
var targetCurrency string

func main() {
	fmt.Println("--- Добро пожаловать в конвертер валют ---")
	value, origin, target := getUserInput()
	result := convert(value, origin, target)
	fmt.Println("В рублях вы получите: ", result)
}

func getUserInput() (float64, string, string) {
	fmt.Print("Введите кол-во валюты для конвертации (0.00): ")
	fmt.Scan(&userInput)
	fmt.Print("Укажите валюту из которой конвертируем (EURO, USD, RUB): ")
	fmt.Scan(&originalCurrency)
	fmt.Print("Укажите валюту в которую конвертируем (EURO, USD, RUB): ")
	fmt.Scan(&targetCurrency)
	return userInput, originalCurrency, targetCurrency
}

func convert(value float64, originalCurrency string, targetCurrency string) float64 {
	switch originalCurrency {
	case EURO:
		switch targetCurrency {
		case USD:
			return value * USDinEUR
		case RUB:
			return value * EURinRUB

		}
	case USD:
		switch targetCurrency {
		case EURO:
			return value / USDinEUR
		case RUB:
			return value * USDinRUB
		}
	case RUB:
		switch targetCurrency {
		case EURO:
			return value / USDinEUR
		case USD:
			return value / USDinRUB
		}
	}
	return 0.0
}
