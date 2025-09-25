package main

import "fmt"

func main() {
	repeatLine("Hello!", 5)
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
