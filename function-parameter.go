package main

import "fmt"

func sayhello(firstname string, lastname string){

	fmt.Println("hello", firstname, lastname)

}

func main() {
	sayhello("eko", "nugraha")
}