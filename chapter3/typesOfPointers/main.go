package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt)) // Получает указатель на myint и выводит тип указателя
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) // Получает указатель на myFloat и выводит тип указателч
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool)) // Получает указатель на myBool и выводит тип указателя

	var myInt2 int
	var myIntPointer *int // Объявление переменной, содержащей указатель на int
	myIntPointer = &myInt2 // Указатель присваивается переменной
	fmt.Println(myIntPointer)

	var myFloat2 float64
	var myFloatPointer *float64 // Объявление переменной для хранения указателя на float64
	myFloatPointer = &myFloat2 // Указатель присваивается переменной
	fmt.Println(myFloatPointer)

	var myBool2 bool
	myBoolPointer := &myBool2 // Короткое объяыление переменной-указателя
	fmt.Println(myBoolPointer)

	var a int = 10
	fmt.Println(reflect.TypeOf(&a))
	var b float64 = 13.14
	fmt.Println(reflect.TypeOf(&b))



}
