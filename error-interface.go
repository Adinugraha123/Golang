package main 

import ( "fmt"
 "errors"
)

func Pembagi(nilai int, pembagi int )(int, error) {
	if  pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {	
	hasil, err := Pembagi(100, 20)
	if  err == nil {
		fmt.Println("hasil: ", hasil)
	} else {
		fmt.Println("err: ", err.Error())
	}
}