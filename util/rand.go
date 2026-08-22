package util

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

var alphabets = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func GenerateAlphanumeric(length int) string {
	random := make([]string, length)
	for i := range length {
		// even = alphabet; odd = number.
		if rand.IntN(10)%2 == 0 {
			random[i] = alphabets[rand.IntN(26)]
		} else {
			random[i] = strconv.Itoa(rand.IntN(10))
		}
	}
	return strings.Join(random, "")
}
