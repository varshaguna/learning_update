// Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
package main

import "fmt"

type Product struct {
	id string
	title string
	prices float64
}

func main() {
	hobbies := [3]string{"dancing", "planting", "skipping"}
	fmt.Println(hobbies)


//Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
     fmt.Println(hobbies[0])
	 fmt.Println(hobbies[1:3])

//      Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
      mainhobbies := hobbies[:2]
	  fmt.Println(mainhobbies)
    
//   Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
       fmt.Println(cap(mainhobbies))
	   mainhobbies = mainhobbies[1:3]
	   fmt.Println(mainhobbies)
    
    // Create a "dynamic array" that contains your course goals (at least 2 goals)

       courseGoals := []string{"Be positive", "Exercise Weekly Once"}
	   fmt.Println(courseGoals)

	 // Set the second goal to a different one AND then add a third goal to that existing dynamic array
	 
	   courseGoals[1] = "Exercise Daily"
	   courseGoals = append(courseGoals, "Be a GoodFit")
	   fmt.Println(courseGoals)
	 // Bonus: Create a "Product" struct with title, id, price and create a
     //		dynamic list of products (at least 2 products).
     //		Then add a third product to the existing list of products.
 
	 products := []Product{
		{
			"first pro",
			"A first product",
			9.2,
		},
		{
			"second pro",
			"A second product",
			7.3,
		},
	 }

	 fmt.Println(products)

	 newProduct := Product{
		"third pro",
		"A third product",
		9.3,
	 }

	 products = append(products, newProduct)

	 fmt.Println(products)
}
