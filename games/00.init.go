package games

import (
	"fmt"
	"math/rand"
	"strconv"
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
	setWordCollection()
}

func refreshSeed() {
	source = rand.NewSource(time.Now().Unix())
	random = rand.New(source)
}

func setCategory() {
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Choose Game Category        " + "\033[0m")
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")

	ctgNum := 1
	categories := make([]string, 0, len(wordCollections)) // Create a slice to hold category names

	// Print available categories and store them in a slice
	for ctg := range wordCollections {
		fmt.Printf("\033[1;34m"+"  %d. %s\n"+"\033[0m", ctgNum, ctg)
		categories = append(categories, ctg) // Store category in the slice
		ctgNum++
	}

	fmt.Print("\033[1;34m" + "Please select the game category: " + "\033[0m")
	fmt.Scanln(&category)
	category = titleCaser.String(strings.ToLower(category))

	// Check if the input is a number
	if num, err := strconv.Atoi(category); err == nil && num > 0 && num <= len(categories) {
		category = categories[num-1] // Update category to the selected category name
	} else {
		category = titleCaser.String(strings.ToLower(category))
	}

	// Validate the chosen category
	_, ok := wordCollections[category]
	for !ok {
		fmt.Print("\033[1;34m" + "Invalid category. Please choose a valid category: " + "\033[0m")
		fmt.Scanln(&category)
		category = titleCaser.String(strings.ToLower(category))

		// Recheck if the input is a number and map it if valid
		if num, err := strconv.Atoi(category); err == nil && num > 0 && num <= len(categories) {
			category = categories[num-1] // Update category to the selected category name
		}
		_, ok = wordCollections[category] // Validate again
	}
}

func setDifficulty() {
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Choose Game Difficulty        " + "\033[0m")
	fmt.Println("\033[1;34m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "  1. Easy                             " + "\033[0m")
	fmt.Println("\033[1;34m" + "  2. Medium                           " + "\033[0m")
	fmt.Println("\033[1;34m" + "  3. Hard                             " + "\033[0m")
	fmt.Print("\033[1;34m" + "Please select the game difficulty(easy/medium/hard): " + "\033[0m")

	fmt.Scanln(&difficulty)
	difficulty = titleCaser.String(strings.ToLower(difficulty))

	switch difficulty {
	case "1", "Easy":
		difficulty = "Easy"
	case "2", "Medium":
		difficulty = "Medium"
	case "3", "Hard":
		difficulty = "Hard"
	}

	for difficulty != "Easy" && difficulty != "Medium" && difficulty != "Hard" {
		fmt.Print("\033[1;34m" + "Invalid choice. Please choose 'easy', 'medium', or 'hard': " + "\033[0m")
		fmt.Scanln(&difficulty)

		switch difficulty {
		case "1":
			difficulty = "Easy"
		case "2":
			difficulty = "Medium"
		case "3":
			difficulty = "Hard"
		default:
			difficulty = titleCaser.String(strings.ToLower(difficulty))
		}
	}

}

func setWordCollection() {
	setCategory()
	setDifficulty()

	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;33m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;33m" + "           Category: " + titleCaser.String(category) + "\033[0m")
	fmt.Println("\033[1;33m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;33m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;33m" + "           Difficulty: " + titleCaser.String(difficulty) + "\033[0m")
	fmt.Println("\033[1;33m" + "======================================" + "\033[0m")

	wordCollection = wordCollections[category][difficulty]
}
