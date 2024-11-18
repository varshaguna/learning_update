package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{8.5, 12.2, 45.4, 88.7}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.23
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
