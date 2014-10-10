package main

import (
	// "bytes"
	// "fmt"
	"bufio"
	// "io"
	"os"
)

func main() {
	arquivo, _ := os.Create(`c:\bufio.txt`)
	buf := bufio.NewWriter(arquivo)
	//buf := bytes.NewBuffer(arquivo)
	for x := 0; x < 1000; x++ {
		buf.WriteString("teste\n")
	}
	buf.Flush()
	//buf.WriteString("teste\n")
	//io.Copy(os.Stdout, buf)
}
