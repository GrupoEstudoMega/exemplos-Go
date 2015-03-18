package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := make(chan int, 0)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("caiu")
			for cmd := range tasks {
				//time.Sleep(1 * time.Second)
				fmt.Println(cmd)
			}
			fmt.Println("terminou for")
			wg.Done()
		}()
	}

	time.Sleep(5 * time.Second)

	// generate some tasks
	for i := 0; i < 20; i++ {

		fmt.Println("mandou")
		tasks <- i
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
