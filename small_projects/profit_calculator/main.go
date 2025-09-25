package main

import (
	"fmt"
	"profit_calc/cmdmanager"
	"profit_calc/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("price.txt", fmt.Sprintf("result_%0.2f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxInludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("couldn't process job")
			fmt.Println(err)
		}

	}

}
