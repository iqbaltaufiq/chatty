package util

import (
	"bytes"
	"math/rand/v2"
	"strconv"
	"sync"
)

var alphabets = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GenerateAlphanumeric(length int) string {
	random := bufferPool.Get().(*bytes.Buffer)
	random.Reset()
	defer bufferPool.Put(random)

	for range length {
		// even = alphabet; odd = number.
		if rand.IntN(10)%2 == 0 {
			random.WriteString(alphabets[rand.IntN(26)])
		} else {
			random.WriteString(strconv.Itoa(rand.IntN(10)))
		}
	}
	return random.String()
}
