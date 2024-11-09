package minigames

import (
	"fmt"

	"github.com/alireza/golang_Minigames/minigames/setting"
)

var (
	min        int
	max        int
	maxGuesses int
)

func (m Minigames) GuessTheNumber() {
	setGTNDiff()
	fmt.Print("\033[1;34m")
	fmt.Println("=======================================")
	fmt.Println("              Guessing Game            ")
	fmt.Println("=======================================")
	fmt.Print("\033[1;0")
	fmt.Printf("\033[1;36m\nGuess a number between [%d and %d]:\033[0m \n", min, max)

	secretNumber := random.Intn(max-min+1) + min

	for guesses := maxGuesses - 1; guesses >= 0; guesses-- {
		var guess int
		fmt.Printf("\033[1;33mGuess #%d: \033[0m", guesses+1)

		for {
			if _, err := fmt.Scan(&guess); err == nil {
				break
			}
			fmt.Println("\033[1;31mInvalid input. Please enter a number:\033[0m")
		}
		fmt.Println("\033[1;36m======================================================\033[0m")

		if guess == secretNumber {
			fmt.Println("\033[1;32mWell done! You guessed the number!\n\033[0m")
			break
		} else if guess < secretNumber {
			fmt.Println("\033[1mYour guess is \033[1;33mLOWER\033[0m \033[1;37mthan the secret number. \033[1;33mTry again!\033[0m")
		} else {
			fmt.Println("\033[1mYour guess is \033[1;33mHIGHER\033[0m \033[1;37mthan the secret number. \033[1;33mTry again!\033[0m")
		}

		fmt.Println("\033[1;36m======================================================\033[0m")

		if guesses == 0 {
			fmt.Printf("\033[1;31mOut of guesses! The number was \033[32m%d\033[0m.\n\n", secretNumber)
		}
	}

	rematch(m.GuessTheNumber)
}

func setGTNDiff() {
	diff := setting.Difficulty

	if diff == "Easy" {
		min = 1
		max = 100
		maxGuesses = 8
	} else if diff == "Medium" {
		min = 1
		max = 200
		maxGuesses = 6
	} else {
		min = 1
		max = 400
		maxGuesses = 6
	}
}
