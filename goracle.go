package main

import (
	"database/sql"
	"fmt"
	_ "github.com/tgulacsi/goracle/godrv"
	"os"
)

func main() {
	os.Setenv("NLS_LANG", ".WE8PC850")
	executa()
}

func executa() {
	db, err := sql.Open("goracle", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select pa_st_sigla, pa_st_nome from mgglo.glo_pais")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var f1 string
		var f2 string
		rows.Scan(&f1, &f2)
		println(f1, f2)
	}
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
