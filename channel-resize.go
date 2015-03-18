package main

import (
	"fmt"
	"github.com/eapache/channels"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	//tasks := make(chan int, 100)

	tasks := channels.NewInfiniteChannel()

	//tasks.Resize(10000)

	//_ = make(chan int, 500000)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {

			time.Sleep(time.Second * 2)
			fmt.Println("caiu")
			debug.FreeOSMemory()

			for cmd := range tasks.Out() {
				//time.Sleep(time.Millisecond)
				inter := cmd.([]interface{})
				if inter[0] == 2 {
					fmt.Println(cmd)
				}
				//debug.FreeOSMemory()
			}
			wg.Done()
		}()
	}

	//time.Sleep(5 * time.Second)

	// generate some tasks
	/*chans := make(map[int]*channels.InfiniteChannel)
	time.Sleep(5 * time.Second)
	for i := 0; i < 10; i++ {
		chans[i] = channels.NewInfiniteChannel()
		//chans[i].Resize(100)
	}
	time.Sleep(5 * time.Second)*/
	/*for i := 0; i < 100000; i++ {
		chans[i] = nil
		//chans[i].Resize(100)
	}*/

	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 10000000; i++ {

		/*if i == 5000 {
			tasks.Resize(500000)
		}*/
		//fmt.Println(tasks.Cap())
		//select {
		tasks.In() <- []interface{}{1, 1, "teste", "outro teste", 1.56, true}
		/*	fmt.Println("mandou", i)
			default:
				fmt.Println("miou", i)

			}*/
	}
	tasks.Close()

	// wait for the workers to finish
	wg.Wait()
}
