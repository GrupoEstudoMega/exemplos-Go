package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/corydoras/go-oci8"
	_ "github.com/mattn/go-oci8"
	_ "github.com/tgulacsi/goracle/godrv"
	"runtime"
	"sync"
	"time"
)

func trace(startTime time.Time) {
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
}

func main() {
	var m sync.Mutex
	runtime.GOMAXPROCS(1)
	//db := connect()
	//defer db.Close()
	db := connect()
	defer db.Close()
	times := 1000
	var wg sync.WaitGroup
	wg.Add(times)
	for i := 1; i <= times; i++ {
		go func() {
			m.Lock()
			/*db := connect()
			defer db.Close()*/
			exec(db)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func connect() *sql.DB {
	db, err := sql.Open("goracle", "mgfin/megafin@127.0.0.1:1521/mega")
	//fmt.Println(db)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func exec(db *sql.DB) {
	defer trace(time.Now())
	//rows, err := db.Query(`select 1 from dual`)
	rows, err := db.Query(`select mov_in_numlancto from fin_movimento where mov_st_complhist like '%teste%'`)
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
