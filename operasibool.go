package main

import (
	"fmt"

)

func main() {

	var abi = 10
	var biu = 89
	var rara = 70

	var a bool = abi > 100
	var b bool = biu > 100
	var e bool = rara > 100

	var c bool = a && b 
	var d bool = b != e
	fmt.Println(c)
	fmt.Println(d)
}