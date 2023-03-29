package main
import "fmt"

func faktorial(value int) int {
	result := 1 
	for i := value; i > 0; i-- {
		result *= i 
	}
	return result
}

func faktorialrecursive(value int) int {
	if value ==	1 {
		return 1 // kondisii harus berhentinya
	} else {
		return value * faktorialrecursive(value -1)
	}
}

func main() {
	loop :=faktorial(10)
	fmt.Println(loop)
	//fmt.Println(5 * 4 * 3 * 2 * 1)
	
	recursive := faktorialrecursive(10)
	fmt.Println(recursive)

}

// recursive function  memanggil diri sendiri

