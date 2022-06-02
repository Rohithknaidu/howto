package utils

import (
	"math/rand"
	"time"
)

func NewID() string {
	numset := "0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 4)
	for i := range b {
		b[i] = numset[seededRand.Intn(len(numset))]
	}

	return string(b)
}
