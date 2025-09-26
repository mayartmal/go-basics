package lists

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{1.5, 7.4}
	fmt.Println(prices)
	prices = append(prices, 5.99)
	fmt.Println(prices)
	var products [4]string
	fmt.Println(products)

	/*
		var products [4]string = [4]string{"Book"}
		products[1] = "Plate"
		prices := [4]float64{1.1, 1.3, 5.6, 3.8}
		featuredPrices := prices[1:]
		highLightedPrices := featuredPrices[:1]
		fmt.Println(highLightedPrices)
		fmt.Println(len(highLightedPrices))
		fmt.Println(cap(highLightedPrices))
		fmt.Println(cap(prices))

		// fmt.Println(prices)
		// 	fmt.Println(products)
		// 	fmt.Println(prices[1:2])
		// 	fmt.Println(prices[:2])
		// 	fmt.Println(prices[1:4])
		// }

		highLightedPrices = highLightedPrices[0:3]
		fmt.Println(highLightedPrices)
	*/

}
