package util

import (
	"math/rand/v2"
	"strconv"
)

var alphabets = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var numbers = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func GenerateAlphanumeric(length int) string {
	var (
		n             = 0
		random        = ""
		alphaOrNumber = false
	)
	for range length {
		// even = alphabet; odd = number.
		n = rand.IntN(10)
		alphaOrNumber = n%2 == 0
		if alphaOrNumber {
			random += alphabets[rand.IntN(26)]
		} else {
			random += strconv.Itoa(numbers[rand.IntN(10)])
		}
	}
	return random
}
