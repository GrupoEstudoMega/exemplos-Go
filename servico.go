package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(err)
}
