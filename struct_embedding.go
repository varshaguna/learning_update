//struct.go

package main

import "fmt"
import "example.com/first.app/user"

func main(){
	userFirstName := getUserData("Enter the first name: ")
	userLastName := getUserData("Enter the last name: ")
	userBirthDate := getUserData("Enter the date of birth (MM/DD/YYYY): ")

	var appUser *user.User

	appUser ,err := user.NewUser(userFirstName, userLastName, userBirthDate)
     
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")
	
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()



	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(text string) (string){
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)
	return value
}

//user.go

package user

import "fmt"
import "errors"
import "time"

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User 
}


func (u *User) OutputUserDetails(){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func(u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstName, lastName and birthDate are required")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
