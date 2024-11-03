package minigames

import (
	"math/rand"
	"time"

	"github.com/alireza/golang_Minigames/minigames/collections"
)

var (
	source         rand.Source
	random         *rand.Rand
	wordCollection []string
)

func init() {
	refreshSeed()
	wordCollection = collections.SetWordCollection()
}

func refreshSeed() {
	source = rand.NewSource(time.Now().Unix())
	random = rand.New(source)
}
