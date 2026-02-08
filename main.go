package main

import "fmt"

const USDinEUR float64 = 1.18
const USDinRUB float64 = 76.75
const EURinRUB float64 = USDinEUR * USDinRUB

func main() {
	var userInput float64
	fmt.Print("Введите кол-во евро для конвертации: ")
	fmt.Scan(&userInput)
	fmt.Println("В рублях вы получите: ", userInput*EURinRUB)
}
