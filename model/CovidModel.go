package model

import (
	"database/sql"

	// i dont know
	_ "github.com/go-sql-driver/mysql"
)

//Covid Model
type Covid struct {
	ID        string
	Province  string
	City      string
	Country   string
	Confirmed int
	Deaths    int
	Recovered int
}

// Insert covid data
func (c Covid) Insert(query string) bool {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Covid")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	return true
}

func (c Covid) Select(query string) string {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Covid")
	var id string
	var res string

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	sel, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for sel.Next() {
		err = sel.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		res = id
	}

	return res

}
