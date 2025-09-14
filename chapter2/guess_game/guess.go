// Guess - Игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбрал случайное число от 1 до 100.")
	fmt.Println("Сможете угадать это число?")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Вы ввели значение, которое меньше загаданного")
		} else if guess > target {
			fmt.Println("Вы ввели значение, которое больше загаданного")
		} else {
			success = true
			fmt.Println("Отличная работа! Вы угадали загаданное число!")
			break
		}
	}
	if !success {
		fmt.Println("Сорян, вы не угадали число, это было:", target)
	}
}
