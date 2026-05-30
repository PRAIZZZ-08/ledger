package main

import (
	"errors"
	"fmt"
)

func calculateTotal(amount float64, taxRate float64) float64 {
	return amount + amount*taxRate
}

func validateUser(username string) (bool, error) {
	if username == "" {
		return false, errors.New("username cannot be empty")
	}
	return true, nil
}

func main() {
	var transactionCount uint16 = 500
	var totalValue float64 = 1250.75
	var currencyCode string = "USD"

	fmt.Printf("Transactions: %d | Total: %.2f | Currency: %s\n", transactionCount, totalValue, currencyCode)

	total := calculateTotal(100.0, 0.05)

	fmt.Println(total)

	namecheck, err := validateUser("")

	if err != nil {
		fmt.Println("Error found:", err)
		return
	}

	fmt.Println("Status:", namecheck)
}
