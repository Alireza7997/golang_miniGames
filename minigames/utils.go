package minigames

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/alireza/golang_Minigames/minigames/setting"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func listGames(m *Minigames) (games []game) {
	t := reflect.TypeOf(m).Elem()

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.Name == "Start" {
			continue
		}

		// Store both the method name and the method value in the Game struct
		games = append(games, game{
			Name:  formatGameName(method.Name),
			Value: reflect.ValueOf(m).MethodByName(method.Name),
		})
	}
	return games
}

func formatGameName(name string) string {
	gameName := strings.ReplaceAll(name, "GuessTheNumber", "Guess The Number")
	gameName = strings.ReplaceAll(gameName, "Hangman", "Hangman")
	gameName = strings.ReplaceAll(gameName, "WordScramble", "Word Scramble")
	gameName = strings.ReplaceAll(gameName, "RockPaperScissors", "Rock Paper Scissors")
	gameName = strings.ReplaceAll(gameName, "WhichHand", "Which Hand")
	return gameName
}

func rematch(minigame func()) {
	var choice string

	fmt.Print("\033[1;34m")
	fmt.Println("====================================")
	fmt.Println("         Play Again?(yes/no)")
	fmt.Println("====================================")
	fmt.Print("\033[1;0m")
	fmt.Scanln(&choice)

	clear()
	switch choice {
	case "yes", "y", "yep":
		refreshSeed()
		minigame()
	case "no", "n", "nope":
		fmt.Println("Thanks For Playing")
	default:
		rematch(minigame)
	}
}

func setCollections() {
	setCategory()
	wordCollection = setting.WordCollections[setting.Category][setting.Difficulty]
}

func setCategory() {
	var category string
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Choose Game Category        " + "\033[0m")
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")

	ctgNum := 1
	categories := make([]string, 0, len(setting.WordCollections))

	for ctg := range setting.WordCollections {
		fmt.Printf("\033[1;34m"+"  %d. %s\n"+"\033[0m", ctgNum, ctg)
		categories = append(categories, ctg)
		ctgNum++
	}

	fmt.Print("\033[1;34m" + "Please select the game Category: " + "\033[0m")
	fmt.Scanln(&category)
	category = titleCaser.String(strings.ToLower(category))

	if num, err := strconv.Atoi(category); err == nil && num > 0 && num <= len(categories) {
		category = categories[num-1]
	} else {
		category = titleCaser.String(strings.ToLower(category))
	}

	_, ok := setting.WordCollections[category]
	for !ok {
		fmt.Print("\033[1;34m" + "Invalid Category. Please choose a valid Category: " + "\033[0m")
		fmt.Scanln(&category)
		category = titleCaser.String(strings.ToLower(category))

		if num, err := strconv.Atoi(category); err == nil && num > 0 && num <= len(categories) {
			category = categories[num-1]
		}
		_, ok = setting.WordCollections[category]
	}
	setting.Category = category
	clear()
}

func setDifficulty() {
	var difficulty string
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Choose Game Difficulty        " + "\033[0m")
	fmt.Println("\033[1;34m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "  1. Easy                             " + "\033[0m")
	fmt.Println("\033[1;34m" + "  2. Medium                           " + "\033[0m")
	fmt.Println("\033[1;34m" + "  3. Hard                             " + "\033[0m")
	fmt.Print("\033[1;34m" + "Please select the game Difficulty(easy/medium/hard): " + "\033[0m")

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

	setting.Difficulty = difficulty
	clear()

}

func printSetting() {
	titleCaser := cases.Title(language.English)
	fmt.Println("\033[1;33m" + "=======================================" + "\033[0m")
	if setting.Category != "" {

		fmt.Println("\033[1;33m" + "            Category:   " + titleCaser.String(setting.Category) + "\033[0m")
	}
	fmt.Println("\033[1;33m" + "            Difficulty: " + titleCaser.String(setting.Difficulty) + "\033[0m")
	fmt.Println("\033[1;33m" + "=======================================" + "\033[0m")
}

func clear() {
	fmt.Print("\033[2J\033[H")
}
