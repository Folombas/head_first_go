package main

import (
	"fmt"
)

func main() {
	amount := 6
	var myInt int = 13
	var myFloat float64 = 3.14
	var myBool bool = true
	fmt.Println(amount)
	fmt.Println(&amount)
	fmt.Println(myInt)
	fmt.Println(&myInt)
	fmt.Println(myFloat)
	fmt.Println(&myFloat)
	fmt.Println(myBool)
	fmt.Println(&myBool)
}
