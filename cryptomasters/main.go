package main

import (
	"cryptomasters/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup // wait group is a value that will wait for a go routine to finish

	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
	// go getCurrencyData("ETH")
	// go getCurrencyData("BCH")
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	// print(rate.Price, err)
	fmt.Printf("the rate for %v is %.2f \n", rate.Currency, rate.Price)
	if err != nil {
		panic("panic")
	}
}
