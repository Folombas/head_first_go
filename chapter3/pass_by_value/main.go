package main

import "fmt"

func main() {
	amount := 10
	double(amount)
}

func double(number int) {
	number *= 2
	fmt.Println(number)
}
