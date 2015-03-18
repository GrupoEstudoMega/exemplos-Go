package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", `c:\temp\teste.db`)

	rows, err := db.Query("select * from Teste")

	fmt.Println(err)

	var nome string
	var idade int

	for rows.Next() {
		rows.Scan(&nome, &idade)
		fmt.Println(nome, idade)
	}

}
