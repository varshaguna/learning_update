package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{8.5, 12.2, 45.4, 88.7}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])
}