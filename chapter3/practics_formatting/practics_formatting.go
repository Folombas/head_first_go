package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s имеет обычно %d колеса\n", "Автомашина", 4)

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", false)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
	fmt.Printf("%f liters needed\n", 1.68999999999998)
}
