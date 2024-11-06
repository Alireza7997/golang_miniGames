package minigames

import (
	"fmt"
	"math"
	"strings"
)

var (
	guesses int
)

func (m Minigames) Hangman() {
	setCollection()

	randomNumber := random.Intn(len(wordCollection))

	word := strings.ToLower(wordCollection[randomNumber])
	length := len(word)
	wordShown := word
	guesses = length
	hiddenLettersN := int(math.Round(float64(length) * 0.7))

	list := make(map[int]bool)
	for i := 0; i < hiddenLettersN; {
		randNum := random.Intn(length)

		if _, ok := list[randNum]; !ok {
			list[randNum] = true
			i++
		}
	}

	for index := range list {
		wordShown = replaceNthChar(wordShown, index, "_")
	}

	fmt.Println("Welcome to the Guessing Game!")
	for {
		if !strings.Contains(wordShown, "_") {
			fmt.Println("\033[32m" + "You Won!" + "\033[0m")
			fmt.Println("The word was: " + "\033[32m" + word + "\033[0m")
			break
		} else if guesses == 0 {
			fmt.Println("\033[31m" + "You Lost! The man has been hanged! (×_×)" + "\033[0m")
			fmt.Println("The word was: " + "\033[32m" + word + "\033[0m")
			break
		}

		fmt.Println("\n\033[33mWord:\033[0m " + wordShown + "\n")

		if guesses >= hiddenLettersN {
			fmt.Printf("\033[32m"+"Guesses remaining: %d\n"+"\033[0m", guesses)

		} else if guesses < hiddenLettersN && guesses >= int(math.Round(float64(length)*0.3)) {
			fmt.Printf("\033[33m"+"Guesses remaining: %d\n"+"\033[0m", guesses)

		} else if guesses < 5 && guesses >= 0 {
			fmt.Printf("\033[31m"+"Guesses remaining: %d\n"+"\033[0m", guesses)

		}

		var guess string
		fmt.Print("Guess a letter: ")
		fmt.Scan(&guess)
		guess = strings.ToLower(guess)
		fmt.Println()
		if guess == word {
			fmt.Println("\033[32m" + "You Won!" + "\033[0m")
			fmt.Println("The word was: " + "\033[32m" + word + "\033[0m")
			break
		}

		contains := strings.Contains(word, guess)
		if !contains {
			fmt.Println("[Your guess was wrong]")
			fmt.Println("\033[31m" + "---------------------" + "\033[0m")
			guesses--
		} else {
			fmt.Println("\033[32m" + "Correct!" + "\033[0m")
			fmt.Println("\033[32m" + "---------------------" + "\033[0m")

		}

		for i, letter := range word {
			if guess == string(letter) {
				wordShown = replaceNthChar(wordShown, i, string(letter))
			}

		}

	}

	rematch(m.Hangman)

}

func replaceNthChar(s string, n int, newChar string) string {
	if n < 0 || n >= len(s) {
		panic("Invalid character index")
	}
	return s[:n] + newChar + s[n+1:]
}
