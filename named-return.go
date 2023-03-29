package main 
import "fmt"


func getcomplete()(firstname, middlename, lastname string ) {
firstname = "eko"
middlename = "kamu"
lastname = "nugraha"

return firstname, middlename, lastname
}

func main() {
	firstname, middlename, lastname := getcomplete()

	fmt.Println(firstname, middlename, lastname)


}