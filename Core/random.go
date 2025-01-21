package core

import (
	"math/rand"
	"time"
)

func GenerateRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
