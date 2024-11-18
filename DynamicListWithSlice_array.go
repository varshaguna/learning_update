package main

import "fmt"

func main() {
	prices := []float64{10.99,8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.89

	prices = append(prices, 5.90)
	prices = prices[1:] //omit 1st element
	fmt.Println(prices)
}
