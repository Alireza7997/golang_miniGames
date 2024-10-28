package games

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())

var random = rand.New(source)
