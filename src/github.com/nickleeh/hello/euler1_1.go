package main

import "fmt"

func main() {

	const Ceilin_number = 1e9
	sum := 0
	for i := 1; i < Ceilin_number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("%s %.0f %s \n%d\n", "The sum of all the multiples of 3 or 5 below", Ceilin_number, "is:", sum)

}
