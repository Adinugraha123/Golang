package main
import "fmt"

type apapun interface{

}

func ups(i int)  interface{}{
	

	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = ups(0)
	fmt.Println(data)
}