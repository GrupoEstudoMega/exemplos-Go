package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
)

func main() {
	db, err := sql.Open("oci8", "system/mega@LOCAL")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select 1,2,3,4,5,6,7,8, sysdate from dual")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		fmt.Println("got row")
	}
}
