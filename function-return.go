package main

import "fmt"

func kamu(name string) string {
	return "Hello, " + name

}

func main() {
	result := kamu("kontol")
	fmt.Println(result)
}