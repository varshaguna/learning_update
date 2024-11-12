package main

import "fmt"

func main() {
	var revenue float64 
	var taxRate float64
	var expenses float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)
	fmt.Print("expenses: ")
	fmt.Scan(&expenses)
	
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	
}
