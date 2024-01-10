package main

import (
	"fmt"
	"sync"

	"fe.com/go/crypto/api"
)

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup // creating a waiting group for keeping the main thread open
	for _, currency := range currencies {
		wg.Add(1) //add 1 for each element of the slice
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done() //substract 1 when the function is done
		}(currency) //IIFE in a go routine to prevent blocking the main thread
	}
	wg.Wait() // wait for the counter to go zero before closing the main go thread
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f USD\n", rate.Currency, rate.Price)
	}
}
