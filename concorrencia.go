package main

import (    
    "fmt"
    "math/rand"
    "runtime"    
)

func process(c chan int) {
  y := rand.Intn(100)
  for x := 0; x < 100000000; x++ {
    y += x;
  }
  c <- y;
  
}

func main() {
  runtime.GOMAXPROCS(8)
  c := make(chan int, 100)
  for x := 0; x < 10000; x++ {    
    go process(c)
  }
  for x := 0; x < 10000; x++ {    
    fmt.Println( <- c)
  }

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}