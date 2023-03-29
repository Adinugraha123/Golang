package main

import "fmt"

// defer
func logging() { 
		pesan := recover()
		if pesan != nil {
			
	fmt.Println("error" , pesan)
		}
	fmt.Println("selesai memanggil function")
}

// func running() {
// 	defer logging() // menghentikan namun masih dieksekusi
// 	fmt.Println("running")
// }

func runner(error bool) {
	defer logging()

	if error {
		panic("lah kok error")
	}

	fmt.Println("running bos")
}


func main() {
 	runner(false)
}


