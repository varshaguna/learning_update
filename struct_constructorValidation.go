package main

import "fmt"
import "time"
import "errors"

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstName, lastName and birthDate are required")
	}
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main(){
	userFirstName := getUserData("Enter the first name: ")
	userLastName := getUserData("Enter the last name: ")
	userBirthDate := getUserData("Enter the date of birth (MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)
     
	if err != nil {
		fmt.Println(err)
		return
	}
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
	fmt.Scanln(&value)
	return value
}
