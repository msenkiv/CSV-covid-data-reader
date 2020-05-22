package services

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/marvinsenkiv/covid/model"
)

// InsertCovidData in db
func InsertCovidData(c model.Covid, file string) string {
	insertCountry(c, file)
	return ""

}

func insertConfirmed(c model.Covid) {
	query := fmt.Sprintf("INSERT INTO Covid.numConfirmados VALUES('%s', %d)", c.ID, c.Confirmed)
	c.Insert(query)

}

func insertNumDeaths(c model.Covid) {
	query := fmt.Sprintf("INSERT INTO Covid.numMortes VALUES('%s', %d)", c.ID, c.Deaths)
	c.Insert(query)

}

func insertNumRecovered(c model.Covid) {
	query := fmt.Sprintf("INSERT INTO Covid.numRecuperados  VALUES('%s', %d)", c.ID, c.Recovered)
	c.Insert(query)
}

func insertCountry(c model.Covid, file string) {
	// create regex
	valid := regexp.MustCompile("^[A-Za-z0-9 ]+$")

	s := strings.TrimSpace(c.Country)
	e := strings.TrimSpace(c.City)

	if !valid.MatchString(s) {
		s = "' '"
	}
	if !valid.MatchString(e) {
		e = "' '"
	}

	if c.Country == "Cote d'Ivoire" {
		fmt.Println("gambi")
		c.Country = "Cote dIvoire"
	}

	id := verifyCoutryDB(c)

	if id == "" {
		id = GenKsuid()
		query := fmt.Sprintf("INSERT INTO Covid.Country values ('%s', '%s')", id, c.Country)
		c.Insert(query)

	}
	idP := verifyProvinceDB(c)

	if idP == "" {
		idP = GenKsuid()
		query := fmt.Sprintf("INSERT INTO Covid.Province values ('%s', '%s')", idP, c.Province)
		c.Insert(query)

	}

	insertCovidData(c, id, idP, file)

}
func insertCovidData(c model.Covid, id string, idP string, file string) {
	query := fmt.Sprintf("INSERT INTO Covid.CovidData VALUES('%s','%s','%s',%d,%d,%d,'%s')", c.ID, id, idP, c.Deaths, c.Confirmed, c.Recovered, file)
	c.Insert(query)
}

func verifyCoutryDB(c model.Covid) string {
	query := fmt.Sprintf("SELECT idCountry FROM Covid.Country WHERE nameCountry ='%s'", c.Country)
	return c.Select(query)

}
func verifyProvinceDB(c model.Covid) string {
	query := fmt.Sprintf("SELECT idProvince from Covid.Province where nameProvince='%s';", c.Province)
	return c.Select(query)
}
