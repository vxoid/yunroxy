package user

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var prefix = []string{"dildo", "slut", "obama", "emo", "nude", "drugs", "pussy", "squirty", "nupid"}
var syfix = []string{"abuser", "smoker", "destroyer", "enjoyer", "stigger"}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func EmailGenerator(min int, max int) string {
	b := make([]byte, rand.Intn(max-min)+min)

	for i := 0; i < len(b); i++ {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	str := string(b[:])
	pref := prefix[seededRand.Intn(len(prefix))]
	syf := syfix[seededRand.Intn(len(syfix))]

	return fmt.Sprintf("%v%v%v@gmail.com", str, pref, syf)
}
