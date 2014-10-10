package main

import "reflect"

func main() {
	var a = make(map[string]string)
	var b = make(map[string]string)

	a["teste1"] = "teste1"
	a["teste2"] = "teste2"

	b["teste2"] = "teste2"
	b["teste1"] = "teste1"

	ar := []map[string]string{a}
	ar2 := []map[string]string{b}

	println(reflect.DeepEqual(ar, ar2))
}
