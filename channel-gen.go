package main

import (
	"fmt"
	"net/http"
	"sync"
	//	"time"
)

type Proc func(param interface{})

func Teste(param interface{}) {
	url, funf := param.(string)
	if funf {
		resp, err := http.Get(url)
		fmt.Println("--------------------------------------------------")
		fmt.Println(url, err, resp)

	}
}

func main() {
	tasks := make(chan interface{}, 100)

	//var call Proc

	call := Teste

	// spawn four worker goroutines
	var wg sync.WaitGroup
	/*for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("caiu")
			for param := range tasks {
				//time.Sleep(1 * time.Second)
				call(param)
			}
			fmt.Println("terminou for")
			wg.Done()
		}()
	}*/

	//time.Sleep(5 * time.Second)

	// generate some tasks
	/*for i := 0; i < 20; i++ {

		fmt.Println("mandou")
		tasks <- i
	}*/

	go func() {
		for x := 0; x < 100; x++ {
			tasks <- "http://www.google.com"
			wg.Add(1)
		}
	}()

	for param := range tasks {

		fmt.Println("passou")
		go func(param interface{}) {
			call(param)
			wg.Done()
		}(param)
	}

	/*
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"
		tasks <- "http://www.google.com"
		tasks <- "http://www.microsoft.com"
		tasks <- "http://www.uol.com.br"
		tasks <- "http://www.golang.org"*/
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
