package main

import "fmt"
import "time"

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}
func main(){
	userFirstName := getUserData("Enter the first name: ")
	userLastName := getUserData("Enter the last name: ")
	userBirthDate := getUserData("Enter the date of birth (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *user){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func getUserData(text string) (string){
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
