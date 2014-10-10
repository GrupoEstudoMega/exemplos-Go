package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	_ "github.com/mattn/go-oci8"
	// "io/ioutil"
	"os"
	"time"
)

type Registro struct {
	Codigo int
	Pai    int
	Query  string
	Rows   *sql.Rows
	Filhos []*Registro
	Nivel  int
}

var db *sql.DB

func main() {
	os.Setenv("NLS_LANG", ".UTF8")
	//os.Setenv("NLS_LANG", ".WE8PC850")
	executa()
}

func executa() {

	var err error
	db, err = sql.Open("oci8", "mega/megamega@local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select to_date('19/10/2013', 'dd/mm/yyyy') from dual")
	rows.Next()
	var data time.Time
	rows.Scan(&data)
	fmt.Println(data)
	fmt.Println(data.Add(6 * time.Hour).Format("01/02/2006"))
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
