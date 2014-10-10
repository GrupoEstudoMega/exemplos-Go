package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/a", teste)
	http.ListenAndServe(":8090", nil)

}

func teste(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("teste"))
}
