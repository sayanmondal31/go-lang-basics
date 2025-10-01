package prices

import (
	"fmt"
	"profit_calc/conversion"
	"profit_calc/iomanager"
)

type TaxInludedPriceJob struct {
	IOMananager       iomanager.IOMananager `json:"-"` // -> this means ignore won't sjow
	TaxRate           float64               `json:"tax_rate"`
	InputPrices       []float64             `json:input_prices`
	TaxIncludedPrices map[string]string     `json:"tax_rates`
}

// constructor
func NewTaxInludedPriceJob(iom iomanager.IOMananager, taxrate float64) *TaxInludedPriceJob {
	fmt.Println(taxrate, "<-taxrate")
	return &TaxInludedPriceJob{
		IOMananager: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxrate,
	}
}

// methods
func (job *TaxInludedPriceJob) Process(donechan chan bool, errChan chan error) {
	err := job.LoadData()

	if err != nil {
		// return err
		errChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOMananager.WriteResult(job)
	donechan <- true

}

func (job *TaxInludedPriceJob) LoadData() error {
	lines, err := job.IOMananager.ReadLines()

	if err != nil {

		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {

		return err
	}

	job.InputPrices = prices

	return nil

}
