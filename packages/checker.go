package main

import "fmt"

func main() {
	exchangeRates := make(map[string]float64)
	
	exchangeRates["USD"] = 1.0
	exchangeRates["EUR"] = 0.92
	exchangeRates["GBP"] = 0.79
	
	target := "JPY"
	
	findRate, found := exchangeRates[target]
	
	if found {
	 fmt.Println("JPY: ", findRate)
	} else {
	 fmt.Printf("Warning: %s not supported.\n", target)
	}
}
