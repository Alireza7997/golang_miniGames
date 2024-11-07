package minigames

import (
	"fmt"
)

func (m Minigames) WhichHand() {
	var (
		goal   int
		choice int
	)

	fmt.Print("\033[1;34m")
	fmt.Println("=======================================")
	fmt.Println("               Which Hand              ")
	fmt.Println("=======================================")
	fmt.Print("\033[0m")

	goal = random.Intn(2) + 1

	fmt.Print("\033[1;36m" + "\nGuess which hand the goal is in: \n" + "\033[0m")
	drawHands()
	fmt.Print("\033[1m" + "Your choice:")
	fmt.Scan(&choice)

	if goal == choice {
		fmt.Print("\033[1;32m" + "Correct!")
	} else if choice < 1 || choice > 2 {
		fmt.Print("\033[1;31m" + "Invalid choice")
	} else {
		fmt.Print("\033[1;31m" + "Incorrect")
	}
	fmt.Printf("\nThe goal was in hand \033[1;32m%d\n\n", goal)

	rematch(m.WhichHand)
}

// Will be enhanced later :D
func drawHands() {
	fmt.Println(" _ _ _ _         _ _ _ _ ")
	fmt.Println("| | | | |       | | | | |")
	fmt.Println("|_|_|_|_||     ||_|_|_|_|")
	fmt.Println("   (1)             (2)   ")
	fmt.Println()
}
