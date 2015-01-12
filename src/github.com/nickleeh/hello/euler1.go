package main

import "fmt"

func main() {

	const CEILIN_NUMBER = 1e10
	var picked []int
	for i := 1; i < CEILIN_NUMBER; i++ {
		if i%3 == 0 || i%5 == 0 {
			picked = append(picked, i)
		}
	}

	// 	fmt.Println(picked)

	sum := 0
	for _, j := range picked {
		sum += j
	}

	fmt.Printf("%s %d %s \n%d\n", "The sum of all the multiples of 3 or 5 below", CEILIN_NUMBER, "is:", sum)

}
