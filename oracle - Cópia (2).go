package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	_ "github.com/mattn/go-oci8"
	// "io/ioutil"
	"os"
	// "time"
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
	db, err := sql.Open("oci8", "system/mega@local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//rows, err := db.Query("select 10 valor, sysdate - 1 data from mgglo.glo_pais where pa_st_sigla = 'SUI'")
	//rows, err := db.Query("select REG,IND_OPER,IND_EMIT,COD_PART,COD_MOD,SER,SUB, NUM_DOC, to_char(c.DT_DOC) from MGMAG.EFD_REGISTRO_C113 c where prm_in_id = 132  and rownum = 1 order by reg_in_sequencial")
	//rows, err := db.Query("select REG, IND_OPER,IND_EMIT,COD_PART,COD_MOD,1,1, 1, DT_DOC from MGMAG.EFD_REGISTRO_C113 c where prm_in_id = 132  and rownum = 1 order by reg_in_sequencial")
	rowsAgente, err := db.Query("select agn_in_codigo from mgglo.glo_agentes where agn_pad_in_codigo = 1 and rownum < 10 order by agn_in_codigo")
	rowsTipo, err := db.Query("select agn_in_codigo, agn_tau_st_codigo from mgglo.glo_agentes_id where agn_pad_in_codigo = 1 order by agn_in_codigo")
	//fmt.Println(rows, err)
	if err != nil {
		fmt.Println(err)
	}
	defer rowsAgente.Close()
	defer rowsTipo.Close()

	var cod int
	var codfilho int
	var tipo string
	rowsTipo.Next()
	rowsTipo.Scan(&codfilho, &tipo)
	for rowsAgente.Next() {
		rowsAgente.Scan(&cod)
		fmt.Println(cod)
		if tipo != "" {
			fmt.Println(codfilho, tipo)
		}
		tipo = ""
		for rowsTipo.Next() {
			rowsTipo.Scan(&codfilho, &tipo)
			if codfilho != cod {
				break
			}
			fmt.Println(codfilho, tipo)

		}
		//fmt.Println("teste")
	}

	//fmt.Println(PackageData(rows))

	/*columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}*/

	//valuesString := make([]string, 2)

	//for rows.Next() {
	/*var f1 string
	var f2 string*/
	/*rows.Scan(&valores[0], &valores[1])
	println(valores[1])*/
	/*var f1 string
	var f2 interface{}
	//var retorno *string
	rows.Scan(&f1, &f2)*/
	//rows.Scan(scanArgs...)
	//rows.Scan(&f1, &f2)
	///=println(f1, f2)
	//println(string(varlues[1]))
	//println(string(values[1]))
	//fmt.Println([]byte("SUÍÇA"))
	//fmt.Println(f1, f2)
	//s := make([]byte, 4000)
	//var d interface{}
	/*var s string
	//var i int64
	rows.Scan(&s)
	//rv := reflect.ValueOf(s)
	i, _ := strconv.ParseInt(s, 10, 32)
	//fmt.Println(d)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s).Name())
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i).Name())
	/*switch s.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Print("int")
	case float32:
		fmt.Println("float32")
	}*/
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
	//}
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
