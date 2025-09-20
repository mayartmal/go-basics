package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer demo")
	age := 32
	agePointer := &age
	aAge := getAdultYears(agePointer)
	fmt.Println("Adult age is", aAge)
	incrementAge(agePointer)
	fmt.Println("New age is", age)

}

func getAdultYears(agePointer *int) int {
	return *agePointer - 18
}

func incrementAge(agePointer *int) {
	*agePointer++
}
