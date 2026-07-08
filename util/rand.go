package util

import (
	"math/rand/v2"
	"strconv"
)

type (
	alphabet [26]string
	number   [10]int
)

func alphabets() alphabet {
	return alphabet{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
}

func (s alphabet) RandomAlphabet() string {
	return s[rand.IntN(len(s))]
}

func numbers() number {
	return number{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func (n number) RandomNumber() int {
	return n[rand.IntN(len(n))]
}

func (n number) RandomNumberToString() string {
	return strconv.Itoa(n.RandomNumber())
}

func GenerateAlphanumeric(length int) string {
	var (
		alphabets     = alphabets()
		numbers       = numbers()
		n             = 0
		random        = ""
		alphaOrNumber = false
	)
	for range length {
		// even = alphabet; odd = number.
		n = numbers.RandomNumber()
		alphaOrNumber = n%2 == 0 && n != 0
		if alphaOrNumber {
			random = random + alphabets.RandomAlphabet()
		} else {
			random = random + numbers.RandomNumberToString()
		}
	}
	return random
}
