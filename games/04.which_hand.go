package games

import (
	"fmt"
)

func (m minigames) WhichHand() {
	var (
		goal   int
		choice int
	)

	goal = random.Intn(2) + 1

	fmt.Print("Guess which hand the goal is in: \n")
	drawHands()
	fmt.Print("Your choice:")
	fmt.Scan(&choice)

	fmt.Println(goal)
	if goal == choice {
		fmt.Println("Correct!")
	} else if choice < 1 || choice > 2 {
		fmt.Println("Incorrect, but nice try")
	} else {
		fmt.Println("Incorrect")
	}

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
