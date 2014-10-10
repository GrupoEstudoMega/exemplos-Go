package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a, err := filepath.Rel(`C:\mega\megaempresarial`, "C:/mega/megaempresarial/Fontes")
	fmt.Println(a, err)
}
