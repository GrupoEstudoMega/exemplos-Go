package main

import (
	"mega/go-util/dbg"
	"time"
)

func main() {
	defer dbg.Trace(time.Now())
	a := 0
	for x := 0; x < 1000000000; x++ {
		a++
	}
}
