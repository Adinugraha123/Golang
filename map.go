package main
import "fmt"

func main() {

	person := map[string]string{
		"person" : "John",
		"person2" : "hasan",
	}
	fmt.Println(person)
	fmt.Println(person["person"])
	fmt.Println(person["person2"])

	book := make(map[string]string)
	book["judul"] = "buku bokep"
	book["penulis"] = "hitomi"

	delete(book, "penulis")

	fmt.Println(book)
	


	}
