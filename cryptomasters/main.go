package main

import "fe.com/go/crypto/api"

func main() {
	rate, err := api.GetRate("BTC")

	print(rate.Price, err)
}
