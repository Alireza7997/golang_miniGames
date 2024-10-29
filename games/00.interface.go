package games

type minigames struct{}

type gameInterface interface {
	GuessTheNumber()
	Hangman()
	WordScramble()
	RockPaperScissors()
	WhichHand()
}

func New() gameInterface {
	return &minigames{}
}
