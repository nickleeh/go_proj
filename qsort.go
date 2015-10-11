package main
import(
	"fmt"
	"math/rand"
)

var mylist = []int{9,4,5,3,6,7,4,43,76,43,34,2,6,73}

func qsort(a []int) []int {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := rand.Int() % len(a)

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qsort(a[:left])
  qsort(a[left + 1:])


  return a
}


func main(){
	result := qsort(mylist)
	fmt.Println(result)
}
