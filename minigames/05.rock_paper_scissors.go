package minigames

import (
	"fmt"
	"strings"
)

func (m Minigames) RockPaperScissors() {
	choices := []string{"rock", "paper", "scissors"}
	var computerChoice, playerChoice string
	var rounds, playerPoints, computerPoints int

	fmt.Print("\033[1;34m")
	fmt.Println("=======================================")
	fmt.Println("           Rock-Paper-Scissors         ")
	fmt.Println("=======================================")
	fmt.Print("\033[1;36m" + "\nEnter number of rounds to play (3 / 5 / 7): ")
	fmt.Print("\033[1;0m")
	fmt.Scanln(&rounds)

	if rounds != 3 && rounds != 5 && rounds != 7 {
		fmt.Println("\033[1;31m" + "\nError: Invalid number of rounds! Please choose 3, 5, or 7." + "\033[1;36m")
		return
	}

	for i := 0; i < rounds; i++ {
		computerChoice = choices[random.Intn(len(choices))]

		fmt.Print("\033[1;33m")
		fmt.Println("\n============================")
		fmt.Printf("          Round %d\n", i+1)
		fmt.Println("============================")
		fmt.Print("\033[0m")

		fmt.Print("\033[1;36m" + "Choose (1.Rock / 2.Paper / 3.Scissors): " + "\033[0m")
		fmt.Scanln(&playerChoice)
		playerChoice = strings.ToLower(playerChoice)

		fmt.Print(getAsciiArt(playerChoice))
		fmt.Printf("\nComputer chose: ")
		fmt.Print(getAsciiArt(computerChoice))

		outcomeMessage := ""
		switch playerChoice {
		case "rock", "1":
			switch computerChoice {
			case "rock":
				outcomeMessage = "\033[1;33m" + "It's a draw" + "\033[0m"
			case "paper":
				outcomeMessage = "\033[1;31m" + "You lost this round" + "\033[0m"
				computerPoints++
			case "scissors":
				outcomeMessage = "\033[1;32m" + "You won this round" + "\033[0m"
				playerPoints++
			}
		case "paper", "2":
			switch computerChoice {
			case "rock":
				outcomeMessage = "\033[1;32m" + "You won this round" + "\033[0m"
				playerPoints++
			case "paper":
				outcomeMessage = "\033[1;33m" + "It's a draw" + "\033[0m"
			case "scissors":
				outcomeMessage = "\033[1;31m" + "You lost this round" + "\033[0m"
				computerPoints++
			}
		case "scissors", "3":
			switch computerChoice {
			case "rock":
				outcomeMessage = "\033[1;31m" + "You lost this round" + "\033[0m"
				computerPoints++
			case "paper":
				outcomeMessage = "\033[1;32m" + "You won this round" + "\033[0m"
				playerPoints++
			case "scissors":
				outcomeMessage = "\033[1;33m" + "It's a draw" + "\033[0m"
			}
		}

		fmt.Println("\033[1m" + "\n============================" + "\033[0m")
		fmt.Printf("\033[1m"+"Outcome: %s\n", outcomeMessage)
		fmt.Println("\033[1m" + "============================" + "\033[0m")
	}

	fmt.Println("\033[1m" + "\n===================================" + "\033[0m")
	fmt.Println("          \033[36mFinal Results\033[0m              ")
	fmt.Println("\033[1m" + "===================================  " + "\033[0m")
	fmt.Printf("\033[1m"+"You: %d points\n"+"\033[0m", playerPoints)
	fmt.Printf("\033[1m"+"Computer: %d points\n", computerPoints)
	fmt.Println("\033[1m" + "===================================" + "\033[0m")

	if playerPoints > computerPoints {
		fmt.Println("\033[1;32m" + "Congratulations! You won the game!" + "\033[0m")
	} else if computerPoints > playerPoints {
		fmt.Println("\033[1;31m" + "The computer wins the game. Better luck next time!" + "\033[0m")
	} else {
		fmt.Println("\033[1;33m" + "The game is a draw. Well played!" + "\033[0m")
	}
	fmt.Println("\033[1;31m" + "===================================" + "\033[0m")

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
