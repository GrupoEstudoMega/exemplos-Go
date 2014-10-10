package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	hash := md5.New()
	io.WriteString(hash, "teste testes fadsjf ioasjfoi ajsoifj saoij oi")
	fmt.Printf("%x", hash.Sum(nil))
}
