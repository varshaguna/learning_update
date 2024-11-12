package main

import "fmt"

func main() {
	var revenue float64 
	var taxRate float64
	var expenses float64

	values("revenue: ")
	fmt.Scan(&revenue)

	values("taxRate: ")
	fmt.Scan(&taxRate)

	values("expenses: ")
	fmt.Scan(&expenses)
	
    ebt,profit,ratio := CalculateEarnings(revenue, taxRate, expenses)

	formatedEBT := fmt.Sprintf("ebt: %.1f\n",ebt)
	formatedPROFIT := fmt.Sprintf("profit: %.1f\n",profit)
	formatedRATIO := fmt.Sprintf("ratio: %.1f\n",ratio)

	fmt.Println(formatedEBT,formatedPROFIT,formatedRATIO)
	
}
func values(text string){
	fmt.Print(text)
}
func CalculateEarnings(revenue, taxRate, expenses float64) (float64,float64,float64){
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt,profit,ratio
}
