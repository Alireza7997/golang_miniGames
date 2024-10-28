package games

import "fmt"

func WordScramble() {
	refresh()
	word := wordCollection[random.Intn(len(wordCollection))]
	scrambled := scramble(word)
	var guess string

	fmt.Printf("Scrambled word: %s\nEnter the actual word: ", scrambled)
	fmt.Scan(&guess)

	if guess == word {
		fmt.Println("Correct! you guessed well")
	} else {
		fmt.Printf("Wrong! the answer is: %s\n", word)
	}

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
