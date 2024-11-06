package minigames

import (
	"fmt"
	"strings"
)

func (m Minigames) RockPaperScissors() {
	choices := []string{"rock", "paper", "scissors"}
	var computerChoice, playerChoice string
	var rounds, playerPoints, computerPoints int

	fmt.Println("===================================")
	fmt.Println("     Welcome to Rock-Paper-Scissors")
	fmt.Println("===================================")
	fmt.Print("Enter number of rounds to play (3, 5, or 7): ")
	fmt.Scanln(&rounds)

	if rounds != 3 && rounds != 5 && rounds != 7 {
		fmt.Println("\nError: Invalid number of rounds! Please choose 3, 5, or 7.")
		return
	}

	fmt.Println("\nLet the game begin!")

	for i := 0; i < rounds; i++ {
		computerChoice = choices[random.Intn(len(choices))]

		fmt.Println("\n-----------------------------")
		fmt.Printf("Round %d\n", i+1)
		fmt.Println("-----------------------------")

		fmt.Print("Your choice (1.Rock / 2.Paper / 3.Scissors): ")
		fmt.Scanln(&playerChoice)
		playerChoice = strings.ToLower(playerChoice)

		fmt.Print(getAsciiArt(playerChoice))
		fmt.Printf("\nComputer chose: %s", computerChoice)
		fmt.Print(getAsciiArt(computerChoice))

		outcomeMessage := ""
		switch playerChoice {
		case "rock", "1":
			switch computerChoice {
			case "rock":
				outcomeMessage = "It's a draw!"
			case "paper":
				outcomeMessage = "You lost this round."
				computerPoints++
			case "scissors":
				outcomeMessage = "You won this round!"
				playerPoints++
			}
		case "paper", "2":
			switch computerChoice {
			case "rock":
				outcomeMessage = "You won this round!"
				playerPoints++
			case "paper":
				outcomeMessage = "It's a draw!"
			case "scissors":
				outcomeMessage = "You lost this round."
				computerPoints++
			}
		case "scissors", "3":
			switch computerChoice {
			case "rock":
				outcomeMessage = "You lost this round."
				computerPoints++
			case "paper":
				outcomeMessage = "You won this round!"
				playerPoints++
			case "scissors":
				outcomeMessage = "It's a draw!"
			}
		}

		fmt.Println("-----------------------------")
		fmt.Printf("Outcome: %s\n", outcomeMessage)
		fmt.Println("-----------------------------")
	}

	fmt.Println("\n===================================")
	fmt.Println("               Final Results")
	fmt.Println("===================================")
	fmt.Printf("You: %d points\n", playerPoints)
	fmt.Printf("Computer: %d points\n", computerPoints)
	fmt.Println("===================================")

	if playerPoints > computerPoints {
		fmt.Println("Congratulations! You won the game!")
	} else if computerPoints > playerPoints {
		fmt.Println("The computer wins the game. Better luck next time!")
	} else {
		fmt.Println("The game is a draw. Well played!")
	}
	fmt.Println("===================================")

	rematch(m.RockPaperScissors)
}

func getAsciiArt(choice string) string {
	rockArt := `
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
`

	paperArt := `
     _______
---'    ____)____
           ______)
          _______)
         _______)
---.__________)
`

	scissorsArt := `
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
`

	switch choice {
	case "rock", "1":
		return rockArt
	case "paper", "2":
		return paperArt
	case "scissors", "3":
		return scissorsArt
	default:
		return ""
	}
}
