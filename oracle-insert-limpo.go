package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	/*"sync"
	"time"*/)

func main() {
	db, err := sql.Open("oci8", "system/mega@127.0.0.1:1521/mega")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("antes do delete")

	//db.Exec("delete from test")

	fmt.Println("depois do delete")
	_, err = db.Exec("insert into test values (:pk, :id)", 8, 7)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("depois do insert")

	/*var wg sync.WaitGroup

	wg.Add(100)
	a := 1
	for x := 1; x <= 100; x++ {
		go func(x int) {
			a = x
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}(x)
	}
	wg.Wait()*/

}
