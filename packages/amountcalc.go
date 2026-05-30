package main

import "fmt"

func calculateTotal(amount float64, taxRate float64) float64 {
	return amount + amount*taxRate
}

func main() {
	
	var transactionCount uint16 = 500
	var totalValue float64 = 1250.75
	var currencyCode string = "USD"

	fmt.Printf("Transactions: %d | Total: %.2f | Currency: %s\n", transactionCount, totalValue, currencyCode)

	total := calculateTotal(100.0, 0.05)

	fmt.Println(total)
}
