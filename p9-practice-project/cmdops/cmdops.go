package cmdops

import "fmt"

type Cmdop struct {
}

func New() Cmdop {
	return Cmdop{}
}

func (fo Cmdop) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices with ENTER confirmation")
	var prices []string 
	for {
		var price string
		fmt.Print("Price:")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (fo Cmdop) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
