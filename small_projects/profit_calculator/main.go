package main

import (
	"fmt"
	"profit_calc/filemanager"
	"profit_calc/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	doneChnas := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChnas[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.2f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxInludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChnas[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("couldn't process job")
		// 	fmt.Println(err)
		// }

	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}

		case <-doneChnas[index]:
			fmt.Println("Done!")

		}
	}

}
