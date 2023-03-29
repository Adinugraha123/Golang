package main 
import "fmt"

func main() {
	name := "adi"
	counter := 0

	hay := func(){
		name := "jaenal"
		// variabel ini menghapus variabel di atasnya
		fmt.Println("Starting")
		counter++
		fmt.Println(name) // bikin variabel yang beda

	}
	hay()
	hay()

	fmt.Println(counter)
	fmt.Println(name)
}