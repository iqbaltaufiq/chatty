package util

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

var alphabets = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func GenerateAlphanumeric(length int) string {
	n := 0
	random := [6]string{}
	alphaOrNumber := false
	i := 0
	for range length {
		// even = alphabet; odd = number.
		n = rand.IntN(10)
		alphaOrNumber = n%2 == 0
		if alphaOrNumber {
			random[i] = alphabets[rand.IntN(26)]
		} else {
			random[i] = strconv.Itoa(rand.IntN(10))
		}
		i++
	}
	return strings.Join(random[:], "")
}
