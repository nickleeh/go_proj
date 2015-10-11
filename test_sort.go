package main
import(
	"fmt"
	"sort"
)

var mylist = []int{9,4,5,3,6,7,4,43,76,43,34,2,6,73}

func main(){
	sort.Ints(mylist)
	fmt.Println(mylist)
	
}
