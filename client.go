package main

import (
	"fmt"
	"io/ioutil"
	"mega/go-util/erro"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for x := 0; x < 3950; x++ {
		wg.Add(1)
		go func() {
			request()
			wg.Done()
		}()
	}
	wg.Wait()

}

func request() {
	resp, err := http.Get("http://localhost:8090/teste")
	erro.Trata(err)
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	erro.Trata(err)
	fmt.Println(string(body))
}
