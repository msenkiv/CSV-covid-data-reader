package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReadingCSV(t *testing.T) {
	f, _ := os.Open("files/01-22-2020.csv")
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record[11])

	}

}
