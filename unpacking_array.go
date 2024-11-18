package main

import "fmt"

func main() {
	prices := []float64{10.99,8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.89

	prices = append(prices, 5.90, 67.9, 76.4, 34.2)
	prices = prices[1:] //omit 1st element
	fmt.Println(prices)

	discountPrices := []float64{656.55, 87.89, 76.34}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}
