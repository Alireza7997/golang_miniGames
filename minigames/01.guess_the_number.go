package minigames

import (
	"fmt"
)

var (
	min = 1
	max = 100
)

func (m Minigames) GuessTheNumber() {
	secretNumber := random.Intn(max-min+1) + min

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println()
	fmt.Printf("Guess a number between %d and %d: \n", min, max)

	maxGuesses := 6

	for guesses := maxGuesses - 1; guesses >= 0; guesses-- {
		var guess int
		fmt.Printf("Guess #%d: ", guesses+1)

		// Validate user input
		for _, err := fmt.Scan(&guess); err != nil; _, err = fmt.Scan(&guess) {
			fmt.Println("Invalid input. Please enter a number:")
		}

		if guess == secretNumber {

			fmt.Println("\033[32m" + "Well! You guessed the number!" + "\033[0m")
			break
		} else if guess < secretNumber {

			fmt.Println("\033[33mYour guess is lower than the goal, Try again.\033[0m")
		} else {

			fmt.Println("\033[33mYour guess is higher than the goal, Try again.\033[0m")
		}

		fmt.Println("\033[36m" + "---------------------" + "\033[0m")

		if guesses == 0 {
			fmt.Printf("\033[31m"+"out of guesses!"+" The number was \033[32m%d\033[0m.\n", secretNumber)

		}
	}

	rematch(m.GuessTheNumber)
}
