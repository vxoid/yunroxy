package user

import (
	"math/rand"
)

func PassGenerator(min int, max int) string {
	b := make([]byte, rand.Intn(max-min)+min)
	for i := 0; i < len(b); i++ {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	str := string(b[:])
	return str
}
