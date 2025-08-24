package main

import "fmt"

func main() {
	// Объявление переменных
	var quantity int
	var length, width float64
	var customerName string

	// Присваивание значений переменным
	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	// Объявлять переменные и присваивать им значения можно сразу в одной строке
	var quantityX int = 4
	var lengthX, widthX float64 = 1.2, 2.4
	var customerNameX string = "Bruce Lee"

	// Чаще всего используют короткое объявление переменных
	quantityY := 4
	lengthY, widthY := 1.2, 2.4
	customerNameY := "Jackie Chan"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
	fmt.Println(quantityX)
	fmt.Println(lengthX*widthX, "square meters also")
	fmt.Println(customerNameX)
	fmt.Println(quantityY)
	fmt.Println(lengthY, widthY, "and this is also square meters")
	fmt.Println(customerNameY)
}
