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

	fmt.Print("\033[1;34m")
	fmt.Println("=======================================")
	fmt.Println("             Hangman Game            ")
	fmt.Println("=======================================")
	fmt.Print("\033[0m")

	for {
		if !strings.Contains(wordShown, "_") {
			fmt.Println("\033[1;32m\nYou Won!\033[0m")
			fmt.Println("\033[1mThe word was: \033[1;32m" + word + "\033[0m")
			break
		} else if guesses == 0 {
			fmt.Println("\033[1;31m\nYou Lost! The man has been hanged! (×_×)\033[0m")
			fmt.Println("\033[1mThe word was: \033[1;32m" + word + "\033[0m")
			break
		}

		fmt.Println("\n\033[1;33mWord:\033[0m " + "[" + wordShown + "]" + "\n")

		fmt.Println("\033[1m==============================\033[0m")

		switch {
		case guesses >= hiddenLettersN:
			fmt.Printf("\033[1;32mGuesses remaining: %d\033[0m\n", guesses)
		case guesses < hiddenLettersN && guesses >= int(math.Round(float64(length)*0.3)):
			fmt.Printf("\033[1;33mGuesses remaining: %d\033[0m\n", guesses)
		case guesses < 5 && guesses >= 0:
			fmt.Printf("\033[1;31mGuesses remaining: %d\033[0m\n", guesses)
		}

		var guess string
		fmt.Print("\033[1;36mGuess a letter: \033[0m")
		fmt.Scan(&guess)
		guess = strings.ToLower(guess)
		fmt.Println("\033[1m==============================\n\033[0m")

		if guess == word {
			fmt.Println("\033[1;32mYou Won!\033[0m")
			fmt.Println("The word was: \033[1;32m" + word + "\033[0m")
			break
		}

		contains := strings.Contains(word, guess)
		if !contains {
			fmt.Println("\033[1;31m==============================\033[0m")
			fmt.Println("\033[1;31m[Your guess was wrong]\033[0m")
			fmt.Println("\033[1;31m==============================\033[0m")
			guesses--
		} else {
			fmt.Println("\033[1;32m==============================\033[0m")
			fmt.Println("\033[1;32mCorrect!\033[0m")
			fmt.Println("\033[1;32m==============================\033[0m")
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
