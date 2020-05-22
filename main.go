package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/marvinsenkiv/covid/services"
)

func main() {
	files := services.ReadFiles()
	for _, file := range files {
		f, _ := os.Open(file)
		r := csv.NewReader(f)
		fmt.Println(file)
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			// Change data Layout
			if len(record) > 9 {
				covidData := services.CreateDataFips(record)
				services.InsertCovidData(covidData, file)

			} else {
				covidData := services.CreateCovidData(record)
				services.InsertCovidData(covidData, file)

			}

		}

	}
}
