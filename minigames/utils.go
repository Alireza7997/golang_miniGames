package minigames

import (
	"fmt"
	"reflect"
	"strings"
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
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "         Play Again?(yes/no)" + "\033[0m]")
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
	fmt.Scanln(&choice)

	switch choice {
	case "yes", "y", "yep":
		refreshSeed()
		minigame()
	case "no", "n", "nope":
	default:
		rematch(minigame)
	}
}
