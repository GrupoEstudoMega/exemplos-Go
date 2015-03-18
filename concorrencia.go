package main

import (
	"fmt"
	//"math/rand"
	"runtime"
	"time"
)

func process(c chan int) {
	time.Sleep(1 * time.Second)
	/*y := rand.Intn(100)
	for x := 0; x < 1000000000; x++ {
		y += x
	}*/
	c <- 1

}

func main() {
	runtime.GOMAXPROCS(8)
	c := make(chan int, 2)
	for x := 0; x < 100; x++ {
		go process(c)
	}
	for x := 0; x < 100; x++ {
		fmt.Println(<-c)
	}

	//var input string
	//fmt.Scanln(&input)
	fmt.Println("done")
}
