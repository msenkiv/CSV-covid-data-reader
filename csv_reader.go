package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// Covid Model diagram, that will provide the info to make data science
type Covid struct {
	Province  string
	City      string
	Country   string
	Confirmed int
	Deaths    int
	Recovered int
}

// Instance a Covid model
func createCovidData(record []string) Covid {
	province := record[0]
	country := record[1]
	confirmed, _ := strconv.Atoi(record[3])
	deaths, _ := strconv.Atoi(record[4])
	recovered, _ := strconv.Atoi(record[5])

	return Covid{province, "", country, confirmed, deaths, recovered}
}

// Instance a Covid model with second data model
func createDataFips(record []string) Covid {
	city := record[1]
	province := record[2]
	country := record[3]
	confirmed, _ := strconv.Atoi(record[7])
	deaths, _ := strconv.Atoi(record[8])
	recovered, _ := strconv.Atoi(record[9])

	return Covid{province, city, country, confirmed, deaths, recovered}

}

// Read a lot of files from path
func readFiles() []string {
	var files []string
	root := "/home/marvinsenkiv/go/src/github.com/marvinsenkiv/covid/files"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	files = removeElement(files)
	return files
}

// remove index from path
func removeElement(a []string) []string {
	i := 0
	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.
	return a

}

func main() {
	files := readFiles()
	for _, file := range files {
		f, _ := os.Open(file)
		r := csv.NewReader(f)

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			// Change data Layout
			if len(record) > 9 {
				covidData := createDataFips(record)
				covidJSON, _ := json.Marshal(covidData)
				fmt.Println(string(covidJSON))
			} else {
				covidData := createCovidData(record)
				covidJSON, _ := json.Marshal(covidData)
				fmt.Println(string(covidJSON))

			}

		}

	}
}
