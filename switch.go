package main
import "fmt"

func main() {

	nama := "eko"

	switch nama {
		case "eko":
			fmt.Printf("hello eko")
		case "joko":
			fmt.Printf("hello joko")
		default: 
		fmt.Printf("kamu siapa kontol")
	}

	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("nama panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

}