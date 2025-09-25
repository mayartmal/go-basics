package main

import (
	"fmt"

	"example.com/practice/product"
)

func main() {
	// task 1
	fmt.Println("----Task1----")
	hobbies := [3]string{"running", "learning", "hiking"}
	fmt.Println(hobbies)
	fmt.Println()

	// task 2
	fmt.Println("----Task2----")
	topHobbie := hobbies[0:1]
	secondaryHobbies := hobbies[1:]
	fmt.Println(topHobbie)
	fmt.Println(secondaryHobbies)
	fmt.Println()

	// task 3
	fmt.Println("----Task3----")
	fmt.Println(hobbies)

	w1 := hobbies[0:2]
	w2 := hobbies[:2]
	fmt.Println(w1)
	fmt.Println(w2)
	fmt.Println()

	// task 4
	fmt.Println("----Task4----")
	w3 := hobbies[1:3]
	fmt.Println(w3)
	fmt.Println()

	// task 5
	fmt.Println("----Task5----")
	goals := []string{"learn go", "learn backend"}
	fmt.Println(goals)
	goals[1] = "learn backend strongly"
	fmt.Println(goals)
	goals = append(goals, "practice web techs")
	fmt.Println(goals)
	fmt.Println()

	// task 6
	fmt.Println("----Task5----")
	products := []product.Product{product.New("Milk", 3.50, 0), product.New("Bread", 1.20, 1)}
	fmt.Println(products)
	products = append(products, product.New("Candy", 2.00, 2))
	fmt.Println(products)
	fmt.Println()

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
