package main

import "fmt"

func main() {

	const Limit = 4e6
	a, b := 1, 1
	total := 0
	for a <= Limit {
		if a%2 == 0 {
			total += a
		}
		a, b = b, a+b
	}
	fmt.Println(total)

}

// 4613732
