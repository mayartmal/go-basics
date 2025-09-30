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
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for i, tax := range taxRates {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)

		fo := fileops.New("prices.txt", fmt.Sprintf("results_%.0f.json", tax*100))
		// co := cmdops.New()
		pj := prices.New(fo, tax)
		go pj.Process(doneChans[i], errChans[i])

		// if err != nil {
		// 	fmt.Println("CANT BE CALCULATED")
		// 	fmt.Println(err)
		// }
	}

	for i, _ := range taxRates {
		select {
		case err := <-errChans[i]:
			if err != nil {
				fmt.Println("CANT BE CALCULATED")
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}



}
