package games

import (
	"fmt"
)

func (g minigames) RockPaperScissors() {
	// Making variables
	choices := []string{"rock", "paper", "scissors"}
	var computerChoice string
	var playerChoice string
	var rounds int
	var playerPoints int
	var computerPoints int
	var rematch string

	// The Game Begins !!
	fmt.Println("Select number of rounds to play(3-5-7):")
	fmt.Scanln(&rounds)

	// Checking if the entered number of rounds is valid
	if rounds != 3 && rounds != 5 && rounds != 7 {
		fmt.Println("Invalid number of rounds!!!")
		return
	}

	// The Real Game Begins !!!
	for i := 0; i < rounds; i++ {
		computerChoice = choices[random.Intn(len(choices))]
		fmt.Println("Choose (rock / paper / scissors) :")
		fmt.Scanln(&playerChoice)

		fmt.Println("      {")
		fmt.Printf("	Computer choice : %s\n", computerChoice)
		if computerChoice == "rock" {
			switch playerChoice {
			case "rock":
				fmt.Println("	*Draw*")
			case "paper":
				fmt.Println("	*You Won*")
				playerPoints++
			case "scissors":
				fmt.Println("	*You Lost*")
				computerPoints++
			}
		}
		if computerChoice == "paper" {
			switch playerChoice {
			case "rock":
				fmt.Println("	*You Lost*")
				computerPoints++
			case "paper":
				fmt.Println("	*Draw*")
			case "scissors":
				fmt.Println("	*You Won*")
				playerPoints++
			}

		}
		if computerChoice == "scissors" {
			switch playerChoice {
			case "rock":
				fmt.Println("	*You Won*")
				playerPoints++
			case "paper":
				fmt.Println("	*You Lost*")
				computerPoints++
			case "scissors":
				fmt.Println("	*Draw*")
			}
		}
		fmt.Println("		  }")
	}
	fmt.Println("**********")
	if playerPoints > computerPoints {
		fmt.Println("Game Result : You Won the Game!")
	} else if computerPoints > playerPoints {
		fmt.Println("Game Result : Computer Won the Game")
	} else {
		fmt.Println("Game Result : Draw")
	}
	fmt.Println("**********")

	// Asking if the player wants to player another game
	fmt.Println("Want to play another Game ? (Yes/No)")
	fmt.Scanln(&rematch)

	if rematch == "Yes" || rematch == "yes" {
		refresh()
		g.RockPaperScissors()
	} else if rematch == "No" || rematch == "no" {
		fmt.Println("Thanks for playing :)")
		return
	} else {
		fmt.Println("Fuck You :)")
		return
	}
}
