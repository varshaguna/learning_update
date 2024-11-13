package main

import "fmt"

func main(){
	age := 32

	var userAge *int
	userAge = &age

	fmt.Println("Age:", *userAge)

	adultYears := getAdultYears(userAge)
	fmt.Println(adultYears) 

}

func getAdultYears(age *int) int{
	return *age - 18
}
