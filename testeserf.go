package main

import (
	"fmt"
	"github.com/hashicorp/serf/client"
	"mega/go-util/erro"
)

func main() {
	config := client.Config{Addr: "127.0.0.1:7373"}
	cli, err := client.ClientFromConfig(&config)
	erro.Trata(err)
	fmt.Println(cli.Members())
}
