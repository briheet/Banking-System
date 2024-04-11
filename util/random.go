package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

// Generates a random string of n characters
func RandomString(n int64) string {
	var sb strings.Builder
	k := int64(len(alphabets))

	for i := int64(0); i < n; i++ {
		c := alphabets[rng.Intn(int(k))]
		sb.WriteByte(c)

	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generates a random money number
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generates a random currency
func RandomCurrency() string {
	currencies := []string{
		"EUR",
		"USD",
		"CAD",
	}

	n := len(currencies)

	return currencies[rng.Intn(n)]
}
