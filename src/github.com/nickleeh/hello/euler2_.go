package main

import "fmt"

func main() {
	fibseq := [...]int{1, 1}
	c := 0
	sum := 0

	for c < 4e6 {
		c = fibseq[0] + fibseq[1]
		fibseq[0] = fibseq[1]
		fibseq[1] = c
		if c % 2 == 0 { sum += c }
	}
	fmt.Printf("Sum of all even fibonnaci seq terms < 4e6 = %d\n", sum)
}
