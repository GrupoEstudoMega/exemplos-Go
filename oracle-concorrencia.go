package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/corydoras/go-oci8"
	_ "github.com/mattn/go-oci8"
	//_ "github.com/tgulacsi/goracle/godrv"
	"runtime"
	"sync"
	"time"
)

func trace(startTime time.Time) {
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
}

func main() {
	//var m sync.Mutex
	runtime.GOMAXPROCS(1)
	//db := connect()
	//defer db.Close()
	db := connect()
	defer db.Close()
	times := 100
	var wg sync.WaitGroup
	wg.Add(times)

	a := "teste"
	for i := 1; i <= times; i++ {
		go func(i, num1 int, num2 int) {
			//m.Lock()
			/*db := connect()
			defer db.Close()*/
			fmt.Println("disparou", i)
			exec(db, num1, num2)
			fmt.Println(a)
			//m.Unlock()
			wg.Done()
		}(i, (i-1)*1000000, i*1000000)
	}
	wg.Wait()
}

func connect() *sql.DB {
	db, err := sql.Open("oci8", "mega/megamega@pc_leandroa:1521/wine")
	//fmt.Println(db)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func exec(db *sql.DB, num1 int, num2 int) {
	defer trace(time.Now())
	//rows, err := db.Query(`select 1 from dual`)
	//rows, err := db.Query(`select mov_in_numlancto from mgfin.fin_movimento where mov_st_complhist like '%teste%'`)
	rows, err := db.Query(`select count(1) from mgfin.fin_movimento where mov_in_numlancto between :num1 and :num2`, num1, num2)
	fmt.Println(num1, num2)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var num int
		rows.Scan(&num)
		println(num)
	}
}
