package main 

import (
    "fmt"
    "math/rand"
    "time"
)

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}
func main() {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)

    xn := make([]float64, 10)

    for n := range xn{
        xn[n] = r.Float64()
    }

	fmt.Println(average(xn))
}
