package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("I am a Golang Developer")
	fmt.Println(reflect.TypeOf('₽'))
	fmt.Println(reflect.TypeOf('好'))
}
