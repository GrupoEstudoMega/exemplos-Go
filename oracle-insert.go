package main

import (
	"database/sql"
	// "encoding/json"
	// "github.com/davecheney/profile"
	"flag"
	"fmt"
	_ "github.com/mattn/go-oci8"
	_ "github.com/mattn/go-sqlite3"
	"mega/go-util/erro"
	//"sync"
)

func main() {
	connect := flag.String("connect", "mega/megamega@127.0.0.1:1521/mega", "Connect string no format usuaro/mega@maquina:porta/instancia ou nome do TNS_NAMES")
	//query := flag.String("query", "select 1 from dual", "Query a ser executada")
	flag.Parse()
	fmt.Println(*connect)
	db, err := sql.Open("oci8", *connect)
	//db, err := sql.Open("sqlite3", `c:\temp\teste.db`)
	erro.Trata(err)

	db.Exec("delete from teste")
	//for x := 1; x <= 100; x++ {
	//go func(x int) {
	x := 1
	_, err = db.Exec("insert into teste values (:pk, :id)", x, x)
	erro.Trata(err)
	//	wg.Done()
	//}(x)
	//}
	//wg.Wait()

}
