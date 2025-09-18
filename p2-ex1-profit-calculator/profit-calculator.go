package main

import "fmt"

func main() {
	revenue := 0.0 // befor expenses
	expenses := 0.0
	taxRate := 0.0

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("TaxRate: ")
	fmt.Scan(&taxRate)

	tax := revenue * taxRate

	erningsBeforeTax := revenue - expenses
	erningsAfterTax := revenue - tax - expenses

	fmt.Print("Earnings before tax: ")
	fmt.Print(erningsBeforeTax)
	fmt.Print()
	fmt.Print("Earnings after tax: ")
	fmt.Print(erningsAfterTax)
}
