package main

import "fmt"

type trasformFn func(int) int

func main(){
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{4, 1, 7}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	trasformFn1 := getTransformerFunction(&numbers)
	trasformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, trasformFn1)
	moretransformedNumbers := transformNumbers(&moreNumbers, trasformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moretransformedNumbers)
}

func transformNumbers(numbers *[]int, transform trasformFn) []int{
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) trasformFn {
	if (*numbers)[0] == 1 {
		return double
	}else{
		return triple
	}
}

func double(number int) int{
	return number*2
}

func triple(number int) int{
	return number*3
}
