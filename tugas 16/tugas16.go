package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type pelajar struct {
	Id    string
	Nama  string
	Umur  int
	Nilai int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/db_belajar_sql")
	if err != nil {
		return nil, err
	}

	return db, nil
}
func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Umur = 20
	rows, err := db.Query("select Id, Nama, Umur , Nilai from tbl_pelajar where Umur = ?", Umur)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pelajar

	for rows.Next() {
		var each = pelajar{}
		var err = rows.Scan(&each.Id, &each.Nama, &each.Umur, &each.Nilai)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Nama)
	}
}
func main() {
	sqlQuery()
}
