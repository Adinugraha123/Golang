package main

import "fmt"

func main(){

var nilai32 int32 = 327678
var nilai64 int64 = int64(nilai32)
var nilai16 int16 = int16(nilai32)

fmt.Println(nilai32)
fmt.Println(nilai64)
fmt.Println(nilai16)

var name = "budi"
var b = name[0]

// bisa diubah eStirng atau b d c dan lain-lain 

var bString =string(b)
fmt.Println(name)
fmt.Println(bString)
}