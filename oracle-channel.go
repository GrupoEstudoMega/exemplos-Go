package main

import (
	"database/sql"
	"fmt"
	//"github.com/eapache/channels"
	_ "github.com/mattn/go-oci8"
	"mega/go-util/erro"
	"time"
)

func main() {
	//tasks := make(chan int, 100)

	query := "select * from mgmag.efd_registro_c170"

	db, err := sql.Open("oci8", "system/mega@local")
	erro.Trata(err)
	defer db.Close()

	//defer Trace(time.Now())

	rows, err := db.Query(query)
	erro.Trata(err)
	defer rows.Close()

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	total := 0

	for rows.Next() {

		total++

		rows.Scan(scanArgs...)

		//linhas := channels.NewInfiniteChannel()

		linhas := make(chan []interface{}, 1000)

		/*if i == 5000 {
			tasks.Resize(500000)
		}*/
		//fmt.Println(tasks.Cap())
		//select {
		linhas <- values
		/*	fmt.Println("mandou", i)
			default:
				fmt.Println("miou", i)

			}*/
		//linhas.Close()
		close(linhas)
	}
	fmt.Println(total)

	time.Sleep(time.Second * 5)
}
