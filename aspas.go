package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	texto := "teste \"aaaa\" sdfsd "
	ioutil.WriteFile("teste.txt", []byte(texto), 0644)
	t, multibyte, tail, err := strconv.UnquoteChar("a", 0)
	fmt.Println(t, multibyte, tail, err)
}
