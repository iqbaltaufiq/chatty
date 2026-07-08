package util

import (
	"math/rand/v2"
	"strconv"
)

type (
	alphabet [26]string
	number   [10]int
)

var Alphabets = alphabet{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

func (s alphabet) RandomAlphabet() string {
	return s[rand.IntN(len(s))]
}

var Numbers = number{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func (n number) RandomNumber() int {
	return n[rand.IntN(len(n))]
}

func (n number) RandomNumberToString() string {
	return strconv.Itoa(n.RandomNumber())
}
