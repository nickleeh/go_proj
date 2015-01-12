package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

func prime_test(n int64, certainty int) (bool, float64) {
	var probobility float64
	i := big.NewInt(n)
	isPrime := i.ProbablyPrime(certainty)
	probobility = 1 - 1/math.Pow(4, 10)
	return isPrime, probobility
}

func why_not_prime(n int64) int64 {
	var i int64
	for i = 2; i < n/2; i++ {
		if n%i == 0 {
			return i
		}
	}
	return i
}

func main() {
	var n int64
	var certainty int
	var isPrime bool
	var probobility float64

	if len(os.Args) > 2 {
		var err error
		n, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		certainty, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	isPrime, probobility = prime_test(n, certainty)
	if isPrime {
		fmt.Printf("%d is probably %0.8f%% a prime.\n", n, probobility*100)
	} else {
		var i int64
		i = why_not_prime(n)
		fmt.Printf("%d is a composite because it can be divided by %d\n", n, i)
	}
}
