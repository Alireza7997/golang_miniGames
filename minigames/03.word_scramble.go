package minigames

import (
	"fmt"
)

func (m Minigames) WordScramble() {

	fmt.Print("\033[1;34m")
	fmt.Println("=======================================")
	fmt.Println("             Word Scramble             ")
	fmt.Println("=======================================")
	fmt.Print("\033[0m")

	word := wordCollection[random.Intn(len(wordCollection))]
	scrambled := scramble(word)
	var guess string

	fmt.Printf("\033[1;33m\nScrambled word: %s\n\033[1;36mEnter the actual word: ", scrambled)
	fmt.Scan(&guess)

	fmt.Println("\033[1;37m\n==============================\033[0m")
	if guess == word {
		fmt.Println("\033[1;32m" + "Correct! you guessed well")
	} else {
		fmt.Printf("\033[1;31m"+"Wrong! the answer is: %s\n", word)
	}
	fmt.Println("\033[1;37m==============================\n\033[0m")

	rematch(m.WordScramble)
}

func scramble(word string) string {
	length := len(word)
	runes := []rune(word)

	for i := 0; i < length; i++ {
		replace := random.Intn(length)
		runes[i], runes[replace] = runes[replace], runes[i]

	}

	return string(runes)
}
