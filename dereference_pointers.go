package main

import "fmt"

func main(){
	age := 32

	var userAge *int
	userAge = &age

	fmt.Println("Age:", *userAge)

	getAdultYears(&age)
	fmt.Println(age) 

}

func getAdultYears(age *int){
	//return *age - 18
	*age = *age -18
}
