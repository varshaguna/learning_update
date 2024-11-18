package main

import "fmt"

type Product struct {
	title string
	id string
	price string
}

func main() {
	prices := [4]float64{8.5, 12.2, 45.4, 88.7}
	fmt.Println(prices)
}
