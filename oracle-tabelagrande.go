package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	// "github.com/davecheney/profile"
	_ "github.com/mattn/go-oci8"
	_ "github.com/tgulacsi/goracle/godrv"
	"os"
	"time"
	// "strings"
	// "reflect"
	// "time"
	//	"strconv"
)

type Valor struct {
	Valor string
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	os.Setenv("NLS_LANG", ".UTF8")
	//os.Setenv("NLS_LANG", ".WE8PC850")
	executa()
}

func executa() {

	defer Trace(time.Now())
	//fmt.Println(getDSN())
	db, err := sql.Open("oci8", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//rows, err := db.Query("select 10 valor, sysdate - 1 data from mgglo.glo_pais where pa_st_sigla = 'SUI'")
	rows, err := db.Query("select * from mgmag.efd_registro_c100-- where rownum <= 20")
	//fmt.Println(rows, err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//fmt.Println(PackageData(rows))

	//valuesString := make([]string, 2)
	x := 0
	for rows.Next() {
		x++
		//fmt.Println("leu fora")
		//fmt.Println(x)
		/*		var i sql.NullInt64
				var campo string
				var tag string
				var tabela sql.NullString
				var s sql.NullString*/
		//rows.Scan(&xml)
		//ioutil.WriteFile(`c:\temp\sped_gerado.txt`, []byte(xml), 0644)
		//fmt.Println(values[1])
		//rows.Scan(&s)
		/*fmt.Println(i)
		fmt.Println(campo)
		fmt.Println(tag)
		fmt.Println(tabela)
		fmt.Println(strings.Replace(s.String, string(0), "", -1))*/
		/*var _ time.Time
		//data, ok := scanArgs[1].(time.Time)
		//data := reflect.ValueOf(scanArgs[0]).(string)
		fmt.Println(values[0])
		numero, ok := values[0].(int)
		fmt.Println(numero, false)
		data, ok := values[1].(time.Time)
		fmt.Println(data)
		texto, ok := scanArgs[0].(string)
		jsonData, err := json.MarshalIndent(map[string]interface{}{"valor": values[0]}, "", "  ")
		fmt.Printf("%s", string(jsonData))
		jsonArq, err := json.MarshalIndent(scanArgs, "", "  ")
		ioutil.WriteFile("c:/teste.json", jsonArq, 0644)
		fmt.Println(string(jsonArq), err)

		fmt.Println(texto, ok)*/

		//convertAssign(&retorno, scanArgs[1])
		//fmt.Println([]byte(retorno))
	}
	fmt.Println(x)
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

func Trace(startTime time.Time) {
	endTime := time.Now()
	fmt.Printf("%.2fs\n", endTime.Sub(startTime).Seconds())
}
