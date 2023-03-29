package main
import "fmt"

func main() {
type NoKTP string
var ktpadi NoKTP = "1111111111111111"
fmt.Println(ktpadi)
fmt.Println(NoKTP("222222222222222"))
}