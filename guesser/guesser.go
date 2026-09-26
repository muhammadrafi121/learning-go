package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	greetings := "=================================\n" +
		"GUESSER GAME\n\n" +
		"Guess a number between 1 to 100\n" +
		"You have five chances\n" +
		"=================================\n"

	fmt.Println(greetings)

	number := rand.IntN(100) + 1
	chances := 5
	var guess int

	for chances > 0 {

		fmt.Print("Your guess: ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if guess > number {
			fmt.Println("Guess lower\n")
			chances--
		} else if guess < number {
			fmt.Println("Guess higher\n")
			chances--
		} else {
			fmt.Println("You won!")
			return
		}
	}

	fmt.Println("You lose :(")
}
