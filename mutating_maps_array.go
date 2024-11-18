package main 

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon Web Services" : "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"]) //extract values with the help of keys

	websites["LinkedIn"] = "https://linkedin.com" //adding new websites
	fmt.Println(websites)

	delete(websites, "Google") //deleting new websites
	fmt.Println(websites)
}
