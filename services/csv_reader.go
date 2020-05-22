package services

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/marvinsenkiv/covid/model"
	"github.com/segmentio/ksuid"
)

// Instance a Covid model
func CreateCovidData(record []string) model.Covid {
	id := GenKsuid()
	province := record[0]
	country := record[1]
	confirmed, _ := strconv.Atoi(record[3])
	deaths, _ := strconv.Atoi(record[4])
	recovered, _ := strconv.Atoi(record[5])
	return model.Covid{id, province, "", country, confirmed, deaths, recovered}
}

// Instance a Covid model with second data model
func CreateDataFips(record []string) model.Covid {
	id := GenKsuid()
	city := record[1]
	province := record[2]
	country := record[3]
	confirmed, _ := strconv.Atoi(record[7])
	deaths, _ := strconv.Atoi(record[8])
	recovered, _ := strconv.Atoi(record[9])
	return model.Covid{id, province, city, country, confirmed, deaths, recovered}

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

// ReadFiles a lot of files from path
func ReadFiles() []string {
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

func GenKsuid() string {
	id := ksuid.New()
	return id.String()
}
