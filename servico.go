package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	)

func main() {
	resp, err := http.Get("http://example.com")
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(err)
}