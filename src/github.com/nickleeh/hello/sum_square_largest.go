package main

//import "fmt"

func sum_s_l(x, y, z int) int {
	if x <= y {
		if y <= z {
			return z
		}
		return y
	}
	return x
}
