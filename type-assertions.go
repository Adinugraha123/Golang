package main 

import "fmt"

func random()interface{}{
	
	return "eko"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)


}
