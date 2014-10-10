package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type teste struct {
	Valor *int `json:"valor" xml:"valor,attr"`
}

func main() {

	testeXml := `<teste>`

	var a teste
	xml.Unmarshal([]byte(testeXml), &a)
	fmt.Println(a.Valor)
	//a.Valor = nil
	//*a.Valor = 1
	//fmt.Println(*a.Valor)
	js, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(js))
	var b teste
	json.Unmarshal(js, &b)
	fmt.Println(b.Valor)
}
