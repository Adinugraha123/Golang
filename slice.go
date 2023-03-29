package main 
import "fmt"

func main() {

	// tanda := adalah pengganti var 
	names := []string{"foo", "bar", "baz", 
	"bazq","kamu", "kamuq", "adi", "selasana" }
	slices := names[0:8]
    fmt.Println(slices[0])
	fmt.Println(slices[1])

	days := [...]string {"senin", "selasa", "rabu", "kamis",
"jumat", "sabtu", "minggu"}

daySlice1 := days[5:]
daySlice1[0] = "Sabtu Baru"
daySlice1[1] = "minggu Baru"

fmt.Println(days)

daySlice2 := append(daySlice1, "libur baru")
daySlice2[0] = "kontol"
fmt.Println(daySlice2)
fmt.Println(days)

newSlice := make([]string, 2 ,7)
newSlice[0] = "wkwk"
newSlice[1] = "KKKK"

fmt.Println(newSlice)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))

fromslice := days[:]
toslice := make([]string, len(fromslice), cap(fromslice))	

copy(toslice, fromslice)

fmt.Println(toslice)

iniarray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
inislice := []int{1, 2, 3, 4, 5,}

fmt.Println(iniarray)
fmt.Println(inislice)
}