package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

var numbers []float64
var sum float64 = 0

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Please enter numbers separated by space or comma.")
		fmt.Println(filepath.Base(os.Args[0]), "-h or --help for help.")
		return
	}

	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err := fmt.Errorf("usage: %s number1 number2 ... numbern",
			filepath.Base(os.Args[0]))
		fmt.Println(err)
		return
	}

	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			numbers = append(numbers, n)
		}
	}

	fmt.Println("Numbers are:", numbers)
	for _, value := range numbers {
		sum += value
	}

	average := sum / float64(len(numbers))
	fmt.Println("The average is:", average)

}
