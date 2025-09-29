package main

import (
	"fmt"

	"example.com/price-calculator/fileops"
	"example.com/price-calculator/prices"
	// "example.com/price-calculator/cmdops"
)

func main() {
	taxRates := []float64{0, 0.18, 0.20, 0.22}
	// result := make(map[float64][]float64)

	for _, tax := range taxRates {
		fo := fileops.New("pricwes.txt", fmt.Sprintf("results_%.0f.json", tax*100))
		// co := cmdops.New()
		pj := prices.New(fo, tax)
		err := pj.IncludeTaxes()
		if err != nil {
			fmt.Println("CANT BE CALCULATED")
			fmt.Println(err)
		}
	}

}
