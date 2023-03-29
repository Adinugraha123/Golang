package main 
import "fmt"

type Kontol func(string)bool
// diatas merupakan type declaration

func daftaruser (name string, kontol Kontol){
	if kontol(name) {
		fmt.Println("kontol kamu panjang", name)
		
	} else {
		fmt.Println("kontol kamu enak", name)
	}
}

func main() { 
	kontol := func (name string) bool  {

		return name == "enak"
		
	}

	daftaruser("punya oci", kontol)
	daftaruser("enak", func (name string) bool {
		return name == "enak"
	})
}
