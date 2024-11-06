package minigames

import (
	"fmt"
)

func (m Minigames) RockPaperScissors() {

	choices := []string{"rock", "paper", "scissors"}
	var computerChoice string
	var playerChoice string
	var rounds int
	var playerPoints int
	var computerPoints int

	fmt.Println("Select number of rounds to play(3-5-7):")
	fmt.Scanln(&rounds)

	if rounds != 3 && rounds != 5 && rounds != 7 {
		fmt.Println("Invalid number of rounds!!!")
		return
	}

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

	rematch(m.RockPaperScissors)
}
