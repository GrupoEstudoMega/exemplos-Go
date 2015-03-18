package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/a", teste)
	http.ListenAndServe(":8090", nil)

}

func teste(res http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("oracle.go")
	io.Copy(res, f)
}
