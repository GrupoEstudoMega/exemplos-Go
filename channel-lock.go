package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := make(chan int, 100)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("caiu")
			for cmd := range tasks {
				time.Sleep(200 * time.Millisecond)
				fmt.Println(cmd)
			}
			wg.Done()
		}()
	}

	//time.Sleep(5 * time.Second)

	// generate some tasks
	for i := 0; i < 10000; i++ {

		fmt.Println("mandou", i)
		tasks <- i
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
