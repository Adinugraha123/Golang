package main 

import "fmt"


 type Customer struct {
 	nama, alamat string
 	age int

}

// func main() {
// 	var adi customer

// 	adi.nama = "adi"
// 	adi.alamat = "alamat"
// 	adi.age = 21

// 	fmt.Println(adi)

// }

// func main() {

// joko := customer{
// 	nama: "joko",
// 	alamat: "surabaya",
// 	age: 21,
// }

// fmt.Println(joko)

// budi := customer{"budi", "bandung", 40}

// fmt.Println(budi)

// }



func (pelanggan Customer) sayhello(){
	fmt.Println("hello", pelanggan.nama)
}

func main() {
	romlah := Customer{nama: "rumlah"}
	romlah.sayhello()
}