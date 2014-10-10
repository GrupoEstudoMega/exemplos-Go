package main

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb"
	"os"
	"os/exec"
	"time"
)

func main() {
	limpa()
	fmt.Print("\rteste\n")
	/*	fmt.Print("\rteste\n")*/
	count := 10
	bar := pb.StartNew(count)
	bar.SetRefreshRate(time.Millisecond * 10)
	w := bufio.NewWriter(os.Stdout)

	for i := 0; i < count; i++ {
		fmt.Fprint(w, "Hello, ")
		fmt.Fprint(w, "world!")

		bar.Increment()

		//fmt.Print("\rteste\n")
		time.Sleep(time.Second)
	}
	bar.FinishPrint("The End!")
	w.Flush() // Don't forget to flush!
}

func limpa() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
