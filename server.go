package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/teste", teste)
	http.ListenAndServe(":8090", nil)

}

func teste(res http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	res.Write([]byte("teste"))
	fmt.Println("respondeu")
}
