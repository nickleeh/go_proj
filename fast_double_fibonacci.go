// Fast doubling Fibonacci algorithm
package main

import (
    "fmt"
    "math/big"
)

//  (Public) Returns F(n).
func fibonacci(n int) *big.Int {
    if n < 0 {
        panic("Negative arguments not implemented")
    }
    fst, _ := fib(n)
    return fst
}

// (Private) Returns the tuple (F(n), F(n+1)).
func fib(n int) (*big.Int, *big.Int) {
    if n == 0 {
        return big.NewInt(0), big.NewInt(1)
    }
    a, b := fib(n / 2)
    c := Mul(a, Sub(Mul(b, big.NewInt(2)), a))
    d := Add(Mul(a, a), Mul(b, b))
    if n%2 == 0 {
        return c, d
    } else {
        return d, Add(c, d)
    }
}

func main() {
    fmt.Println(fibonacci(123))
    fmt.Println(fibonacci(124))
}

func Mul(x, y *big.Int) *big.Int {
    return big.NewInt(0).Mul(x, y)
}
func Sub(x, y *big.Int) *big.Int {
    return big.NewInt(0).Sub(x, y)
}
func Add(x, y *big.Int) *big.Int {
    return big.NewInt(0).Add(x, y)
}
