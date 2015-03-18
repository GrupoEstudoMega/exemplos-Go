package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	// "github.com/davecheney/profile"
	"flag"
	_ "github.com/mattn/go-oci8"
	//_ "github.com/mattn/go-sqlite3"
	"strings"
	"time"
	// "reflect"
	// "time"
	//	"strconv"
	"mega/go-util/dbg"
	"mega/go-util/erro"
)

type Valor struct {
	Valor string
}

func main() {
	connect := flag.String("connect", "mega/megamega@127.0.0.1:1521/mega", "Connect string no format usuaro/mega@maquina:porta/instancia ou nome do TNS_NAMES")
	query := flag.String("query", "select 1 from dual", "Query a ser executada")
	flag.Parse()
	//defer profile.Start(profile.CPUProfile).Stop()
	//os.Setenv("NLS_LANG", ".UTF8")
	//os.Setenv("NLS_LANG", ".WE8PC850")
	executa(*connect, *query)
}

func executa(connect string, query string) {
	fmt.Println(connect, query)

	dbg.SetDebug(false)

	//var err error
	//var db *sql.DB
	var banco string
	if strings.Contains(connect, ".db") {
		banco = "sqlite3"
	} else {
		banco = "oci8"
	}
	db, err := sql.Open(banco, connect)

	//db, err := sql.Open("oci8", connect)
	//db, err = sql.Open("sqlite3", `c:\temp\teste.db`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	defer Trace(time.Now())

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	columns, _ := rows.Columns()

	fmt.Println(columns)

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		erro.Trata(err)
		fmt.Println(values)
	}

}

func Trace(startTime time.Time) {
	endTime := time.Now()
	fmt.Printf("%.2fs\n", endTime.Sub(startTime).Seconds())
}
