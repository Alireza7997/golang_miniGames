package minigames

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/alireza/golang_Minigames/minigames/setting"
)

type Minigames struct{}

type game struct {
	Name  string
	Value reflect.Value
}

type gameInterface interface {
	GuessTheNumber()
	Hangman()
	WordScramble()
	RockPaperScissors()
	WhichHand()
	Start()
}

func New() gameInterface {
	return &Minigames{}
}

func (m *Minigames) Start() {
	for {
		fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
		fmt.Println("\033[1;34m" + "        Select a Mini-Game         " + "\033[0m")
		fmt.Println("\033[1;34m" + "====================================" + "\033[0m")

		games := listGames(m)
		gamesMap := map[int]string{}

		for i, game := range games {
			fmt.Printf("\033[1;34m"+"%d. %s\n", i+1, game.Name+"\033[0m")
			gamesMap[i+1] = game.Name
		}

		fmt.Print("\n\033[1;33m" + "Select the game of your choice (by number): " + "\033[0m")

		var choice string
		fmt.Scanln(&choice)

		if num, err := strconv.Atoi(choice); err == nil {
			if gameName, exists := gamesMap[num]; exists {
				for _, game := range games {
					if game.Name == gameName {
						setDifficulty()
						if setting.GamesWithCollection[game.Name] {
							setCollections()
						}
						printSetting()
						game.Value.Call(nil)
						clear()
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
		clear()
		fmt.Println("\033[1;31m" + "====================================" + "\033[0m")
		fmt.Println("\033[1;31m" + "Invalid selection. Please try again." + "\033[0m")
		fmt.Println("\033[1;31m" + "====================================" + "\033[0m")
		fmt.Println()

	}

}

//* to be used later
// func (m Minigames) Exit() {
// 	fmt.Println("Thanks for playing :D")
// }
