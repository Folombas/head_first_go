package main

import (
	"fmt"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("Can`t get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squareRoot(81.00)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f", root)
	}
}
