package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"mega/go-util/dbg"
	"os"
	"sync"
	"time"
)

func main() {
	//os.Setenv("NLS_LANG", ".UTF8")
	executa()
}

func executa() {
	db, err := sql.Open("oci8", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	times := 3

	selects := []string{"select uf_st_sigla, log_st_nome from mgglo.glo_logradouro",
		"select acaom_st_documento, lan_ch_tipocontab from mgglo.glo_acaomovimento",
		"select ref_st_tipo, org_tau_st_codigo from mgfin.fin_referenciafin"}
	queries := make([]*sql.Rows, 0, times)

	/*for x := 0; x < times; x++ {

		rows, err := db.Query(selects[x])
		if err != nil {
			fmt.Println(err)
		}
		rows.Next()
		queries = append(queries, rows)
	}*/

	defer dbg.Trace(time.Now())

	var wg sync.WaitGroup
	wg.Add(times)

	//var m sync.Mutex

	total := 0

	for x := 0; x < times; x++ {
		rows, err := db.Query(selects[x])
		if err != nil {
			fmt.Println(err)
		}
		queries = append(queries, rows)
		for {

			//m.Lock()
			if queries[x].Next() {
				var f1 string
				var f2 string
				queries[x].Scan(&f1, &f2)
				total++
				//fmt.Println(f1)
				if f1 == f2 {
				}
				//println(f1, f2)
				//m.Unlock()
			} else {
				//m.Unlock()
				break
			}

		}
		rows.Close()
		wg.Done()

		/*go func(rows *sql.Rows, wg *sync.WaitGroup) {

			for {
				//m.Lock()
				if rows.Next() {
					var f1 string
					var f2 string
					rows.Scan(&f1, &f2)
					//fmt.Println(f1)
					total++
					if f1 == f2 {
					}
					//println(f1, f2)
					//m.Unlock()
				} else {
					//m.Unlock()
					break
				}

			}
			wg.Done()
		}(queries[x], &wg)*/
	}
	wg.Wait()
	fmt.Println(total)

	/*if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var f1 string
		var f2 string
		rows.Scan(&f1, &f2)
		println(f1, f2)
	}*/
}

func getDSN() string {
	var dsn string
	if len(os.Args) > 1 {
		dsn = os.Args[1]
		if dsn != "" {
			return dsn
		}
	}
	dsn = os.Getenv("GO_OCI8_CONNECT_STRING")
	if dsn != "" {
		return dsn
	}
	fmt.Fprintln(os.Stderr, `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
or as the first argument! (The format is user/name@host:port/sid)`)
	os.Exit(1)
	return ""
}
