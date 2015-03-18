package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	senha := os.Getenv("SENHA")
	fmt.Println(senha)
	os.Unsetenv("SENHA")
	time.Sleep(100 * time.Second)
}
