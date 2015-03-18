package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"io/ioutil"
	"os"
	// "strings"
	// "reflect"
	// "time"
	//	"strconv"
)

type Valor struct {
	Valor string
}

func main() {
	os.Setenv("NLS_LANG", ".UTF8")
	//os.Setenv("NLS_LANG", ".WE8PC850")
	executa()
}

func executa() {
	//fmt.Println(getDSN())
	db, err := sql.Open("oci8", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//rows, err := db.Query("select 10 valor, sysdate - 1 data from mgglo.glo_pais where pa_st_sigla = 'SUI'")
	//rows, err := db.Query("select xml_bl_xml from mgint.int_xml where trn_in_id = 3626")
	//rows, err := db.Query("select xml_bl_xml from mgint.int_xml where trn_in_id = 3626")
	rows, err := db.Query("select value from clobtest where id = 3")
	//fmt.Println(rows, err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//fmt.Println(PackageData(rows))

	var xml sql.NullString

	//valuesString := make([]string, 2)

	for rows.Next() {
		fmt.Println("teste")

		/*		var i sql.NullInt64
				var campo string
				var tag string
				var tabela sql.NullString
				var s sql.NullString*/
		rows.Scan(&xml)
		fmt.Println(xml.String)
		ioutil.WriteFile(`c:\temp\teste_clob.xml`, []byte(xml.String), 0644)
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
