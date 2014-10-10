package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"os"
	"sync"
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

	times := 100

	queries := make([]*sql.Rows, 0, times)

	for x := 0; x < times; x++ {
		rows, err := db.Query("select pa_st_sigla, pa_st_nome from mgglo.glo_pais")
		if err != nil {
			fmt.Println(err)
		}
		rows.Next()
		queries = append(queries, rows)
	}

	var wg sync.WaitGroup
	wg.Add(times)

	var m sync.Mutex

	for x := 0; x < times; x++ {
		go func(rows *sql.Rows, wg *sync.WaitGroup, m *sync.Mutex) {

			for {
				//m.Lock()
				if rows.Next() {
					var f1 string
					var f2 string
					rows.Scan(&f1, &f2)
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
		}(queries[x], &wg, &m)
	}
	wg.Wait()

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
