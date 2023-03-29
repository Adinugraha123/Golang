package main
import "fmt"

type namaseorang interface {
getname() string
}

func sayhi(Nama namaseorang) {
fmt.Println("hello", Nama.getname())
}

// func main(){
// 	var adi namaseorang
// 	sayhi()
// }

type person struct {
	Nama string
}

func (p person) getname() string {
	return p.Nama
}
func main() {
	var adi person
	adi.Nama = "ekontol"
sayhi(adi)

}