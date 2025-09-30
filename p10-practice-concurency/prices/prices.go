package prices

import (
	"fmt"

	"example.com/price-calculator/convertion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"output_prices"`
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

func (j *TaxIncludedPricesJob) Process(doneChan chan bool, errorChan chan error) {
	err := j.LoadData()
	if err != nil {
		// return err
		errorChan <- err
		return
	}
	result := make(map[string]string)
	for _, price := range j.InputPrices {
		taxIncludedPrice := price + (price * j.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	j.TaxIncludedPrices = result
	j.IOManager.WriteResult(j)
	doneChan <- true

}

func (j *TaxIncludedPricesJob) LoadData() error {
	lines, err := j.IOManager.ReadLines()
	if err != nil {
		// fmt.Println("ERROR DURING FILEOPS")
		// fmt.Println(err)
		return err
	}

	loadedPrices, err := convertion.Strings2Floats(lines)
	if err != nil {
		// fmt.Println("ERROR DURING CONVERTION")
		// fmt.Println(err)
		return err
	}

	j.InputPrices = loadedPrices
	return nil

}
