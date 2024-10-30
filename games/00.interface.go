package games

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type minigames struct{}

type gameInterface interface {
	GuessTheNumber()
	Hangman()
	WordScramble()
	RockPaperScissors()
	WhichHand()
	Start()
}

func New() gameInterface {
	return &minigames{}
}

type Game struct {
	Name  string
	Value reflect.Value
}

func (m *minigames) Start() {
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Select a Mini-Game         " + "\033[0m")
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")

	games := listGames(m)
	gamesMap := map[int]string{}

	for i, game := range games {
		fmt.Printf("\033[1;34m"+"%d. %s\n", i+1, game.Name+"\033[0m")
		gamesMap[i+1] = game.Name
	}

	fmt.Print("\033[1;33m" + "Select the game of your choice (by number): " + "\033[0m")

	var choice string
	fmt.Scanln(&choice)

	if num, err := strconv.Atoi(choice); err == nil {
		if gameName, exists := gamesMap[num]; exists {
			for _, game := range games {
				if game.Name == gameName {
					game.Value.Call(nil)
					return
				}
			}

		}
	}

	//* Will be added later
	// for _, gameName := range gamesMap {
	// 	if strings.TrimSpace(strings.ToLower(gameName)) == strings.TrimSpace(strings.ToLower(choice)) {
	// 		for _, game := range games {
	// 			if game.Name == gameName {
	// 				game.Value.Call(nil)
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	fmt.Println("Invalid selection. Please try again.")

}

func listGames(m *minigames) []Game {
	t := reflect.TypeOf(m).Elem()

	var games []Game

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.Name == "Start" {
			continue
		}

		// Store both the method name and the method value in the Game struct
		games = append(games, Game{
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
