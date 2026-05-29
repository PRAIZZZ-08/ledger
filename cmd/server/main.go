package main

import "fmt"

func main() {
	var transactionCount uint16 = 500
	var totalValue float64 = 1250.75
	var currencyCode string = "USD"

	fmt.Printf("Transactions: %d | Total: %.2f | Currency: %s\n", transactionCount, totalValue, currencyCode)
}
