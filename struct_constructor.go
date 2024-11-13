package main

import "fmt"
import "time"

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}


func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main(){
	userFirstName := getUserData("Enter the first name: ")
	userLastName := getUserData("Enter the last name: ")
	userBirthDate := getUserData("Enter the date of birth (MM/DD/YYYY): ")

	var appUser *user

	appUser = newUser(userFirstName, userLastName, userBirthDate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func (u *user) outputUserDetails(){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func(u *user) clearUserName(){
	u.firstName = ""
	u.lastName = ""
}


func getUserData(text string) (string){
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
