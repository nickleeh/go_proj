package main

import (
    "fmt"
    "math/big"
)

func main() {
    // x := new(big.Int)
    // x.MulRange(1, 10)
    fmt.Println(factorial(123))
}

func factorial(n int64) *big.Int {
	// var x int
	result := new(big.Int)
	result.MulRange(1, n)
	return result
}