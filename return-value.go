package main

import "fmt"

func kamu()(string, string){
	return "aku", "iyen"
}

func main() {
	firstname, _ := kamu() // untuk menghilangkan lastname itu harus menggunakan _ (garis bawah)
	fmt.Println(firstname)
}