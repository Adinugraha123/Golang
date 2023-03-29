package main 
import "fmt"

func main() {
	var name [3] string
	name [0] = "eko"
	name [1] = "nugraha"
	name [2] = "paleo"

	fmt.Println(name[0], name[1], name[2])

	var values = [3] int{
		90,
		80,
		95,
	}
	fmt.Println(values)

	var nama = [3] string{
		"A",
		"B",
		"C",

	}

	fmt.Println(nama)
	var kamu = [...] string {
		"aku", 
		"cinta",
		"kamu",

	}
	fmt.Println(kamu)
	fmt.Println(len(kamu))
	kamu[0] = "kontol"
	fmt.Println(kamu)
}