package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
)

func main() {
	db, err := sql.Open("oci8", "system/mega@127.0.0.1:1521/mega")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//for x := 1; x <= 2; x++ {
	//rows, err := db.Query(fmt.Sprintf("select id, value from clobtest where id = %d", x))
	rows, err := db.Query("select trd_in_id, trd_bl_scriptregistro from mgint.int_traducao where pro_in_id = 8090 and trd_ch_tipotag = 'A'")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var id int
	var value sql.NullString

	for rows.Next() {
		rows.Scan(&id, &value)
		if value.Valid {
			fmt.Println("------------------------")
			fmt.Printf("Row: %d, length: %d \n", id, len(value.String))
			fmt.Println(value.String)
		}
	}
	//}
}
