package minigames

import (
	"math/rand"
	"time"
)

var (
	source         rand.Source
	random         *rand.Rand
	wordCollection []string
)

func init() {
	refreshSeed()
}

func refreshSeed() {
	source = rand.NewSource(time.Now().Unix())
	random = rand.New(source)
}
