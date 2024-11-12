package main

import ("fmt"
       "os"
	   "strconv"
)

const accountBalanceFile = "file.txt"
func getBalanceFromFile() float64{
	data, _ := os.ReadFile(accountBalanceFile)
	fileText := string(data)
	file, _ := strconv.ParseFloat(fileText, 64)
	return file
}
func writeBalanceToFile(file float64){
	fileText := fmt.Sprint(file)
	os.WriteFile("file.txt", []byte(fileText),0644)
}
func main(){
 
	var accountBalance = getBalanceFromFile()
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice :")
	fmt.Scan(&choice)

	
	if choice == 1{
		fmt.Print("Your Balance is ",accountBalance)
		return
	}else if choice == 2{
		fmt.Print("Your Deposit is: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0{
			fmt.Println("Invalid Amount. Must be greater than 0.")
			return
		}
		accountBalance=accountBalance+depositAmount
		fmt.Println("Balance updated! new amount: ", accountBalance)
		writeBalanceToFile(accountBalance)
	}else if choice == 3{
		fmt.Print("Enter the amount to Withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0{
			fmt.Println("Invalid Amount. Must be greater than 0.")
			return
		}else if withdrawAmount > accountBalance{
			fmt.Println("You can't Withdraw more than you have")
			return
		}
		accountBalance=accountBalance-withdrawAmount
		fmt.Println("Balance updated! new amount: ", accountBalance)
		writeBalanceToFile(accountBalance)
	}else if choice == 4{
		fmt.Print("Good Bye!, Visit again")
	}
}
