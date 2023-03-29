package main

import "fmt"

func sayhellofilter(name string, saringan func(string) string) {
	fmt.Println("hello", saringan(name))

}
// filter disini itu sebagai function


func spamfilter(name string) string{
	if name == "anjing"{
		return "..."
	}else {
		return name
	}
}

func main() {
	sayhellofilter("eko", spamfilter)

	saringan := spamfilter
	sayhellofilter("anjing", saringan)
}