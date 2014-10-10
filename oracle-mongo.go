package main

import (
	"../comparador-banco/comparador"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"runtime"
	"sync"
	"time"
)

var session *mgo.Session
var mongo *mgo.Database

var db *sql.DB
var query string = "select mov_in_numlancto, m.mov_st_documento, mov_dt_datadocto, mov_Re_valor from mgfin.fin_movimento m"

func main() {
	runtime.GOMAXPROCS(3)
	os.Setenv("NLS_LANG", ".UTF8")
	var err error
	db, err = sql.Open("oci8", "system/mega@127.0.0.1:1521/mega")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//mongo = GetMongoDb()
	//os.Setenv("NLS_LANG", ".WE8PC850")
	fmt.Println("Só query")
	executaSoQuery()
	fmt.Println("Executando e transformando em array de map")
	executaConverteMap()
	fmt.Println("Executando e serializando JSON")
	executaSerializaJson()
	fmt.Println("Gravando no Mongo")
	executaGravaMongo()
}

func executaSoQuery() {

	defer comparador.Trace(time.Now())
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	total := 0
	for rows.Next() {
		total++
	}
	fmt.Println(total)

}

func executaConverteMap() *comparador.Resultado {

	//fmt.Println(getDSN())

	defer comparador.Trace(time.Now())

	linhas, colunas, err := comparador.ExecutaSql(db, query)

	fmt.Println(len(linhas))
	fmt.Println(linhas[0])
	fmt.Println(colunas)
	fmt.Println(err)

	return &linhas
}

func executaSerializaJson() {

	//fmt.Println(getDSN())

	defer comparador.Trace(time.Now())

	linhas, _, _ := comparador.ExecutaSql(db, query)
	linhasJson, _ := json.Marshal(linhas)
	fmt.Println(len(linhasJson))
}

type Movimento struct {
	Id        bson.ObjectId `bson:"_id"`
	Lancto    int
	Documento string
	Data      time.Time
	Valor     float32
}

func executaGravaMongo() {
	defer comparador.Trace(time.Now())
	session := GetMongoSession()
	collection := session.DB("local").C("Movimentos")
	collection.DropCollection()
	session.Close()

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}

	/*scanArgs := make([]interface{}, 2)

	values := make([]interface{}, 2)

	for i := range values {
		scanArgs[i] = &values[i]
	}*/

	var wg sync.WaitGroup
	wg.Add(264427)
	count := 0
	total := 0

	for rows.Next() {

		var mov Movimento
		mov.Id = bson.NewObjectId()
		rows.Scan(&mov.Lancto, &mov.Documento, &mov.Data, &mov.Valor)
		go func(mov Movimento) {
			session := GetMongoSession()
			defer session.Close()
			collection := session.DB("local").C("Movimentos")
			total++
			count++
			if count > 1000 {
				fmt.Println(total)
				count = 0
			}
			collection.Insert(mov)
			wg.Done()
		}(mov)
	}
	wg.Wait()

	//resultado := executaConverteMap()
	//collection.Insert(resultado)
	//collection.Insert(...)
	//fmt.Println(collection)
	//collection.

}

func GetMongoSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("mongodb://localhost/local")
		session.SetMode(mgo.Monotonic, true)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return session.Clone()
}
