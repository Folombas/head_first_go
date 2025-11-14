package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(myInt)
	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt)) // Получает указатель на myint и выводит тип указателя
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) // Получает указатель на myFloat и выводит тип указателя
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool)) // Получает указатель на myBool и выводит тип указателя
	fmt.Println("============================================================================")

	var myInt2 int
	fmt.Println(myInt2)        // 0
	fmt.Println(&myInt2)       // Получаем указатель на myInt2 - 0xc00000a110
	var myIntPointer *int      // Объявление переменной, содержащей указатель на int
	fmt.Println(&myIntPointer) // 0xc00004e040
	fmt.Println(myIntPointer)  // <nil>
	myIntPointer = &myInt2     // Указатель присваивается переменной
	fmt.Println(myIntPointer)  // 0xc00000a110

	var myFloat2 float64
	var myFloatPointer *float64 // Объявление переменной для хранения указателя на float64
	myFloatPointer = &myFloat2  // Указатель присваивается переменной
	fmt.Println(myFloatPointer) // 0xc00000a130

	var myBool2 bool
	myBoolPointer := &myBool2  // Короткое объяыление переменной-указателя
	fmt.Println(myBoolPointer) // 0xc00000a138

	fmt.Println("============================================================================")
	var a int = 10
	fmt.Println(reflect.TypeOf(&a))
	var b float64 = 13.14
	fmt.Println(reflect.TypeOf(&b))

}
