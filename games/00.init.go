package games

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	source     rand.Source
	random     *rand.Rand
	category   string
	difficulty string
)

func init() {
	refreshSeed()
	chooseCategory()
	chooseDifficulty()
	setWordCollection()
}

func refreshSeed() {
	source = rand.NewSource(time.Now().Unix())
	random = rand.New(source)
}

func chooseDifficulty() {
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "=====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "     Choose Game Difficulty          " + "\033[0m")
	fmt.Println("\033[1;34m" + "=====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "  1. Easy                            " + "\033[0m")
	fmt.Println("\033[1;34m" + "  2. Medium                          " + "\033[0m")
	fmt.Println("\033[1;34m" + "  3. Hard                            " + "\033[0m")
	fmt.Print("\033[1;34m" + "Please enter your choice (easy/medium/hard): " + "\033[0m")

	fmt.Scanln(&difficulty)

	difficulty = titleCaser.String(strings.ToLower(difficulty))

	for difficulty != "Easy" && difficulty != "Medium" && difficulty != "Hard" {
		fmt.Print("\033[1;34m" + "Invalid choice. Please choose 'easy', 'medium', or 'hard': " + "\033[0m")
		fmt.Scanln(&difficulty)
		difficulty = titleCaser.String(strings.ToLower(difficulty))
	}

}

func chooseCategory() {
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "=====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "           Choose Game Category      " + "\033[0m")
	fmt.Println("\033[1;34m" + "=====================================" + "\033[0m")

	for categories := range wordCollections {
		fmt.Printf("\033[1;34m"+"  - %s\n"+"\033[0m", categories)
	}

	fmt.Print("\033[1;34m" + "Please enter the game category: " + "\033[0m")
	fmt.Scanln(&category)
	category = titleCaser.String(strings.ToLower(category))

	for _, ok := wordCollections[category]; !ok; {
		fmt.Print("\033[1;34m" + "Invalid category. Please choose a valid category: " + "\033[0m")
		fmt.Scanln(&category)
		category = titleCaser.String(strings.ToLower(category))
	}
}

func setWordCollection() {
	wordCollection = wordCollections[category][difficulty]
}
