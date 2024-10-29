package games

import (
	"math/rand"
	"time"
)

var (
	source rand.Source

	random *rand.Rand
)

var wordCollection = []string{"pineapple", "watermelon", "grapefruit", "blackberry", "blueberry", "raspberry", "octopus", "elephant", "platypus"}

func refresh() {
	source = rand.NewSource(time.Now().Unix())
	random = rand.New(source)
}

func init() {
	refresh()
}
